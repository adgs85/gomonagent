package diskstats

import (
	"io/ioutil"

	"github.com/adgs85/gomonagent/agentlogger"
	"github.com/adgs85/gomonagent/agentmessagesdispatcher"

	"github.com/ricochet2200/go-disk-usage/du"
)

func CollectDiskInfo(path string, sink agentmessagesdispatcher.StatSinkFuncType) {
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		agentlogger.Logger().Fatalln(err)
	}
	arr := []DiskStatPayload{}
	for _, file := range fileInfo {
		fullPath := path + "/" + file.Name()
		dUsage := du.NewDiskUsage(fullPath)
		arr = append(arr, NewDiskStateEntry(dUsage.Size(), dUsage.Available(), fullPath))
	}

	sink(CreatePayload(NewDiskStatMetadata(cfg), arr))
}
