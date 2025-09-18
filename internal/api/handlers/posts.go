package handlers

import (
	"blog-api/internal/models"
	"blog-api/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type PostHandler struct {
	DB *gorm.DB
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	if result := h.DB.Order("created_at desc").Find(&posts); result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch posts")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, posts)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if result := h.DB.Create(&post); result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create post")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, post)
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")

	var post models.Post
	if result := h.DB.First(&post, postID); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
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
	if result := h.DB.First(&post, postID); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			utils.RespondWithError(w, http.StatusNotFound, "Post not found")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve post")
		}
		return
	}

	var updatedData models.Post
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	post.Title = updatedData.Title
	post.Content = updatedData.Content

	if result := h.DB.Save(&post); result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update post")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, post)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")

	result := h.DB.Delete(&models.Post{}, postID)

	if result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete post")
		return
	}

	if result.RowsAffected == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "Post not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Post deleted successfully"})
}
