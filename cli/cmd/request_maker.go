package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"klutzer/conanical-library-app/shared"
	"net/http"
)

func makeRequest[R any](apiData shared.ApiRequest, apiResp *R, url string, method string, httpClient *http.Client) error {
	var body io.Reader
	if apiData != nil {
		if err := apiData.Validate(); err != nil {
			return err
		}

		b, err := json.Marshal(apiData)
		if err != nil {
			return err
		}

		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	// If api response is not nil then unmarshal the response body into it
	if resp != nil && resp.StatusCode == http.StatusOK {
		if apiResp != nil {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			json.Unmarshal(b, apiResp)
		}
	}

	// Close the api response body
	resp.Body.Close()

	return nil
}
