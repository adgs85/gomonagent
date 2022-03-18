package diskstats

import (
	"io/ioutil"
	"time"

	"github.com/adgs85/gomonagent/agentlogger"
	"github.com/adgs85/gomonagent/agentmessagesdispatcher"
	"github.com/adgs85/gomonmarshalling/monmarshalling"

	"github.com/ricochet2200/go-disk-usage/du"
)

func StartDiskInfoLoopGoRoutine(sink agentmessagesdispatcher.StatSinkFuncType) {
	go func() {
		for {
			sink(collectDiskInfo(cfg.DiskFreeSpacePath, sink))
			time.Sleep(time.Duration(cfg.DiskPollingRateMs) * time.Millisecond)
		}
	}()
}

func collectDiskInfo(path string, sink agentmessagesdispatcher.StatSinkFuncType) monmarshalling.Stat {
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		agentlogger.Logger().Fatalln(err)
	}
	arr := []DiskStatPayload{}
	for _, file := range fileInfo {
		if file.IsDir() {
			fullPath := path + "/" + file.Name()
			dUsage := du.NewDiskUsage(fullPath)
			arr = append(arr, NewDiskStateEntry(dUsage.Size(), dUsage.Available(), fullPath))
		} else {
			agentlogger.Logger().Println("WARN files found in", path)
		}
	}

	return CreatePayload(NewDiskStatMetadata(cfg), arr)
}
