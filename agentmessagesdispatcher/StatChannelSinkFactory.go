package agentmessagesdispatcher

import (
	"time"

	"github.com/adgs85/gomonagent/agentconfiguration"
	"github.com/adgs85/gomonagent/agentlogger"
	"github.com/adgs85/gomonmarshalling/monmarshalling"
)

func NewChannelSinkStallLogging(c chan monmarshalling.Stat) StatSinkFuncType {
	senderStatStallDuration := time.Millisecond * time.Duration(agentconfiguration.GlobalCfg().SenderStatStallMs)
	return func(stat monmarshalling.Stat) {
		//TODO perhaps add max retries and throw away the stat
		for !addToChannelOrWarn(c, stat, senderStatStallDuration) {
		}
	}
}

func addToChannelOrWarn(c chan monmarshalling.Stat, object monmarshalling.Stat, senderStatStallMs time.Duration) bool {
	for {
		select {
		case c <- object:
			return true
		case <-time.After(senderStatStallMs):
			agentlogger.Logger().Println("WARN timed out while publishing to channel")
			return false
		}
	}
}
