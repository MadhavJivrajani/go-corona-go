/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

// getStatesCmd represents the getStates command
var getStatesCmd = &cobra.Command{
	Use:   "states",
	Short: "Gets a list of valid states whose stats can be retreived",
	Long: `Syntax:
	go-corona-go getStates`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("The states command does not take any arguments")
		}
		return api.States()
	},
}

func init() {
	rootCmd.AddCommand(getStatesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getStatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getStatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
