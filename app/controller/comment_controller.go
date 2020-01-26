package controller

import (
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/response"
)

// A CommentControlller is a controller for a comment.
type CommentControlller struct {
	Logger   *logger.Logger
	Client   *api.Client
	Response *response.Response
}

// NewCommentController creates a NewCommentController.
func NewCommentController(logger *logger.Logger, client *api.Client, response *response.Response) *CommentControlller {
	return &CommentControlller{
		Logger:   logger,
		Client:   client,
		Response: response,
	}
}

// Store displays a listing of the resource.
func (cc *CommentControlller) Store(w http.ResponseWriter, r *http.Request) {
	// TODO:
	cc.Response.Error(w, http.StatusInternalServerError)
}
