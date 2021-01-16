package api

import (
	"net/http"
	"strconv"
)

const (
	// getCategoriesPath is a path for getting categories.
	getCategoriesPath = "/categories"
)

// GetCategories requests categories.
func (c *Client) GetCategories(page int, limit int) (*http.Response, error) {
	resp, err := c.Do(http.MethodGet, getCategoriesPath, map[string]string{"page": strconv.Itoa(page), "limit": strconv.Itoa(limit)}, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
