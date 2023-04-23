package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/model"
	"github.com/bmf-san/gobel-client-example/app/presenter"
)

// A HomeController is a controller for a home.
type HomeController struct {
	Logger    *logger.Logger
	Client    *api.Client
	Presenter *presenter.Presenter
}

// NewHomeController creates a HomeController.
func NewHomeController(logger *logger.Logger, client *api.Client, presenter *presenter.Presenter) *HomeController {
	return &HomeController{
		Logger:    logger,
		Client:    client,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (hc *HomeController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, limit, err := hc.Client.GetPageAndLimit(r)
		if err != nil {
			hc.Logger.Error(err.Error())
			hc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		resp, err := hc.Client.GetPosts(page, limit)
		if err != nil {
			hc.Logger.Error(err.Error())
			hc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			hc.Logger.Error(err.Error())
			hc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var posts model.Posts

		if err := json.Unmarshal(body, &posts); err != nil {
			hc.Logger.Error(err.Error())
			hc.Logger.Error(string(body))
			hc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		if err = hc.Presenter.ExecuteHomeIndex(w, r, &presenter.PostIndex{
			Posts: &posts,
		}); err != nil {
			hc.Logger.Error(err.Error())
			hc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
	})
}
