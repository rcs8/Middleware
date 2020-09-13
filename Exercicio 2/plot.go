package main

import (
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type protocolBarConfig struct {
	color  int
	offset vg.Length
}

const barWidth = vg.Centimeter

var protocolBar = map[string]protocolBarConfig{
	"UDP": {1, -barWidth / 2},
	"TCP": {2, barWidth / 2},
}

func Plot(results map[string][]float64, maxAvg float64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	for protocol, results := range results {
		barConfig := protocolBar[protocol]

		values := results[:]
		values = append(values, 0)
		bars, err := plotter.NewBarChart(plotter.Values(values), barWidth)
		if err != nil {
			panic(err)
		}
		bars.LineStyle.Width = vg.Length(0)
		bars.Color = plotutil.Color(barConfig.color)
		bars.Offset = barConfig.offset

		p.Add(bars)
		p.Legend.Add(protocol, bars)
	}

	p.Title.Text = "TCP vs UDP performance"
	p.X.Label.Text = "Number of clients"
	p.Y.Label.Text = "Time in μs"
	p.Y.Max = maxAvg * 1.3
	p.Legend.Top = true

	nominalClientAmounts := make([]string, 0)
	for _, clientAmount := range ClientAmounts {
		nominalClientAmounts = append(nominalClientAmounts, strconv.Itoa(clientAmount))
	}
	p.NominalX(nominalClientAmounts...)

	var hSize vg.Length = 2 * vg.Centimeter
	hSize += 2.5 * vg.Points(float64(len(ClientAmounts)+1)) * barWidth

	if err := p.Save(hSize, 4*vg.Inch, "performance.png"); err != nil {
		panic(err)
	}
}
