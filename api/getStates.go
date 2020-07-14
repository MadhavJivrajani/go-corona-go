package api

import (
	"io/ioutil"
	"net/http"

	"github.com/MadhavJivrajani/go-corona-go/utils"
)

// GetStates returns a slice containing a list of valid states.
func GetStates() error {
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

	keys, err := utils.GetKeys(subJSON)
	if err != nil {
		return err
	}

	utils.Beautify(keys)
	return nil
}
