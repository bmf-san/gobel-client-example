package controller

import (
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/presenter"
)

// A CommentControlller is a controller for a comment.
type CommentControlller struct {
	Logger    *logger.Logger
	Client    *api.Client
	Presenter *presenter.Presenter
}

// NewCommentController creates a NewCommentController.
func NewCommentController(logger *logger.Logger, client *api.Client, presenter *presenter.Presenter) *CommentControlller {
	return &CommentControlller{
		Logger:    logger,
		Client:    client,
		Presenter: presenter,
	}
}

// Store displays a listing of the resource.
func (cc *CommentControlller) Store(w http.ResponseWriter, r *http.Request) {
	// TODO: implement later.
}
