package main

import (
	"context"
	"embed"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"syscall"
	"time"

	"github.com/bmf-san/gobel-client-example/app/api"
	"github.com/bmf-san/gobel-client-example/app/controller"
	"github.com/bmf-san/gobel-client-example/app/logger"
	"github.com/bmf-san/gobel-client-example/app/presenter"
	"github.com/bmf-san/goblin"
)

//go:embed templates/*
var templates embed.FS

const timeout time.Duration = 10 * time.Second

func main() {
	threshold, _ := strconv.Atoi(os.Getenv("LOG_THRESHOLD"))
	offset, _ := strconv.Atoi(os.Getenv("LOG_TIME_ZONE_OFFSET"))
	location := time.FixedZone(os.Getenv("TIME_ZONE"), offset)

	logger := logger.NewLogger(threshold, location)

	defer func() {
		if x := recover(); x != nil {
			logger.Error(string(debug.Stack()))
		}
		os.Exit(1)
	}()

	client := api.NewClient()
	presenter := presenter.NewPresenter(templates)

	hc := controller.NewHomeController(logger, client, presenter)
	pc := controller.NewPostController(logger, client, presenter)
	cc := controller.NewCategoryController(logger, client, presenter)
	tc := controller.NewTagController(logger, client, presenter)
	// TODO: implement later.
	// cmc := controller.NewCommentController(logger, client, presenter)
	sc := controller.NewSitemapController(logger, client, presenter)
	fc := controller.NewFeedController(logger, client, presenter)

	r := goblin.NewRouter()

	r.Methods(http.MethodGet).Handler("/debug/pprof/", http.HandlerFunc(pprof.Index))
	r.Methods(http.MethodGet).Handler("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	r.Methods(http.MethodGet).Handler("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	r.Methods(http.MethodGet).Handler("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	r.Methods(http.MethodGet).Handler("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	r.Methods(http.MethodGet).Handler("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	r.Methods(http.MethodGet).Handler("/debug/pprof/heap", pprof.Handler("heap"))
	r.Methods(http.MethodGet).Handler("/debug/pprof/mutex", pprof.Handler("mutex"))
	r.Methods(http.MethodGet).Handler("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	r.Methods(http.MethodGet).Handler("/debug/pprof/block", pprof.Handler("block"))

	r.Methods(http.MethodGet).Handler(`/favicon.ico`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	}))

	r.Methods(http.MethodGet).Handler(`/ads.txt`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	}))

	r.Methods(http.MethodGet).Handler(`/style.css`, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	}))

	r.Methods(http.MethodGet).Handler(`/`, hc.Index())
	r.Methods(http.MethodGet).Handler(`/posts`, pc.Index())
	r.Methods(http.MethodGet).Handler(`/posts/search`, pc.IndexByKeyword())
	r.Methods(http.MethodGet).Handler(`/posts/:title`, pc.Show())
	r.Methods(http.MethodGet).Handler(`/posts/categories/:name`, pc.IndexByCategory())
	r.Methods(http.MethodGet).Handler(`/posts/tags/:name`, pc.IndexByTag())
	// TODO: implement later.
	// r.Methods("/posts/:title/comments").Handler(`/posts/:title/comments`, cc.Store())
	r.Methods(http.MethodGet).Handler(`/categories`, cc.Index())
	r.Methods(http.MethodGet).Handler(`/tags`, tc.Index())
	r.Methods(http.MethodGet).Handler(`/sitemap`, sc.Index())
	r.Methods(http.MethodGet).Handler(`/feed`, fc.Index())

	s := http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: r,
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
