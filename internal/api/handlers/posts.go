package handlers

import (
	"blog-api/internal/models"
	"blog-api/pkg/utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	DB *sql.DB
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT post_id, title, content, published_at, created_at, updated_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch posts")
		return
	}
	defer rows.Close()

	posts := []models.Post{}
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.PostID, &post.Title, &post.Content, &post.PublishedAt, &post.CreatedAt, &post.UpdatedAt); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to scan post")
			return
		}
		posts = append(posts, post)
	}

	utils.RespondWithJSON(w, http.StatusOK, posts)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := h.DB.QueryRow(
		"INSERT INTO posts (title, content) VALUES ($1, $2) RETURNING post_id, published_at, created_at, updated_at",
		post.Title, post.Content,
	).Scan(&post.PostID, &post.PublishedAt, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create post")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, post)
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")

	var post models.Post
	err := h.DB.QueryRow(
		"SELECT post_id, title, content, published_at, created_at, updated_at FROM posts WHERE post_id = $1",
		postID,
	).Scan(&post.PostID, &post.Title, &post.Content, &post.PublishedAt, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithError(w, http.StatusNotFound, "Post not found")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve post")
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, post)
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	_, err := h.DB.Exec(
		"UPDATE posts SET title = $1, content = $2, updated_at = NOW() WHERE post_id = $3",
		post.Title, post.Content, postID,
	)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update post")
		return
	}

	post.PostID, _ = strconv.Atoi(postID)
	utils.RespondWithJSON(w, http.StatusOK, post)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")

	res, err := h.DB.Exec("DELETE FROM posts WHERE post_id = $1", postID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete post")
		return
	}

	count, err := res.RowsAffected()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to check rows affected")
		return
	}

	if count == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "Post not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Post deleted successfully"})
}
