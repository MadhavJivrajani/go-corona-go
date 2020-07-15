package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/MadhavJivrajani/go-corona-go/utils"
)

// State prints stats about Covid-19 for a particular state in India.
func State(args []string) error {
	url := "https://corona-virus-world-and-india-data.p.rapidapi.com/api_india"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	req.Header.Add("x-rapidapi-host", "corona-virus-world-and-india-data.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "39cadffecemshb8c6b64846a679dp1c212bjsn6386a8140e6e")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

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
		return fmt.Errorf("requested state is not supported, try the getStates command to get a list of the valid states")
	}

	stateStats, err := utils.GetSubJSON(subJSON, nameInJSON)
	if err != nil {
		return err
	}

	fmt.Println(string(stateStats))
	return nil
}
