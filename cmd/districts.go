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

// getDistrictsCmd represents the getDistricts command
var getDistrictsCmd = &cobra.Command{
	Use:   "districts",
	Short: "Gets a list of valid districts of a particular state whose stats can be retreived",
	Long: `Syntax:
	go-corona-go getDistricts [state]
	
Example:
	go-corona-go getDistricts Karnataka`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("districts command takes exactly one argument")
		}
		return api.GetDistricts(args)
	},
}

func init() {
	rootCmd.AddCommand(getDistrictsCmd)
}
