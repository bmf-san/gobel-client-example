package api

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const (
	// defaultPage is a default page.
	defaultPage = 1
	// defaultPage is a default limit.
	defaultLimit = 10
)

// Client represents the singular of api client.
type Client struct {
	httpClient *http.Client
	URL        *url.URL
}

// NewClient creates a Client.
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Duration(100 * time.Second),
		},
		URL: &url.URL{
			Scheme: os.Getenv("HTTP_API_SCHEME"),
			Host:   os.Getenv("HTTP_API_HOST"),
		},
	}
}

// Do sends an HTTP request and returns an HTTP response.
func (c *Client) Do(method string, path string, query map[string]string, data []byte) (*http.Response, error) {
	c.URL.Path = path
	req, err := http.NewRequest(method, c.URL.String(), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	// query
	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// header
	for key, value := range map[string]string{"key": "value"} {
		req.Header.Add(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}

	return resp, nil
}

// GetPageAndLimit gets page and limit.
func (c *Client) GetPageAndLimit(r *http.Request) (int, int, error) {
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
