package controller

import (
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/response"
)

// A HomeController is a controller for a home.
type HomeController struct {
	Logger   *logger.Logger
	Client   *api.Client
	Response *response.Response
}

// NewHomeController creates a HomeController.
func NewHomeController(logger *logger.Logger, client *api.Client, response *response.Response) *HomeController {
	return &HomeController{
		Logger:   logger,
		Client:   client,
		Response: response,
	}
}

// Index displays a listing of the resource.
func (hc *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	if err := hc.Response.ExecuteHomeIndex(w); err != nil {
		hc.Logger.Error(err.Error())
		hc.Response.Error(w, http.StatusInternalServerError)
		return
	}
}
