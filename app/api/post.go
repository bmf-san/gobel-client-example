package api

import (
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/bmf-san/goblin"
)

// GetPosts requests posts
func (a *Client) GetPosts(r *http.Request, defaultPage int, defaultLimit int) (*http.Response, []byte, error) {
	page, limit, err := a.GetPageAndLimit(r, defaultPage, defaultLimit)
	if err != nil {
		return nil, nil, err
	}

	values := url.Values{}
	values.Add("page", strconv.Itoa(page))
	values.Add("limit", strconv.Itoa(limit))

	req, err := http.NewRequest(http.MethodGet, os.Getenv("API_URL")+"/posts?", nil)
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

// GetPostsByCategory requests posts by category.
func (a *Client) GetPostsByCategory(r *http.Request, defaultPage int, defaultLimit int) (*http.Response, []byte, error) {
	page, limit, err := a.GetPageAndLimit(r, defaultPage, defaultLimit)
	if err != nil {
		return nil, nil, err
	}

	values := url.Values{}
	values.Add("page", strconv.Itoa(page))
	values.Add("limit", strconv.Itoa(limit))

	name := goblin.GetParam(r.Context(), "name")

	req, err := http.NewRequest(http.MethodGet, os.Getenv("API_URL")+"/posts/categories/"+name+"?", nil)
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

// GetPostsByTag requests posts by tag.
func (a *Client) GetPostsByTag(r *http.Request, defaultPage int, defaultLimit int) (*http.Response, []byte, error) {
	page, limit, err := a.GetPageAndLimit(r, defaultPage, defaultLimit)
	if err != nil {
		return nil, nil, err
	}

	values := url.Values{}
	values.Add("page", strconv.Itoa(page))
	values.Add("limit", strconv.Itoa(limit))

	name := goblin.GetParam(r.Context(), "name")

	req, err := http.NewRequest(http.MethodGet, os.Getenv("API_URL")+"/posts/tags/"+name+"?", nil)
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

// GetPost requests a post.
func (a *Client) GetPost(r *http.Request) (*http.Response, []byte, error) {
	title := goblin.GetParam(r.Context(), "title")

	req, err := http.NewRequest(http.MethodGet, os.Getenv("API_URL")+"/posts/"+title, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, body, err := a.HandleResponseBody(req)
	if err != nil {
		return nil, nil, err
	}

	return resp, body, err
}
