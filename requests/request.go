package requests

import (
	"io"
	"net/http"
)

type Response struct {
	Headers    http.Header
	Content    []byte
	StatusCode int
	request    *http.Request
	response   *http.Response
}

func NewResponse(headers http.Header, content []byte, status_code int, request *http.Request, response *http.Response) *Response {
	return &Response{headers, content, status_code, request, response}
}

func Request(method string, url string, headers map[string]string, data io.Reader) (Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest(method, url, data)

	if err != nil {
		return Response{}, err
	}
	for header, value := range headers {
		req.Header.Add(header, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return Response{nil, nil, 0, resp.Request, resp}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return *NewResponse(resp.Header, nil, resp.StatusCode, resp.Request, resp), err
	}

	return *NewResponse(resp.Header, body, resp.StatusCode, resp.Request, resp), nil
}

func Get(url string, headers map[string]string) (Response, error) {
	return Request("GET", url, headers, nil)
}

func Post(url string, headers map[string]string, data io.Reader) (Response, error) {
	return Request("POST", url, headers, data)
}

func Put(url string, headers map[string]string, data io.Reader) (Response, error) {
	return Request("PUT", url, headers, data)
}
