package api

import (
	"fmt"
	"strings"

	"github.com/MadhavJivrajani/go-corona-go/utils"
)

// State prints stats about Covid-19 for a particular state in India.
func State(args []string) error {
	body, err := ApiIndia()

	subJSON, err := utils.GetSubJSON(body, "state_wise")
	if err != nil {
		return err
	}

	states, err := utils.GetKeys(subJSON)
	if err != nil {
		return err
	}

	state := strings.ToLower(args[0])
	nameInJSON, exists := utils.ExistsString(states, state)

	if !exists {
		return fmt.Errorf("requested state is not supported, try the states command to get a list of the valid states")
	}

	stateStats, err := utils.GetSubJSON(subJSON, nameInJSON)
	if err != nil {
		return err
	}

	err = utils.StateInfo(stateStats)
	if err != nil {
		return err
	}

	return nil
}
