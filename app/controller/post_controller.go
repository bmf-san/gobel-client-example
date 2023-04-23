package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/model"
	"github.com/bmf-san/gobel-client-example/app/presenter"
	"github.com/bmf-san/goblin"
)

// A PostController is a controller for a post.
type PostController struct {
	Logger    *logger.Logger
	Client    *api.Client
	Presenter *presenter.Presenter
}

// NewPostController creates a NewPostController.
func NewPostController(logger *logger.Logger, client *api.Client, presenter *presenter.Presenter) *PostController {
	return &PostController{
		Logger:    logger,
		Client:    client,
		Presenter: presenter,
	}
}

// Index displays a listing of the resource.
func (pc *PostController) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, limit, err := pc.Client.GetPageAndLimit(r)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		resp, err := pc.Client.GetPosts(page, limit)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var posts model.Posts

		if err := json.Unmarshal(body, &posts); err != nil {
			pc.Logger.Error(err.Error())
			pc.Logger.Error(string(body))
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		if err = pc.Presenter.ExecutePostIndex(w, r, &presenter.PostIndex{
			Posts: &posts,
			Pagination: &presenter.Pagination{
				Pager:       &pagination,
				QueryParams: "",
			},
		}); err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
	})
}

// IndexByKeyword displays a listing of the resource.
func (pc *PostController) IndexByKeyword() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, limit, err := pc.Client.GetPageAndLimit(r)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		keyword := r.URL.Query().Get("keyword")
		resp, err := pc.Client.GetPostsByKeyword(keyword, page, limit)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var posts model.Posts

		if err := json.Unmarshal(body, &posts); err != nil {
			pc.Logger.Error(err.Error())
			pc.Logger.Error(string(body))
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		if err = pc.Presenter.ExecutePostIndexByKeyword(w, r, &presenter.PostIndexBySearch{
			Posts: &posts,
			Pagination: &presenter.Pagination{
				Pager:       &pagination,
				QueryParams: template.URL(fmt.Sprintf("keyword=%s", keyword)),
			},
		}); err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
	})
}

// IndexByCategory displays a listing of the resource.
func (pc *PostController) IndexByCategory() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, limit, err := pc.Client.GetPageAndLimit(r)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		name := goblin.GetParam(r.Context(), "name")
		resp, err := pc.Client.GetPostsByCategory(name, page, limit)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var posts model.Posts

		if err := json.Unmarshal(body, &posts); err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		if err = pc.Presenter.ExecutePostIndexByCategory(w, r, &presenter.PostIndexByCategory{
			CategoryName: name,
			Posts:        &posts,
			Pagination: &presenter.Pagination{
				Pager:       &pagination,
				QueryParams: "",
			},
		}); err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
	})
}

// IndexByTag displays a listing of the resource.
func (pc *PostController) IndexByTag() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, limit, err := pc.Client.GetPageAndLimit(r)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		name := goblin.GetParam(r.Context(), "name")
		resp, err := pc.Client.GetPostsByTag(name, page, limit)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var posts model.Posts

		if err := json.Unmarshal(body, &posts); err != nil {
			pc.Logger.Error(err.Error())
			pc.Logger.Error(string(body))
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var pagination model.Pagination
		if err := pagination.Convert(resp.Header); err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		if err = pc.Presenter.ExecutePostIndexByTag(w, r, &presenter.PostIndexByTag{
			TagName: name,
			Posts:   &posts,
			Pagination: &presenter.Pagination{
				Pager:       &pagination,
				QueryParams: "",
			},
		}); err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
	})
}

// Show displays the specified resource.
func (pc *PostController) Show() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		title := goblin.GetParam(r.Context(), "title")

		resp, err := pc.Client.GetPost(title)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		var post model.Post

		if err := json.Unmarshal(body, &post); err != nil {
			pc.Logger.Error(err.Error())
			pc.Logger.Error(string(body))
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}

		if err = pc.Presenter.ExecutePostShow(w, r, &presenter.PostShow{
			Post: &post,
		}); err != nil {
			pc.Logger.Error(err.Error())
			pc.Presenter.Error(w, http.StatusInternalServerError)
			return
		}
	})
}
