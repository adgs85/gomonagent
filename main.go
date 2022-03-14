package main

import (
	"fmt"

	"os"
	"time"

	//	"github.com/adgs85/gomonagent/diskstats"
	//"github.com/spf13/viper"
	//"github.com/adgs85/gomonmarshalling/monmarshalling"
	"github.com/davecgh/go-spew/spew"

	"github.com/adgs85/gomonagent/agentmessagesdispatcher"
	"github.com/adgs85/gomonagent/diskstats"

	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
)

func main() {

	//	fmt.Println(viper.GetString("disk_polling_rate_ms"))

	hostname, err := os.Hostname()
	now := time.Now().UnixMilli()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//str := viper.GetString("disk_polling_rate_ms");
	//fmt.Println(str)
	fmt.Println(spew.Sdump((envconfig.GetViperConfig().GetString("disk_polling_rate_ms"))))
	fmt.Printf("Hostname: %s %v\n", hostname, now)

	diskstats.New(agentmessagesdispatcher.Dispatch)
}
