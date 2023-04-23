package api

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	// getPostsPath is a path for getting posts.
	getPostsPath = "/posts"
	// getPostsBySearchPath is a path for getting posts by search.
	getPostsBySearchPath = "/posts/search"
	// getPostsByCategoryPath is a path for getting posts by category.
	getPostsByCategoryPath = "/posts/categories/%s"
	// getPostsByTagPath is a path for getting posts by tag.
	getPostsByTagPath = "/posts/tags/%s"
	// getPostPath is a path for getting a post.
	getPostPath = "/posts/%s"
)

// GetPosts requests posts
func (c *Client) GetPosts(page int, limit int) (*http.Response, error) {
	resp, err := c.Do(http.MethodGet, getPostsPath, map[string]string{"page": strconv.Itoa(page), "limit": strconv.Itoa(limit)}, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetPostsByKeyword requests posts by keyword.
func (c *Client) GetPostsByKeyword(keyword string, page int, limit int) (*http.Response, error) {
	resp, err := c.Do(http.MethodGet, getPostsBySearchPath, map[string]string{"keyword": keyword, "page": strconv.Itoa(page), "limit": strconv.Itoa(limit)}, nil)
	fmt.Printf("%#v\n", resp.Request.URL)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetPostsByCategory requests posts by category.
func (c *Client) GetPostsByCategory(name string, page int, limit int) (*http.Response, error) {
	resp, err := c.Do(http.MethodGet, fmt.Sprintf(getPostsByCategoryPath, name), map[string]string{"page": strconv.Itoa(page), "limit": strconv.Itoa(limit)}, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetPostsByTag requests posts by tag.
func (c *Client) GetPostsByTag(name string, page int, limit int) (*http.Response, error) {
	resp, err := c.Do(http.MethodGet, fmt.Sprintf(getPostsByTagPath, name), map[string]string{"page": strconv.Itoa(page), "limit": strconv.Itoa(limit)}, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetPost requests a post.
func (c *Client) GetPost(title string) (*http.Response, error) {
	resp, err := c.Do(http.MethodGet, fmt.Sprintf(getPostPath, title), nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
