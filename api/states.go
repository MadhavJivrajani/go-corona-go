package api

import (
	"github.com/MadhavJivrajani/go-corona-go/utils"
)

// GetStates returns a slice containing a list of valid states.
func GetStates() error {
	body, err := ApiIndia()

	stateJSON, err := utils.GetSubJSON(body, "state_wise")
	if err != nil {
		return err
	}

	states, err := utils.GetKeys(stateJSON)
	if err != nil {
		return err
	}

	utils.ListEntities(states, "State")
	return nil
}
