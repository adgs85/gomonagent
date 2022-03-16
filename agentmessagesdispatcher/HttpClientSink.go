package agentmessagesdispatcher

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/adgs85/gomonagent/agentconfiguration"
	"github.com/adgs85/gomonagent/agentlogger"
	"github.com/adgs85/gomonmarshalling/monmarshalling"
)

var logger = agentlogger.Logger()

type httpClient struct {
	clientChannel      chan monmarshalling.Stat
	clientSink         StatSinkFuncType
	baseUrl            string
	httpClientInstance *http.Client
}

func (client *httpClient) GetStatSink() StatSinkFuncType {
	return client.clientSink
}

func (client *httpClient) PostImmediately(stat *monmarshalling.Stat) {
	client.post(stat)
}

func NewHttpClient(baseUrl string) *httpClient {
	httpClientChannel := make(chan monmarshalling.Stat)
	hClient := httpClient{
		clientChannel: httpClientChannel,
		clientSink:    NewChannelSinkStallLogging(httpClientChannel),
		baseUrl:       baseUrl,
	}

	return &hClient
}

func (client *httpClient) post(stat *monmarshalling.Stat) {

	json_data, err := json.Marshal(stat)

	if err != nil {
		logger.Fatalln(err)
	}

	url := buildUrl(client, stat)

	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		logger.Println(err, "Event lost!")
		return
	}

	if resp.StatusCode >= 300 {
		logger.Println("WARN Request to", url, "received", resp.StatusCode, "http status.", "Event lost!")
	}

	if resp.Body != nil {
		resp.Body.Close()
	}
}

func buildUrl(client *httpClient, stat *monmarshalling.Stat) string {
	url := client.baseUrl + "/" + stat.MetaData.MessageType
	if len(stat.MetaData.StatType) > 0 {
		return url + "/" + stat.MetaData.StatType
	}
	return url
}

func StartHttpClientSenderLoopReturnSink() StatSinkFuncType {
	cfg := agentconfiguration.GlobalCfg()

	println("Starting http client loop with base url:", cfg.ServerUrl)
	warnInactivityMinutes := 2
	baseUrl := cfg.ServerUrl
	hClient := NewHttpClient(baseUrl)
	go func() {
		hClient.httpClientInstance = &http.Client{
			Timeout: time.Duration(cfg.RequestTimeoutSec) * time.Second,
		}
		for {
			select {
			case stat := <-hClient.clientChannel:
				hClient.post(&stat)
			case <-time.After(time.Duration(warnInactivityMinutes) * time.Minute):
				logger.Println("WARN nothing to post to server for", warnInactivityMinutes, "minutes")
			}
		}
	}()
	return hClient.clientSink
}
