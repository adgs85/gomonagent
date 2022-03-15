package main

import (
	"fmt"
	"sync"

	"os"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/adgs85/gomonagent/agentmessagesdispatcher"
	"github.com/adgs85/gomonagent/diskstats"

	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
)

func main() {

	hostname, err := os.Hostname()
	now := time.Now().UnixMilli()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(spew.Sdump((envconfig.GetViperConfig().GetString("disk_polling_rate_ms"))))
	fmt.Printf("Hostname: %s %v\n", hostname, now)

	// statChannel := make(chan monmarshalling.Stat)

	var wg sync.WaitGroup
	wg.Add(1)

	httpClientSink := agentmessagesdispatcher.StartHttpClientSenderLoopReturnSink()

	diskstats.StartDiskInfoLoopGoRoutine(httpClientSink)

	wg.Wait()
}
