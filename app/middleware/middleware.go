// TODO: Implement csrf
package middleware

// import (
// 	"context"
// 	"net/http"

// 	"github.com/bmf-san/gobel-admin-client/app/cookie"
// 	"github.com/bmf-san/gobel-admin-client/app/logger"
// 	"github.com/bmf-san/gobel-admin-client/app/response"
// 	"github.com/bmf-san/gobel-admin-client/app/session"
// )

// const (
// 	ctxCSRFToken = "csrf-token"
// )

// // middelware represents the singular of middleware.
// type middleware func(http.HandlerFunc) http.HandlerFunc

// // Middlewares represents the plural of middleware.
// type Middlewares struct {
// 	middlewares []middleware
// }

// // Asset represents the plural of middelware services.
// type Asset struct {
// 	redisHandler *session.RedisHandler
// 	logger       *logger.Logger
// 	cookie       *cookie.Cookie
// 	response     *response.Response
// }

// // NewAsset creates a templates.
// func NewAsset(redisHandler *session.RedisHandler, logger *logger.Logger, cookie *cookie.Cookie, presenter *presenter.Presenter) Asset {
// 	return Asset{
// 		redisHandler: redisHandler,
// 		logger:       logger,
// 		cookie:       cookie,
// 		response:     response,
// 	}
// }

// // NewMiddlewares creates a middlewares.
// func NewMiddlewares(mws ...middleware) Middlewares {
// 	return Middlewares{
// 		middlewares: append([]middleware(nil), mws...),
// 	}
// }

// // Then handles chaining middlewares.
// func (mws Middlewares) Then(h http.HandlerFunc) http.HandlerFunc {
// 	for i := range mws.middlewares {
// 		h = mws.middlewares[len(mws.middlewares)-1-i](h)
// 	}

// 	return h
// }

// // CSRF is a middelware for CSRF.
// func (a *Asset) CSRF(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodGet:
// 			token, err := a.redisHandler.SetCSRFToken()
// 			if err != nil {
// 				a.logger.Error(err.Error())
// 				a.Presenter.Error(w, http.StatusInternalServerError)
// 				return
// 			}
// 			ctx := context.WithValue(r.Context(), ctxCSRFToken, token)
// 			next.ServeHTTP(w, r.WithContext(ctx))
// 			return

// 		case http.MethodPost:
// 			r.ParseForm()
// 			token := r.Form.Get("csrf_token")
// 			if token == "" {
// 				a.logger.Error("CSRF token is invalid")
// 				a.Presenter.Error(w, http.StatusInternalServerError)
// 				return
// 			}

// 			if err := a.redisHandler.HasCSRFToken(token); err != nil {
// 				a.logger.Error(err.Error())
// 				a.Presenter.Error(w, http.StatusInternalServerError)
// 				return
// 			}

// 			next.ServeHTTP(w, r)
// 			return
// 		default:
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 	}
// }
