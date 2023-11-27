package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func makeRequest(apiData interface{}, url string, method string, httpClient *http.Client) error {
	b, err := json.Marshal(apiData)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()

	return nil
}
