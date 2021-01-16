package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/controller"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/presenter"
	"github.com/bmf-san/gobel-client-example/app/router"
)

const timeout time.Duration = 10 * time.Second

func main() {
	threshold, _ := strconv.Atoi(os.Getenv("LOG_THRESHOLD"))
	offset, _ := strconv.Atoi(os.Getenv("LOG_TIME_ZONE_OFFSET"))
	location := time.FixedZone(os.Getenv("TIME_ZONE"), offset)

	logger := logger.NewLogger(threshold, location)
	client := api.NewClient()
	presenter := presenter.NewPresenter()

	hc := controller.NewHomeController(logger, client, presenter)
	pc := controller.NewPostController(logger, client, presenter)
	cc := controller.NewCategoryController(logger, client, presenter)
	tc := controller.NewTagController(logger, client, presenter)
	// TODO: implement later.
	// cmc := controller.NewCommentController(logger, client, presenter)
	sc := controller.NewSitemapController(logger, client, presenter)
	fc := controller.NewFeedController(logger, client, presenter)

	r := router.NewRouter()
	r.SetHome(hc)
	r.SetPosts(pc)
	r.SetCategories(cc)
	r.SetTags(tc)
	// TODO: implement later.
	// r.SetComments(cmc)
	r.SetSitemap(sc)
	r.SetFeed(fc)

	s := http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: r.Mux,
	}

	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			logger.Error(err.Error())
		}
	}()

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-q

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Gracefully shutdown")
}
