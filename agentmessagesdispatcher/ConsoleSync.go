package agentmessagesdispatcher

import (
	"fmt"

	"github.com/adgs85/gomonmarshalling/monmarshalling"
	"github.com/davecgh/go-spew/spew"
)

func SpewToConsoleSink(stats monmarshalling.Stat) {
	fmt.Println(spew.Sdump(stats))
}
