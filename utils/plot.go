package utils

import (
	"fmt"

	ag "github.com/guptarohit/asciigraph"
)

type PlotData struct {
	XAxis   string
	YAxis   string
	Caption string
}

// PrintPlot generates a graph and prints it onto the terminal.
func PrintPlot(data []float64, plotData PlotData) {
	fmt.Println("Plot information:")
	fmt.Println("\tIMPORTANT: This plot is intended to show ONLY the trend of data and should not be interpreted accurately")
	fmt.Println("\n\tX Axis:", plotData.XAxis)
	fmt.Println("\tY Axis:", plotData.YAxis)

	graph := ag.Plot(data, ag.Caption(plotData.Caption), ag.Width(50), ag.Height(20))
	fmt.Println(graph)
}
