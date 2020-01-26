package controller

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/model"
	"github.com/bmf-san/gobel-client-example/app/response"
)

// A PostController is a controller for a post.
type PostController struct {
	Logger   *logger.Logger
	Client   *api.Client
	Response *response.Response
}

// NewPostController creates a NewPostController.
func NewPostController(logger *logger.Logger, client *api.Client, response *response.Response) *PostController {
	return &PostController{
		Logger:   logger,
		Client:   client,
		Response: response,
	}
}

// Index displays a listing of the resource.
func (pc *PostController) Index(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	_, body, err := pc.Client.GetPosts(r, defaultPage, defaultLimit)
	if err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var posts model.Posts

	if err := json.Unmarshal(body, &posts); err != nil {
		pc.Logger.Error(err.Error())
		pc.Logger.Error(string(body))
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	if err = pc.Response.ExecutePostIndex(w, &response.PostIndex{
		Posts: &posts,
	}); err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}
}

// IndexByCategory displays a listing of the resource.
func (pc *PostController) IndexByCategory(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	_, body, err := pc.Client.GetPostsByCategory(r, defaultPage, defaultLimit)
	if err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var posts model.Posts

	if err := json.Unmarshal(body, &posts); err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	if err = pc.Response.ExecutePostIndexByCategory(w, &response.PostIndexByCategory{
		Posts: &posts,
	}); err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}
}

// IndexByTag displays a listing of the resource.
func (pc *PostController) IndexByTag(w http.ResponseWriter, r *http.Request) {
	const defaultPage = 1
	const defaultLimit = 10

	_, body, err := pc.Client.GetPostsByCategory(r, defaultPage, defaultLimit)
	if err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var posts model.Posts

	if err := json.Unmarshal(body, &posts); err != nil {
		pc.Logger.Error(err.Error())
		pc.Logger.Error(string(body))
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	if err = pc.Response.ExecutePostIndexByTag(w, &response.PostIndexByTag{
		Posts: &posts,
	}); err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}
}

// Show displays the specified resource.
func (pc *PostController) Show(w http.ResponseWriter, r *http.Request) {
	_, body, err := pc.Client.GetPost(r)
	if err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	var post model.Post

	if err := json.Unmarshal(body, &post); err != nil {
		pc.Logger.Error(err.Error())
		pc.Logger.Error(string(body))
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}

	if err = pc.Response.ExecutePostShow(w, &response.PostShow{
		Post: &post,
	}); err != nil {
		pc.Logger.Error(err.Error())
		pc.Response.Error(w, http.StatusInternalServerError)
		return
	}
}
