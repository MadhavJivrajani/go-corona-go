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

func trimQuotes(toTrim []byte) string {
	return string(toTrim[1 : len(toTrim)-1])
}

func formatDate(date string) string {
	var formatted []byte

	for _, char := range date {
		if char != '\\' {
			formatted = append(formatted, byte(char))
		}
	}
	return string(formatted)
}

// StateInfo prints stats about a state leaving out details on a per-district basis.
func StateInfo(jsonArray []byte) error {
	unmarshalInto := make(map[string]json.RawMessage)

	err := json.Unmarshal(jsonArray, &unmarshalInto)

	if err != nil {
		return err
	}

	fmt.Printf("Info was last updated on: %s\n\n", formatDate(
		trimQuotes(unmarshalInto["lastupdatedtime"])),
	)

	stats := [][]string{
		[]string{"Active cases", trimQuotes(unmarshalInto["active"])},
		[]string{"Confirmed cases", trimQuotes(unmarshalInto["confirmed"])},
		[]string{"No. of deaths", trimQuotes(unmarshalInto["deaths"])},
		[]string{"No. of confirmed deaths", trimQuotes(unmarshalInto["deaths"])},
		[]string{"No. of recoveries (yay!)", trimQuotes(unmarshalInto["recovered"])},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{trimQuotes(unmarshalInto["state"])})

	for _, v := range stats {
		table.Append(v)
	}

	table.Render()
	return nil
}
