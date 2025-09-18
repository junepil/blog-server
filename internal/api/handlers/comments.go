package handlers

import (
	"database/sql"
	"net/http"
)

type CommentHandler struct {
	DB *sql.DB
}

func (h *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	// Placeholder
	w.Write([]byte("GetComments"))
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	// Placeholder
	w.Write([]byte("CreateComment"))
}
