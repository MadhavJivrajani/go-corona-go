package api

import (
	"encoding/json"

	"github.com/MadhavJivrajani/go-corona-go/utils"
)

type arrayJson []json.RawMessage

func PlotGeneral(key string) ([]float64, error) {
	body, err := ApiIndiaTimeline()
	if err != nil {
		var empty []float64
		return empty, err
	}

	var data arrayJson
	json.Unmarshal(body, &data)

	dataPoints, err := utils.ExtractDataPoints(data, key)
	if err != nil {
		var empty []float64
		return empty, err
	}

	return dataPoints, nil
}

// PlotConfirmed provides plotting functionality for confirmed cases per day.
func PlotConfirmed() error {
	dataPoints, err := PlotGeneral("dailyconfirmed")
	if err != nil {
		return err
	}

	plotDesc := utils.PlotData{
		XAxis:   "Dates starting from 30th January 2020.",
		YAxis:   "Number of confirmed Covid-19 cases in India.",
		Caption: "A plot showing the trend of the number of confirmed Covid-19 cases per day.",
	}

	utils.PrintPlot(dataPoints, plotDesc)
	return nil
}

// PlotDeaths provides plotting functionality for number of deaths per day.
func PlotDeaths() error {
	dataPoints, err := PlotGeneral("dailydeceased")
	if err != nil {
		return err
	}

	plotDesc := utils.PlotData{
		XAxis:   "Dates starting from 30th January 2020.",
		YAxis:   "Number of deaths due to Covid-19 cases in India.",
		Caption: "A plot showing the trend of the number of deaths due to Covid-19 per day.",
	}

	utils.PrintPlot(dataPoints, plotDesc)
	return nil
}

// PlotRecovered provides plotting functionality for number of recoveries from Covid-19 per day.
func PlotRecovered() error {
	dataPoints, err := PlotGeneral("dailyrecovered")
	if err != nil {
		return err
	}

	plotDesc := utils.PlotData{
		XAxis:   "Dates starting from 30th January 2020.",
		YAxis:   "Number of deaths due to Covid-19 cases in India.",
		Caption: "A plot showing the trend of the number of deaths due to Covid-19 per day.",
	}

	utils.PrintPlot(dataPoints, plotDesc)
	return nil
}
