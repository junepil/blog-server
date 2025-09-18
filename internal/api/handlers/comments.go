package handlers

import (
	"blog-api/internal/models"
	"blog-api/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type CommentHandler struct {
	DB *gorm.DB
}

func (h *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.Atoi(chi.URLParam(r, "postID"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var comments []models.Comment
	if result := h.DB.Where("post_id = ?", postID).Find(&comments); result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch comments")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, comments)
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	postID, err := strconv.Atoi(chi.URLParam(r, "postID"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	comment.PostID = postID

	if result := h.DB.Create(&comment); result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create comment")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, comment)
}
