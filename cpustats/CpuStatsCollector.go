package cpustats

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/adgs85/gomonagent/agentmessagesdispatcher"
)

func startCpuUsageInfoLoop(sink agentmessagesdispatcher.StatSinkFuncType) {
	go func() {
		collectTotalCpuUsage(sink)
	}()
}

func collectTotalCpuUsage(sink agentmessagesdispatcher.StatSinkFuncType) {

	file, err := os.Open(cpuStatCfg.CpuStatsPath)
	if err != nil {
		logger.Fatalln(err)
	}
	defer file.Close()

	var prevIdleTime, prevTotalTime uint64
	for {
		if err != nil {
			logger.Fatal(err)
		}

		totalTime, idleTime := calculateTotalAndIdleTime(getFirstLine(file))
		if prevTotalTime > 0 {
			deltaIdleTime := idleTime - prevIdleTime
			deltaTotalTime := totalTime - prevTotalTime
			cpuUsage := (1.0 - float64(deltaIdleTime)/float64(deltaTotalTime)) * 100.0
			sink(*newCpuUsageStat("cpu", cpuUsage))
		}
		prevIdleTime = idleTime
		prevTotalTime = totalTime
		_, err := file.Seek(0, 0)

		if err != nil {
			logger.Fatalln(err)
		}

		time.Sleep(time.Second)
	}

}

func getFirstLine(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	firstLine := scanner.Text()[5:] // get first line (aggregated cores time) + skip "cpu  " from /proc/stat
	if err := scanner.Err(); err != nil {
		logger.Fatal(err)
	}
	return firstLine
}

func calculateTotalAndIdleTime(cpuLine string) (uint64, uint64) {
	split := strings.Fields(cpuLine)
	idleTime, _ := strconv.ParseUint(split[3], 10, 64)
	totalTime := uint64(0)
	for _, s := range split {
		u, _ := strconv.ParseUint(s, 10, 64)
		totalTime += u
	}

	return totalTime, idleTime
}
