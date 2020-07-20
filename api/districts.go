package api

import (
	"fmt"
	"strings"

	"github.com/MadhavJivrajani/go-corona-go/utils"
)

// GetStates returns a slice containing a list of valid states.
func GetDistricts(args []string) error {
	body, err := ApiIndia()

	stateWise, err := utils.GetSubJSON(body, "state_wise")
	if err != nil {
		return err
	}

	state := strings.ToLower(args[0])
	states, err := utils.GetKeys(stateWise)
	if err != nil {
		return err
	}
	nameInJSON, exists := utils.ExistsString(states, state)

	if !exists {
		return fmt.Errorf("entered state is not supported, try the states command to get a list of the valid states")
	}

	reqState, err := utils.GetSubJSON(stateWise, nameInJSON)
	if err != nil {
		return err
	}

	districtJSON, err := utils.GetSubJSON(reqState, "district")
	if err != nil {
		return err
	}

	districts, err := utils.GetKeys(districtJSON)
	if err != nil {
		return err
	}

	utils.ListEntities(districts, "District")
	return nil
}
