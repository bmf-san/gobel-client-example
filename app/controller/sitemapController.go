package controller

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"os"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/model"
	"github.com/bmf-san/gobel-client-example/app/presenter"
)

// A SitemapController is a controller for a sitemap.
type SitemapController struct {
	Logger    *logger.Logger
	Client    *api.Client
	Presenter *presenter.Presenter
}

// NewSitemapController creates a SitemapController.
func NewSitemapController(logger *logger.Logger, client *api.Client, presenter *presenter.Presenter) *SitemapController {
	return &SitemapController{
		Logger:    logger,
		Client:    client,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (si *SitemapController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// NOTE: Since api does not support getting all items, so taking a rough method.
		resp, err := si.Client.GetPosts(1, 99999)
		if err != nil {
			si.Logger.Error(err.Error())
			si.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			si.Logger.Error(err.Error())
			si.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var posts model.Posts
		if err = json.Unmarshal(body, &posts); err != nil {
			si.Logger.Error(err.Error())
			si.Logger.Error(string(body))
			si.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		// NOTE: Since api does not support getting all items, so taking a rough method.
		resp, err = si.Client.GetCategories(1, 99999)
		if err != nil {
			si.Logger.Error(err.Error())
			si.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			si.Logger.Error(err.Error())
			si.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var categories model.Categories
		if err = json.Unmarshal(body, &categories); err != nil {
			si.Logger.Error(err.Error())
			si.Logger.Error(string(body))
			si.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		// NOTE: Since api does not support getting all items, so taking a rough method.
		resp, err = si.Client.GetTags(1, 99999)
		if err != nil {
			si.Logger.Error(err.Error())
			si.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			si.Logger.Error(err.Error())
			si.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var tags model.Tags
		if err = json.Unmarshal(body, &tags); err != nil {
			si.Logger.Error(err.Error())
			si.Logger.Error(string(body))
			si.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var urlset model.URLSet
		urlset.Version = "http://www.sitemaps.org/schemas/sitemap/0.9"

		loc := os.Getenv("BASE_URL")
		for _, s := range [...]string{"/", "/posts", "/categories", "/tags", "/sitemap", "/feed"} {
			url := model.URL{
				Loc: loc + s,
			}
			urlset.URLs = append(urlset.URLs, url)
		}

		for _, p := range posts {
			url := model.URL{
				Loc: loc + "/posts/" + p.Title,
			}
			urlset.URLs = append(urlset.URLs, url)
		}

		for _, c := range categories {
			url := model.URL{
				Loc: loc + "/posts/categories/" + c.Name,
			}
			urlset.URLs = append(urlset.URLs, url)
		}

		for _, t := range tags {
			url := model.URL{
				Loc: loc + "/posts/tags/" + t.Name,
			}
			urlset.URLs = append(urlset.URLs, url)
		}

		buf, _ := xml.MarshalIndent(urlset, "", "  ")

		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte(xml.Header + string(buf)))
	})
}
