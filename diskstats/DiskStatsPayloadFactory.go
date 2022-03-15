package diskstats

import (
	"encoding/json"

	"github.com/adgs85/gomonagent/agentlogger"
	"github.com/adgs85/gomonmarshalling/monmarshalling"
)

const DiskSpaceStatType = "disk"

type DiskStatPayload struct {
	Size        uint64
	Available   uint64
	StoragePath string
}

func NewDiskStatMetadata(config DiskCollectorConfig) monmarshalling.MetaData {
	metaData := monmarshalling.NewStatsMetaDataWithTs()
	metaData.StatType = DiskSpaceStatType
	metaData.PollRateMs = config.DiskPollingRateMs
	metaData.HostName = config.HostName
	return *metaData
}

func NewDiskStateEntry(diskSize uint64, diskSpaceAvailable uint64, storagePath string) DiskStatPayload {
	return DiskStatPayload{
		diskSize,
		diskSpaceAvailable,
		storagePath,
	}
}

func CreatePayload(metaData monmarshalling.MetaData, statsArr []DiskStatPayload) monmarshalling.Stat {
	payload, err := json.Marshal(statsArr)
	if err != nil {
		agentlogger.Logger().Fatalln(err)
	}
	return monmarshalling.Stat{MetaData: metaData, Payload: string(payload)}
}
