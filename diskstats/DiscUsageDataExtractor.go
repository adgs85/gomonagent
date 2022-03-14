package diskstats

import (
	"io/ioutil"
	"log"

	"github.com/adgs85/gomonmarshalling/monmarshalling"

	"github.com/ricochet2200/go-disk-usage/du"
)

func CollectDiskInfo(path string, sink func(stat monmarshalling.Stat)) {
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}
	arr := []DiskStatPayload{}
	for _, file := range fileInfo {
		fullPath := path + "/" + file.Name()
		dUsage := du.NewDiskUsage(fullPath)
		arr = append(arr, NewDiskStateEntry(dUsage.Size(), dUsage.Available(), fullPath))
	}

	sink(CreatePayload(NewDiskStatMetadata(cfg), arr))

}
