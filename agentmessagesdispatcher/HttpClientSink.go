package agentmessagesdispatcher

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/adgs85/gomonagent/agentconfiguration"
	"github.com/adgs85/gomonagent/agentlogger"
	"github.com/adgs85/gomonmarshalling/monmarshalling"
)

type httpClient struct {
	clientChannel chan monmarshalling.Stat
	clientSink    StatSinkFuncType
	baseUrl       string
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
		agentlogger.Logger().Fatalln(err)
	}

	url := buildUrl(client, stat)
	println("%%%%%%%%%" + url)
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		agentlogger.Logger().Println(err)
		return
	}

	if resp.StatusCode != 202 {
		agentlogger.Logger().Println("WARN unexpected http status code:", resp.StatusCode)
	}

	if resp.Body != nil {
		resp.Body.Close()
	}
}

func buildUrl(client *httpClient, stat *monmarshalling.Stat) string {
	return client.baseUrl + "/" + stat.MetaData.MessageType + "/" + stat.MetaData.StatType
}

func StartHttpClientSenderLoopReturnSink() StatSinkFuncType {
	warnInactivityMinutes := 2
	baseUrl := agentconfiguration.GlobalCfg().ServerUrl
	hClient := NewHttpClient(baseUrl)
	go func() {
		for {
			select {
			case stat := <-hClient.clientChannel:
				hClient.post(&stat)
			case <-time.After(time.Duration(warnInactivityMinutes) * time.Minute):
				log.Println("WARN nothing to post to server for", warnInactivityMinutes, "minutes")
			}
		}
	}()
	return hClient.clientSink
}
