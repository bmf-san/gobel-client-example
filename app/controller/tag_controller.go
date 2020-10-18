package controller

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/model"
	"github.com/bmf-san/gobel-client-example/app/response"
)

// A TagController is a controller for a tag.
type TagController struct {
	Logger   *logger.Logger
	Client   *api.Client
	Response *response.Response
}

// NewTagController creates a TagController.
func NewTagController(logger *logger.Logger, client *api.Client, response *response.Response) *TagController {
	return &TagController{
		Logger:   logger,
		Client:   client,
		Response: response,
	}
}

// Index displays a listing of the resource.
func (tc *TagController) Index(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	resp, body, err := tc.Client.GetTags(r, defaultPage, defaultLimit)
	if err != nil {
		tc.Logger.Error(err.Error())
		tc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var tags model.Tags

	if err := json.Unmarshal(body, &tags); err != nil {
		tc.Logger.Error(err.Error())
		tc.Logger.Error(string(body))
		tc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var pagination model.Pagination
	if err := pagination.Convert(resp.Header); err != nil {
		tc.Logger.Error(err.Error())
		tc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	if err = tc.Response.ExecuteTagIndex(w, &response.TagIndex{
		Tags:       &tags,
		Pagination: &pagination,
	}); err != nil {
		tc.Logger.Error(err.Error())
		tc.Response.Error(w, http.StatusInternalServerError)
		return
	}
}
