package requests

import (
	"io"
	"net/http"
)

func Request(method string, url string, headers map[string]string, data io.Reader) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, data)

	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
