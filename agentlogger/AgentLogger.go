package agentlogger

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "agent", log.Default().Flags()|log.LUTC|log.Lshortfile)

func Logger() *log.Logger {
	return logger
}
