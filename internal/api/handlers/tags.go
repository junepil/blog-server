package handlers

import (
	"blog-api/internal/models"
	"blog-api/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type TagHandler struct {
	DB *gorm.DB
}

func (h *TagHandler) GetTags(w http.ResponseWriter, r *http.Request) {
	var tags []models.Tag
	if result := h.DB.Find(&tags); result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch tags")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, tags)
}

func (h *TagHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	var tag models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if result := h.DB.Create(&tag); result.Error != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create tag")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, tag)
}

func (h *TagHandler) GetPostsByTag(w http.ResponseWriter, r *http.Request) {
	tagName := chi.URLParam(r, "tagName")

	var tag models.Tag
	if result := h.DB.Preload("Posts").Where("name = ?", tagName).First(&tag); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			utils.RespondWithError(w, http.StatusNotFound, "Tag not found")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve tag")
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, tag.Posts)
}
