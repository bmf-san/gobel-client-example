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

// A FeedController is a controller for feed.
type FeedController struct {
	Logger   *logger.Logger
	Client   *api.Client
	Response *response.Response
}

// NewFeedController creates a FeedController.
func NewFeedController(logger *logger.Logger, client *api.Client, response *response.Response) *FeedController {
	return &FeedController{
		Logger:   logger,
		Client:   client,
		Response: response,
	}
}

// Index displays a listing of the resource.
func (fc *FeedController) Index(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 99999

	_, body, err := fc.Client.GetPosts(r, defaultPage, defaultLimit)
	if err != nil {
		fc.Logger.Error(err.Error())
		fc.Response.Error(w, http.StatusInternalServerError)
		return
	}
	var posts model.Posts

	if err := json.Unmarshal(body, &posts); err != nil {
		fc.Logger.Error(err.Error())
		fc.Logger.Error(string(body))
		fc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var entries []model.FeedEntry
	for _, p := range posts {
		entry := model.FeedEntry{
			Title: p.Title,
			Link: model.FeedLink{
				Href: os.Getenv("SITE_URL") + "/posts/" + p.Title,
			},
			ID:        os.Getenv("SITE_URL") + "/posts/" + p.Title,
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
			Href: os.Getenv("SITE_URL"),
		},
		Updated: posts[len(posts)-1].UpdatedAt,
		Author: model.FeedAuthor{
			Name: "bmf_san",
		},
		ID:      os.Getenv("SITE_URL"),
		Entries: entries,
	}

	buf, _ := xml.MarshalIndent(feed, "", "  ")

	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(xml.Header + string(buf)))
}
