package api

import (
	"io/ioutil"
	"net/http"
)

//ApiIndia queries the API and returns the response body containing stats about Covid-19 in India.
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

// ApiIndiaTimeline queries the timeline feature of the API and returns per-day statistics.
func ApiIndiaTimeline() ([]byte, error) {

	url := "https://corona-virus-world-and-india-data.p.rapidapi.com/api_india_timeline"

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
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}
