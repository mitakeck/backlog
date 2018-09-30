package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// Client : Backlog HTTP Client
type Client struct {
	BaseURL         *url.URL
	HTTPClient      *http.Client
	APIRefreshToken string
	APIAccessToken  string
}

// NewClient : create new backlog http client
func NewClient(baseURL *url.URL, APIRefreshToken string, APIAccessToken string) (*Client, error) {
	client := &Client{
		BaseURL:         baseURL,
		APIRefreshToken: APIRefreshToken,
		APIAccessToken:  APIAccessToken,
	}

	return client, nil
}

// ComposeURL : compose url
func ComposeURL(base url.URL, pathStr string, params url.Values) string {
	copiedURL := base
	copiedURL.Path = path.Join(copiedURL.Path, pathStr)
	copiedURL.RawQuery = params.Encode()
	return copiedURL.String()
}

// Get http method
func (c *Client) Get(endPoint string, params url.Values) ([]byte, error) {
	return c.execute("GET", endPoint, params, true)
}

// Post http method
func (c *Client) Post(endPoint string, params url.Values) ([]byte, error) {
	return c.execute("POST", endPoint, params, true)
}

func (c *Client) execute(method, endpoint string, params url.Values, setAuth bool) ([]byte, error) {
	if c.HTTPClient == nil {
		c.HTTPClient = http.DefaultClient
	}

	copiedURL := *c.BaseURL
	copiedURL.Path = path.Join(copiedURL.Path, endpoint)
	url := copiedURL.String()

	var (
		req *http.Request
		err error
	)

	if method != "GET" {
		req, err = http.NewRequest(method, url, bytes.NewBufferString(params.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	if setAuth {
		req.Header.Set("Authorization", "Bearer "+c.APIAccessToken)
	}

	resp, err := c.HTTPClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return []byte(""), err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	if resp.StatusCode != 200 {
		return []byte(""), fmt.Errorf("response status = %v", resp.StatusCode)
	}

	return body, nil
}
