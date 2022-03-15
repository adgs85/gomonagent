package agentmessagesdispatcher

import (
	"time"

	"github.com/adgs85/gomonagent/agentlogger"
	"github.com/adgs85/gomonmarshalling/monmarshalling"
	"github.com/adgs85/gomonmarshalling/monmarshalling/envconfig"
)

const senderStallMsKey = "sender_stall_ms"

var senderStallTimeoutDuration = envconfig.GetViperConfig().GetDuration(senderStallMsKey) * time.Millisecond

func NewChannelSink(c chan monmarshalling.Stat) StatSinkFuncType {

	return func(stat monmarshalling.Stat) {
	out:
		for {
			select {
			case c <- stat:
				break out
			case <-time.After(senderStallTimeoutDuration):
				agentlogger.Logger().Println("WARN timed out while publishing to channel")
			}
		}
	}
}
