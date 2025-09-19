package api

import (
	"blog-api/internal/api/handlers"
	"blog-api/internal/api/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(chiMiddleware.Recoverer)
	r.Use(middleware.RequestLogger)

	postHandler := &handlers.PostHandler{DB: db}
	tagHandler := &handlers.TagHandler{DB: db}
	commentHandler := &handlers.CommentHandler{DB: db}

	r.Route("/posts", func(r chi.Router) {
		r.Get("/", postHandler.GetPosts)
		r.Post("/", postHandler.CreatePost)
		
		r.Route("/{postID}", func(r chi.Router) {
			r.Use(middleware.PostTrafficTracker(db))
			r.Get("/", postHandler.GetPost)
		})

		r.Put("/{postID}", postHandler.UpdatePost)
		r.Delete("/{postID}", postHandler.DeletePost)

		r.Get("/{postID}/comments", commentHandler.GetComments)
		r.Post("/{postID}/comments", commentHandler.CreateComment)
	})

	r.Route("/tags", func(r chi.Router) {
		r.Get("/", tagHandler.GetTags)
		r.Post("/", tagHandler.CreateTag)
		r.Get("/{tagName}/posts", tagHandler.GetPostsByTag)
	})

	return r
}
