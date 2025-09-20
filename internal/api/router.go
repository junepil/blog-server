package api

import (
	"blog-api/internal/api/handlers"
	"blog-api/internal/api/middleware"
	"blog-api/internal/s3"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, uploader *s3.Uploader) http.Handler {
	r := chi.NewRouter()
	r.Use(chiMiddleware.Recoverer)
	r.Use(middleware.RequestLogger)

	authHandler := &handlers.AuthHandler{DB: db}
	postHandler := &handlers.PostHandler{DB: db}
	tagHandler := &handlers.TagHandler{DB: db}
	commentHandler := &handlers.CommentHandler{DB: db}
	imageHandler := &handlers.ImageHandler{Uploader: uploader}

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

		// Add the new protected image upload route
		r.Post("/upload/image", imageHandler.UploadImage)
	})

	return r
}
