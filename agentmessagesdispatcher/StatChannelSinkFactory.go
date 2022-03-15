package agentmessagesdispatcher

import (
	"time"

	"github.com/adgs85/gomonagent/agentlogger"
	"github.com/adgs85/gomonmarshalling/monmarshalling"
	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
)

const senderStallMsKey = "sender_stall_ms"

var senderStallTimeoutDuration = envconfig.GetViperConfig().GetDuration(senderStallMsKey) * time.Millisecond

func NewChannelSinkStallLogging(c chan monmarshalling.Stat) StatSinkFuncType {

	return func(stat monmarshalling.Stat) {
		//TODO perhaps add max retries and throw away the stat
		for !addToChannelOrWarn(c, stat) {
		}
	}
}

func addToChannelOrWarn(c chan monmarshalling.Stat, object monmarshalling.Stat) bool {
	for {
		select {
		case c <- object:
			return true
		case <-time.After(senderStallTimeoutDuration):
			agentlogger.Logger().Println("WARN timed out while publishing to channel")
			return false
		}
	}
}
