package waniapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func BuildGetRequest(resource string, parameters map[string]string) (*http.Request, error) {
	req, err := http.NewRequest("GET", waikaniURL+"/"+resource, nil)
	if err != nil {
		return nil, err
	}

	addAuth(req)

	values := req.URL.Query()
	for name, value := range parameters {
		values.Add(name, value)
	}
	req.URL.RawQuery = values.Encode()

	return req, nil
}

func RequestNextPage(pageURL string) (*http.Request, error) {
	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return nil, err
	}

	addAuth(req)
	return req, nil
}

func DoAndDecode(client *http.Client, req *http.Request, colention interface{}, subjetType string) error {
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[Request%s] Do request fail %w", subjetType, err)
	}

	if resp.StatusCode == http.StatusOK {
		err = decodeColection(resp.Body, colention)
		if err != nil {
			return fmt.Errorf("[Request%s] decode body fail %w", subjetType, err)
		}
		return nil
	} else {

		return fmt.Errorf("[Request%s] got code %s \non url: %s", subjetType, resp.Status, req.URL)
	}
}

func addAuth(req *http.Request) {
	req.Header.Add("Authorization", "Bearer d95f8d87-4b01-4326-bcc9-95347b5ddc29")
}

func decodeColection(reader io.ReadCloser, colection interface{}) error {
	defer reader.Close()
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(data), colection)
	if err != nil {
		return err
	}
	return nil
}
