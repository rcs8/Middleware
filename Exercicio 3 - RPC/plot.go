package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
)

func createDataFile(results map[string][]BenchResult) {
	f, err := os.Create("benchmark.dat")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines := make([][]string, len(ClientAmounts))
	for i := range ClientAmounts {
		lines[i] = make([]string, 9)
	}
	for protocol, protocolResults := range results {
		startIdx := 1
		if protocol == "UDP" {
			startIdx = 5
		}
		for i, result := range protocolResults {
			lines[i][startIdx+0] = protocol
			lines[i][startIdx+1] = fmt.Sprintf("%.2f", result.mean)
			lines[i][startIdx+2] = fmt.Sprintf("%.2f", result.mean-result.sd)
			lines[i][startIdx+3] = fmt.Sprintf("%.2f", result.mean+result.sd)
		}
	}

	for i, clientAmount := range ClientAmounts {
		fmt.Fprintf(
			f, "%d %s %s %s %s %s %s %s %s\n",
			clientAmount,
			lines[i][1],
			lines[i][2],
			lines[i][3],
			lines[i][4],
			lines[i][5],
			lines[i][6],
			lines[i][7],
			lines[i][8],
		)
	}
}

func gnuPlot(maxMean, minMeanSD, maxMeanSD float64) {
	maxMeanStr := strconv.Itoa(int(math.Ceil(maxMean * 1.3)))
	minMeanSDStr := strconv.Itoa(int(math.Floor(minMeanSD * 1.5)))
	maxMeanSDStr := strconv.Itoa(int(math.Ceil(maxMeanSD * 1.5)))

	exec.Command("./gnuplot", "-c", "benchmark-script.txt", "0", maxMeanStr).Output()
	exec.Command("./gnuplot", "-c", "benchmark-sd-script.txt", minMeanSDStr, maxMeanSDStr).Output()
}

func Plot(results map[string][]BenchResult, maxMean, minMeanSD, maxMeanSD float64) {
	createDataFile(results)
	gnuPlot(maxMean, minMeanSD, maxMeanSD)
}
