package controller

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"os"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/model"
	"github.com/bmf-san/gobel-client-example/app/response"
)

// A SitemapController is a controller for a sitemap.
type SitemapController struct {
	Logger   *logger.Logger
	Client   *api.Client
	Response *response.Response
}

// NewSitemapController creates a SitemapController.
func NewSitemapController(logger *logger.Logger, client *api.Client, response *response.Response) *SitemapController {
	return &SitemapController{
		Logger:   logger,
		Client:   client,
		Response: response,
	}
}

// Index displays a listing of the resource.
func (si *SitemapController) Index(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 99999

	_, body, err := si.Client.GetPosts(r, defaultPage, defaultLimit)
	if err != nil {
		si.Logger.Error(err.Error())
		si.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var posts model.Posts
	if err = json.Unmarshal(body, &posts); err != nil {
		si.Logger.Error(err.Error())
		si.Logger.Error(string(body))
		si.Response.Error(w, http.StatusInternalServerError)
		return
	}

	_, body, err = si.Client.GetCategories(r, defaultPage, defaultLimit)
	if err != nil {
		si.Logger.Error(err.Error())
		si.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var categories model.Categories
	if err = json.Unmarshal(body, &categories); err != nil {
		si.Logger.Error(err.Error())
		si.Logger.Error(string(body))
		si.Response.Error(w, http.StatusInternalServerError)
		return
	}

	_, body, err = si.Client.GetTags(r, defaultPage, defaultLimit)
	if err != nil {
		si.Logger.Error(err.Error())
		si.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var tags model.Tags
	if err = json.Unmarshal(body, &tags); err != nil {
		si.Logger.Error(err.Error())
		si.Logger.Error(string(body))
		si.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var urlset model.URLSet

	for _, s := range [...]string{"/", "/posts", "/categories", "/tags", "/sitemap", "/feed"} {
		url := model.URL{
			Loc: os.Getenv("SITE_URL") + s,
		}
		urlset.URLs = append(urlset.URLs, url)
	}

	for _, p := range posts {
		url := model.URL{
			Loc: os.Getenv("SITE_URL") + "/posts/" + p.Title,
		}
		urlset.URLs = append(urlset.URLs, url)
	}

	for _, c := range categories {
		url := model.URL{
			Loc: os.Getenv("SITE_URL") + "/posts/categories/" + c.Name,
		}
		urlset.URLs = append(urlset.URLs, url)
	}

	for _, t := range tags {
		url := model.URL{
			Loc: os.Getenv("SITE_URL") + "/posts/tags/" + t.Name,
		}
		urlset.URLs = append(urlset.URLs, url)
	}

	buf, _ := xml.MarshalIndent(urlset, "", "  ")

	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(xml.Header + string(buf)))
}
