package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Client interface {
	TestApi() (string, error)
}

type client struct {
	plugin     *Plugin
	HTTPClient *http.Client
}

func (azureDevops *client) TestApi() (string, error) {

	return "hello world", nil
}

// Wrapper to make REST API requests with "application/json" type content
func (azureDevops *client) callJSON(url, path, method string, in, out interface{}) (responseData []byte, err error) {
	contentType := "application/json"
	buf := &bytes.Buffer{}
	err = json.NewEncoder(buf).Encode(in)
	if err != nil {
		return nil, err
	}
	return azureDevops.call(url, method, path, contentType, buf, out)
}

// Makes HTTP request to REST APIs
func (azureDevops *client) call(basePath, method, path, contentType string, inBody io.Reader, out interface{}) (responseData []byte, err error) {
	errContext := fmt.Sprintf("Azure Devops: Call failed: method:%s, path:%s", method, path)
	pathURL, err := url.Parse(path)
	if err != nil {
		return nil, errors.WithMessage(err, errContext)
	}

	if pathURL.Scheme == "" || pathURL.Host == "" {
		var baseURL *url.URL
		baseURL, err = url.Parse(basePath)
		if err != nil {
			return nil, errors.WithMessage(err, errContext)
		}
		if path[0] != '/' {
			path = "/" + path
		}
		path = baseURL.String() + path
	}

	req, err := http.NewRequest(method, path, inBody)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}

	resp, err := azureDevops.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()

	responseData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated:
		if out != nil {
			err = json.Unmarshal(responseData, out)
			if err != nil {
				return responseData, err
			}
		}
		return responseData, nil

	case http.StatusNoContent:
		return nil, nil

	case http.StatusNotFound:
		return nil, errors.Errorf("not found")
	}

	type ErrorResponse struct {
		Message string `json:"message"`
	}
	errResp := ErrorResponse{}
	err = json.Unmarshal(responseData, &errResp)
	if err != nil {
		return responseData, errors.WithMessagef(err, "status: %s", resp.Status)
	}
	return responseData, fmt.Errorf("errorMessage %s", errResp.Message)
}

func InitClient(p *Plugin) Client {
	return &client{
		plugin:     p,
		HTTPClient: &http.Client{},
	}
}
