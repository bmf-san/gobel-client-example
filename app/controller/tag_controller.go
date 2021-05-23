package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/model"
	"github.com/bmf-san/gobel-client-example/app/presenter"
)

// A TagController is a controller for a tag.
type TagController struct {
	Logger    *logger.Logger
	Client    *api.Client
	Presenter *presenter.Presenter
}

// NewTagController creates a TagController.
func NewTagController(logger *logger.Logger, client *api.Client, presenter *presenter.Presenter) *TagController {
	return &TagController{
		Logger:    logger,
		Client:    client,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (tc *TagController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, _, err := tc.Client.GetPageAndLimit(r)

		resp, err := tc.Client.GetTags(page, 100)
		if err != nil {
			tc.Logger.Error(err.Error())
			tc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			tc.Logger.Error(err.Error())
			tc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var tags model.Tags

		if err := json.Unmarshal(body, &tags); err != nil {
			tc.Logger.Error(err.Error())
			tc.Logger.Error(string(body))
			tc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			tc.Logger.Error(err.Error())
			tc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		if err = tc.Presenter.ExecuteTagIndex(w, &presenter.TagIndex{
			Tags:       &tags,
			Pagination: &pagination,
		}); err != nil {
			tc.Logger.Error(err.Error())
			tc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
	})
}
