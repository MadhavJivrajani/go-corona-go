package api

import (
	"fmt"
	"strings"

	"github.com/MadhavJivrajani/go-corona-go/utils"
)

// District provides district-wise stats of a particular state for Covid-19 in India.
func District(args []string) error {
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

	district := strings.ToLower(args[1])
	districts, err := utils.GetKeys(districtJSON)
	if err != nil {
		return err
	}

	districtName, exists := utils.ExistsString(districts, district)
	if !exists {
		return fmt.Errorf("entered district is not supported, try the districts command to get a list of the valid districts")
	}

	districtStats, err := utils.GetSubJSON(districtJSON, districtName)
	if err != nil {
		return err
	}

	lastUpdated, err := utils.GetSubJSON(reqState, "lastupdatedtime")
	if err != nil {
		return err
	}

	lastUpdatedString := utils.FormatDate(utils.TrimQuotes(lastUpdated))
	err = utils.DistrictInfo(districtStats, districtName, lastUpdatedString)
	if err != nil {
		return err
	}

	return nil
}
