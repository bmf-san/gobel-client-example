package api

import (
	"net/http"
	"net/url"
	"os"
	"strconv"
)

// GetTags requests categories
func (a *Client) GetTags(r *http.Request, defaultPage int, defaultLimit int) (*http.Response, []byte, error) {
	page, limit, err := a.GetPageAndLimit(r, defaultPage, defaultLimit)
	if err != nil {
		return nil, nil, err
	}

	values := url.Values{}
	values.Add("page", strconv.Itoa(page))
	values.Add("limit", strconv.Itoa(limit))

	req, err := http.NewRequest(http.MethodGet, os.Getenv("API_URL")+"/tags?", nil)
	if err != nil {
		return nil, nil, err
	}
	req.URL.RawQuery = values.Encode()

	resp, body, err := a.HandleResponseBody(req)
	if err != nil {
		return nil, nil, err
	}

	return resp, body, nil
}
