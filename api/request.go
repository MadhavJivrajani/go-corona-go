package api

import (
	"io/ioutil"
	"net/http"
)

func ApiIndia() ([]byte, error) {
	url := "https://corona-virus-world-and-india-data.p.rapidapi.com/api_india"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		var empty []byte
		return empty, err
	}

	req.Header.Add("x-rapidapi-host", "corona-virus-world-and-india-data.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "39cadffecemshb8c6b64846a679dp1c212bjsn6386a8140e6e")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		var empty []byte
		return empty, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		var empty []byte
		return empty, err
	}

	return body, nil
}
