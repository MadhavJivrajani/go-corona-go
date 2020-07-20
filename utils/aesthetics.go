package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

// BeautifyIndexed prints a lsit of strings alphabetically and numerically indexed.
func ListEntities(items []string, commandType string) {
	sort.Strings(items)

	var states [][]string

	for index, item := range items {
		states = append(states, []string{strconv.Itoa(index + 1), item})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Sl no.", commandType})

	for _, v := range states {
		table.Append(v)
	}

	table.Render()
}

// StateInfo prints stats about a state leaving out details on a per-district basis.
func StateInfo(jsonArray []byte) error {
	unmarshalInto := make(map[string]json.RawMessage)

	err := json.Unmarshal(jsonArray, &unmarshalInto)

	if err != nil {
		return err
	}

	fmt.Printf("Info was last updated on: %s\n\n", FormatDate(
		TrimQuotes(unmarshalInto["lastupdatedtime"])),
	)

	stats := [][]string{
		[]string{"Active cases", TrimQuotes(unmarshalInto["active"])},
		[]string{"Confirmed cases", TrimQuotes(unmarshalInto["confirmed"])},
		[]string{"No. of deaths", TrimQuotes(unmarshalInto["deaths"])},
		[]string{"No. of recoveries (yay!)", TrimQuotes(unmarshalInto["recovered"])},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{TrimQuotes(unmarshalInto["state"])})

	for _, v := range stats {
		table.Append(v)
	}

	table.Render()
	return nil
}

// StateInfo prints stats about a district of a particular state.
func DistrictInfo(jsonArray []byte, entity string, lastUpdated string) error {
	unmarshalInto := make(map[string]json.RawMessage)

	err := json.Unmarshal(jsonArray, &unmarshalInto)

	if err != nil {
		return err
	}

	fmt.Printf("Info was last updated on: %s\n\n", lastUpdated)

	stats := [][]string{
		[]string{"Active cases", TrimQuotes(unmarshalInto["active"])},
		[]string{"Confirmed cases", TrimQuotes(unmarshalInto["confirmed"])},
		[]string{"No. of deaths", TrimQuotes(unmarshalInto["deceased"])},
		[]string{"No. of recoveries (yay!)", TrimQuotes(unmarshalInto["recovered"])},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{entity})

	for _, v := range stats {
		table.Append(v)
	}

	table.Render()
	return nil
}
