package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
)

func createDataFile(results map[string][]BenchResult) {
	f, err := os.Create("data/benchmark-" + os.Args[1] + ".dat")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines := make([][]string, len(ClientAmounts))
	for i := range ClientAmounts {
		lines[i] = make([]string, len(results)*4+1)
	}
	startIdx := 0
	for protocol, protocolResults := range results {
		for i, result := range protocolResults {
			lines[i][startIdx+0] = protocol
			lines[i][startIdx+1] = fmt.Sprintf("%.2f", result.mean)
			lines[i][startIdx+2] = fmt.Sprintf("%.2f", result.mean-result.sd)
			lines[i][startIdx+3] = fmt.Sprintf("%.2f", result.mean+result.sd)
		}
		startIdx += 4
	}

	for i, clientAmount := range ClientAmounts {
		line := strconv.Itoa(clientAmount)
		for _, val := range lines[i] {
			line += " " + val
		}

		fmt.Fprint(f, line+"\n")
	}
}

func gnuPlot(maxMean, minMeanSD, maxMeanSD float64) {
	maxMeanStr := strconv.Itoa(int(math.Ceil(maxMean * 1.3)))
	minMeanSDStr := strconv.Itoa(int(math.Floor(minMeanSD * 1.5)))
	maxMeanSDStr := strconv.Itoa(int(math.Ceil(maxMeanSD * 1.5)))

	exec.Command("gnuplot", "-c", "benchmark-script.txt", "0", maxMeanStr).Output()
	exec.Command("gnuplot", "-c", "benchmark-sd-script.txt", minMeanSDStr, maxMeanSDStr).Output()
}

func Plot(results map[string][]BenchResult, maxMean, minMeanSD, maxMeanSD float64) {
	createDataFile(results)
	gnuPlot(maxMean, minMeanSD, maxMeanSD)
}
