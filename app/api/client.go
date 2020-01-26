package api

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Client represents the singular of api client.
type Client struct {
	Client *http.Client
}

// NewClient creates a Client.
func NewClient() *Client {
	return &Client{
		Client: &http.Client{
			Timeout: time.Duration(100 * time.Second),
		},
	}
}

// HandleResponseBody handles an api client for getting response and body.
func (a *Client) HandleResponseBody(req *http.Request) (*http.Response, []byte, error) {
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp, body, nil
}

// HandleResponse handles an api client for getting response.
func (a *Client) HandleResponse(req *http.Request) (*http.Response, error) {
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp, nil
}

// GetPageAndLimit gets page and limit.
func (a *Client) GetPageAndLimit(r *http.Request, defaultPage int, defaultLimit int) (int, int, error) {
	var err error

	paramPage := r.URL.Query().Get("page")
	var page int
	if paramPage == "" {
		page = defaultPage
	} else {
		page, err = strconv.Atoi(paramPage)
		if err != nil {
			return 0, 0, err
		}
	}

	paramLimit := r.URL.Query().Get("limit")
	var limit int
	if paramLimit == "" {
		limit = defaultLimit
	} else {
		limit, err = strconv.Atoi(paramLimit)
		if err != nil {
			return 0, 0, err
		}
	}

	return page, limit, nil
}
