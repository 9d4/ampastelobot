package httprequest

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

var AcceptedMethod []string = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
}

type HttpRequest struct {
	Url    string
	Method string
}

// check wheter method exists in AcceptedMethod or not
func (r *HttpRequest) checkMethod() error {
	for _, m := range AcceptedMethod {
		if m == r.Method {
			return nil
		}
	}

	return errors.New("invalid method")
}

func (r *HttpRequest) checkUrl() error {
	url, err := url.Parse(r.Url)
	if err != nil {
		return err
	}

	if url.Scheme == "" {
		url.Scheme = "http"
	}
	r.Url = url.String()

	return nil
}

// do the http request based on the
// *HttpRequest.Url
// *HttpRequest.Method
func (r *HttpRequest) do() (*http.Response, error) {
	var client *http.Client = &http.Client{}
	var response *http.Response

	if err := r.checkMethod(); err != nil {
		return response, err
	}

	if err := r.checkUrl(); err != nil {
		return response, err
	}

	req, _ := http.NewRequest(r.Method, r.Url, nil)

	response, err := client.Do(req)
	if err != nil {
		return response, err
	}

	return response, nil
}

// do the request
// returns the response body as string or error
func (r *HttpRequest) Do() (string, error) {
	res, err := r.do()
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// do simple request
// only returns the status code or error
func (r *HttpRequest) DoSimple() (int, error) {
	res, err := r.do()
	if err != nil {
		return 0, err
	}

	return res.StatusCode, nil
}

// Simple Request:
// make request with method HEAD
func NewSimpleRequest(url string) *HttpRequest {
	return &HttpRequest{
		Url:    url,
		Method: http.MethodHead,
	}
}

// create a new HttpRequest with custom method
func NewRequest(url string, method string) *HttpRequest {
	// capitalize method
	method = strings.ToUpper(method)

	return &HttpRequest{
		Url:    url,
		Method: method,
	}
}
