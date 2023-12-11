package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

	// Unmarshal success responses in the passed apiResp
	if resp.StatusCode == http.StatusOK {
		if apiResp != nil {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			// Swallow error because we assume the api will always give us a valid response when there is a 200
			json.Unmarshal(b, apiResp)
		}

		// Handle error responses
	} else {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		apiErr := shared.ApiResponse[shared.AppError]{}
		if err := json.Unmarshal(b, &apiErr); err != nil {
			return errors.New(fmt.Sprintf("Failed to unmarshal error response for a %d", resp.StatusCode))
		}

		// Return app error from the api
		return errors.New(fmt.Sprintf("API Error: %s", apiErr.Data.Msg))

	}

	// Close the api response body
	resp.Body.Close()

	return nil
}
