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

// A FeedController is a controller for feed.
type FeedController struct {
	Logger    *logger.Logger
	Client    *api.Client
	Presenter *presenter.Presenter
}

// NewFeedController creates a FeedController.
func NewFeedController(logger *logger.Logger, client *api.Client, presenter *presenter.Presenter) *FeedController {
	return &FeedController{
		Logger:    logger,
		Client:    client,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (fc *FeedController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// NOTE: Since api does not support getting all items, so taking a rough method.
		resp, err := fc.Client.GetPosts(1, 99999)
		if err != nil {
			fc.Logger.Error(err.Error())
			fc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fc.Logger.Error(err.Error())
			fc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var posts model.Posts

		if err := json.Unmarshal(body, &posts); err != nil {
			fc.Logger.Error(err.Error())
			fc.Logger.Error(string(body))
			fc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var entries []model.FeedEntry
		url := os.Getenv("BASE_URL")
		for _, p := range posts {
			u := url + "/posts/" + p.Title
			entry := model.FeedEntry{
				Title: p.Title,
				Link: model.FeedLink{
					Href: u,
				},
				ID:        u,
				Updated:   p.UpdatedAt,
				Published: p.CreatedAt,
				Author: model.FeedAuthor{
					Name: p.Admin.Name,
				},
				Content: model.FeedContent{
					Type:  "html",
					CDATA: p.HTMLBody,
				},
			}
			entries = append(entries, entry)
		}

		feed := model.Feed{
			Title:    "Gobel",
			Subtitle: "Gobel is a headless cms built with golang.",
			Link: model.FeedLink{
				Href: url,
			},
			Updated: posts[len(posts)-1].UpdatedAt,
			Author: model.FeedAuthor{
				Name: "bmf_san",
			},
			ID:      url,
			Entries: entries,
		}

		buf, _ := xml.MarshalIndent(feed, "", "  ")

		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte(xml.Header + string(buf)))
	})
}
