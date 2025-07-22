package cloudbase

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	domain  = "api.tcloudbasegateway.com"
	timeout = time.Second * 60
)

// client
type Client struct {
	EnvId    string `json:"envId"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewClient
func NewClient(username string, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

// httpPost Send HTTP Post Request
func httpPost(url string, headers map[string]string, reqBody interface{}, respBody interface{}) error {
	reqData, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqData))
	if err != nil {
		return err
	}
	// Add Request Headers
	addHeaders(req, headers)

	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, respBody)
	if err != nil {
		return err
	}
	return nil
}

// httpGet Send HTTP GET Request
func httpGet(url string, headers map[string]string, respBody interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	// Add Request Headers
	addHeaders(req, headers)

	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 200 && resp.StatusCode <= 204 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, respBody)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

// addHeaders
func addHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
}
