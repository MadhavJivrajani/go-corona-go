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

// districtCmd represents the district command
var districtCmd = &cobra.Command{
	Use:   "district",
	Short: "Provides district-wise stats of a particular state for Covid-19 in India",
	Long: `Syntax:
  go-corona-go district [state] [district]

Example:
  go-corona-go state Karnataka Bengaluru`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("district command requires exactly 2 arguments: [state] [district]")
		}
		return api.District(args)
	},
}

func init() {
	rootCmd.AddCommand(districtCmd)
}
