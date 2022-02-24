package httprequest

import (
	"errors"
	"io"
	"net/http"
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

func (r *HttpRequest) checkMethod() error {
	for _, m := range AcceptedMethod {
		if m == r.Method {
			return nil
		}
	}

	return errors.New("invalid method")
}

func (r *HttpRequest) do() (*http.Response, error) {
	var client *http.Client = &http.Client{}
	var response *http.Response

	if err := r.checkMethod(); err != nil {
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

func (r *HttpRequest) DoSimple() (int, error) {
	res, err := r.do()
	if err != nil {
		return 0, err
	}

	return res.StatusCode, nil
}

func NewSimpleRequest(url string) *HttpRequest {
	return &HttpRequest{
		Url:    url,
		Method: http.MethodHead,
	}
}
