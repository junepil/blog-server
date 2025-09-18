package handlers

import (
	"database/sql"
	"net/http"
)

type TagHandler struct {
	DB *sql.DB
}

func (h *TagHandler) GetTags(w http.ResponseWriter, r *http.Request) {
	// Placeholder
	w.Write([]byte("GetTags"))
}

func (h *TagHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	// Placeholder
	w.Write([]byte("CreateTag"))
}

func (h *TagHandler) GetPostsByTag(w http.ResponseWriter, r *http.Request) {
	// Placeholder
	w.Write([]byte("GetPostsByTag"))
}
