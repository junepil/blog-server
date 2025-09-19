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

	authHandler := &handlers.AuthHandler{DB: db}
	postHandler := &handlers.PostHandler{DB: db}
	tagHandler := &handlers.TagHandler{DB: db}
	commentHandler := &handlers.CommentHandler{DB: db}

	// Public routes
	r.Post("/login", authHandler.Login)

	r.Route("/posts", func(r chi.Router) {
		r.Get("/", postHandler.GetPosts)

		r.Route("/{postID}", func(r chi.Router) {
			r.Use(middleware.PostTrafficTracker(db))
			r.Get("/", postHandler.GetPost)
		})

		r.Get("/{postID}/comments", commentHandler.GetComments)
		r.Post("/{postID}/comments", commentHandler.CreateComment)
	})

	r.Route("/tags", func(r chi.Router) {
		r.Get("/", tagHandler.GetTags)
		r.Get("/{tagName}/posts", tagHandler.GetPostsByTag)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/posts", postHandler.CreatePost)
		r.Put("/posts/{postID}", postHandler.UpdatePost)
		r.Delete("/posts/{postID}", postHandler.DeletePost)

		r.Post("/tags", tagHandler.CreateTag)
	})

	return r
}
