/*
Copyright © 2020 Madhav Jivrajani madhav.jiv@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/MadhavJivrajani/go-corona-go/api"
	"github.com/spf13/cobra"
)

// plotCmd represents the plot command
var plotCmd = &cobra.Command{
	Use:   "plot",
	Short: "Plots daily data about a particular statistic about Covid-19 in India",
	Long: `Syntax:
	go-corona-go plot -t [type]
	[type]: confirmed (default), deaths, recovered

	confirmed: plots trend of number of confirmed cases of Covid-19 in India each day
	deceased : plots trend of number of deaths in India due to Covid-19 each day
	recovered: plots trend of number of recoveries from Covid-19 in India each day

Example:
	go-corona-go plot -t deceased`,
	RunE: func(cmd *cobra.Command, args []string) error {
		plotType, _ := cmd.Flags().GetString("type")

		if plotType != "confirmed" && plotType != "deaths" && plotType != "recovered" {
			return fmt.Errorf("type of plot specified is not valid, see go-corona-go plot --help")
		}

		switch plotType {
		case "confirmed":
			return api.PlotConfirmed()
		case "deceased":
			//return api.PlotDeaths()
		case "recovered":
			//return api.PlotRecovered()
		}
		fmt.Println("plot called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(plotCmd)
	plotCmd.Flags().StringP("type", "t", "confirmed", "type of data to plot")
}
