package api

import (
	"blog-api/internal/api/handlers"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	postHandler := &handlers.PostHandler{DB: db}
	tagHandler := &handlers.TagHandler{DB: db}
	commentHandler := &handlers.CommentHandler{DB: db}

	r.Route("/posts", func(r chi.Router) {
		r.Get("/", postHandler.GetPosts)
		r.Post("/", postHandler.CreatePost)
		r.Get("/{postID}", postHandler.GetPost)
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
