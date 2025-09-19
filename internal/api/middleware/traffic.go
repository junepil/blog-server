package middleware

import (
	"blog-api/internal/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func PostTrafficTracker(db *gorm.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Let the next handler run first. This ensures we only log traffic
			// for successful requests (e.g., status 200 OK).
			next.ServeHTTP(w, r)

			postIDStr := chi.URLParam(r, "postID")
			if postIDStr == "" {
				// Not a route with a postID, so we do nothing.
				return
			}

			postID, err := strconv.Atoi(postIDStr)
			if err != nil {
				// Invalid postID, just log it and do nothing.
				log.Printf("Invalid postID for traffic tracking: %s", postIDStr)
				return
			}

			// Run the database insert in a separate goroutine
			// so it doesn't block the HTTP response.
			go func() {
				traffic := models.PostTraffic{
					PostID:    postID,
					ViewedAt:  time.Now(),
					IPAddress: r.RemoteAddr,
				}
				if err := db.Create(&traffic).Error; err != nil {
					log.Printf("Error saving post traffic: %v", err)
				}
			}()
		})
	}
}
