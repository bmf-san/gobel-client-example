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

// A CategoryController is a controller for category.
type CategoryController struct {
	Logger    *logger.Logger
	Client    *api.Client
	Presenter *presenter.Presenter
}

// NewCategoryController creates a CategoryController.
func NewCategoryController(logger *logger.Logger, client *api.Client, presenter *presenter.Presenter) *CategoryController {
	return &CategoryController{
		Logger:    logger,
		Client:    client,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (cc *CategoryController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, _, err := cc.Client.GetPageAndLimit(r)
		if err != nil {
			cc.Logger.Error(err.Error())
			cc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		resp, err := cc.Client.GetCategories(page, 100)
		if err != nil {
			cc.Logger.Error(err.Error())
			cc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			cc.Logger.Error(err.Error())
			cc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var categories model.Categories

		if err := json.Unmarshal(body, &categories); err != nil {
			cc.Logger.Error(err.Error())
			cc.Logger.Error(string(body))
			cc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			cc.Logger.Error(err.Error())
			cc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		if err = cc.Presenter.ExecuteCategoryIndex(w, &presenter.CategoryIndex{
			Categories: &categories,
			Pagination: &pagination,
		}); err != nil {
			cc.Logger.Error(err.Error())
			cc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
	})
}
