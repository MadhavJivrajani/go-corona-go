package api

import (
	"encoding/json"

	"github.com/MadhavJivrajani/go-corona-go/utils"
)

type arrayJson []json.RawMessage

// PlotConfirmed provides plotting functionality for confirmed cases per day.
func PlotConfirmed() error {
	body, err := ApiIndiaTimeline()
	if err != nil {
		return err
	}

	var data arrayJson
	json.Unmarshal(body, &data)

	dataPoints, err := utils.ExtractDataPoints(data, "dailyconfirmed")
	if err != nil {
		return err
	}

	plotDesc := utils.PlotData{
		XAxis:   "Dates starting from 30th January 2020.",
		YAxis:   "Number of confirmed Covid-19 cases in India.",
		Caption: "A plot showing the trend of the number of confirmed Covid-19 cases per day.",
	}

	utils.PrintPlot(dataPoints, plotDesc)
	//fmt.Println(dataPoints)
	return nil
}
