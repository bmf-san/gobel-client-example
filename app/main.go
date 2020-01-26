package main

import (
	"net/http"
	"os"
	"time"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/controller"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/response"
	"github.com/bmf-san/goblin"
)

func init() {
	location := os.Getenv("TIME_ZONE")
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	logger := logger.NewLogger()
	client := api.NewClient()
	response := response.NewResponse()

	homeController := controller.NewHomeController(logger, client, response)
	postController := controller.NewPostController(logger, client, response)
	categoryController := controller.NewCategoryController(logger, client, response)
	tagController := controller.NewTagController(logger, client, response)
	commentController := controller.NewCommentController(logger, client, response)
	sitemapController := controller.NewSitemapController(logger, client, response)
	feedController := controller.NewFeedController(logger, client, response)

	r := goblin.NewRouter()

	r.GET("/", homeController.Index)
	r.GET("/posts", postController.Index)
	r.GET("/posts/:title", postController.Show)
	r.GET("/posts/categories/:name", postController.IndexByCategory)
	r.GET("/posts/tags/:name", postController.IndexByTag)
	r.POST("/posts/:title/comments", commentController.Store)
	r.GET("/categories", categoryController.Index)
	r.GET("/tags", tagController.Index)
	r.GET("/sitemap", sitemapController.Index)
	r.GET("/feed", feedController.Index)

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r); err != nil {
		logger.Error(err.Error())
	}
}
