package router

import (
	"net/http"

	"github.com/bmf-san/gobel-client-example/app/controller"
	"github.com/bmf-san/goblin"
)

// Router represents the singular of router.
type Router struct {
	Mux *goblin.Router
}

// NewRouter creates a Router.
func NewRouter() *Router {
	return &Router{
		Mux: goblin.NewRouter(),
	}
}

// SetHome sets routes for home.
func (r *Router) SetHome(hc *controller.HomeController) {
	r.Mux.GET("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hc.Index(w, r)
	}))
}

// SetPosts sets routes for posts.
func (r *Router) SetPosts(pc *controller.PostController) {
	r.Mux.GET("/posts", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pc.Index(w, r)
	}))
	r.Mux.GET("/posts/:title", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pc.Show(w, r)
	}))
	r.Mux.GET("/posts/categories/:name", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pc.IndexByCategory(w, r)
	}))
	r.Mux.GET("/posts/tags/:name", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pc.IndexByTag(w, r)
	}))
}

// SetComments sets routes for comments.
func (r *Router) SetComments(cc *controller.CommentControlller) {
	// TODO: implement later.
	// r.Mux.POST("/posts/:title/comments", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	cc.Store(w, r)
	// }))
}

// SetCategories sets routes for categories.
func (r *Router) SetCategories(cc *controller.CategoryController) {
	r.Mux.GET("/categories", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cc.Index(w, r)
	}))
}

// SetTags sets routes for tags.
func (r *Router) SetTags(tc *controller.TagController) {
	r.Mux.GET("/tags", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tc.Index(w, r)
	}))
}

// SetSitemap sets routes for sitemap.
func (r *Router) SetSitemap(sc *controller.SitemapController) {
	r.Mux.GET("/sitemap", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sc.Index(w, r)
	}))
}

// SetFeed sets routes for feed.
func (r *Router) SetFeed(fc *controller.FeedController) {
	r.Mux.GET("/feed", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fc.Index(w, r)
	}))
}
