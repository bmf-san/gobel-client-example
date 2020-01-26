package controller

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/model"
	"github.com/bmf-san/gobel-client-example/app/response"
)

// A CategoryController is a controller for category.
type CategoryController struct {
	Logger   *logger.Logger
	Client   *api.Client
	Response *response.Response
}

// NewCategoryController creates a CategoryController.
func NewCategoryController(logger *logger.Logger, client *api.Client, response *response.Response) *CategoryController {
	return &CategoryController{
		Logger:   logger,
		Client:   client,
		Response: response,
	}
}

// Index displays a listing of the resource.
func (cc *CategoryController) Index(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	_, body, err := cc.Client.GetCategories(r, defaultPage, defaultLimit)
	if err != nil {
		cc.Logger.Error(err.Error())
		cc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var categories model.Categories

	if err := json.Unmarshal(body, &categories); err != nil {
		cc.Logger.Error(err.Error())
		cc.Logger.Error(string(body))
		cc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	if err = cc.Response.ExecuteCategoryIndex(w, &response.CategoryIndex{
		Categories: &categories,
	}); err != nil {
		cc.Logger.Error(err.Error())
		cc.Response.Error(w, http.StatusInternalServerError)
		return
	}
}
