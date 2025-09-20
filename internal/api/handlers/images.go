package handlers

import (
	"blog-api/internal/s3"
	"blog-api/pkg/utils"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/google/uuid"
)

type ImageHandler struct {
	Uploader *s3.Uploader
}

// UploadImage handles the image upload request.
func (h *ImageHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	// 1. Parse the multipart form data.
	// 10 << 20 specifies a max upload size of 10 MB.
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "File too large")
		return
	}

	// 2. Retrieve the file from the form data.
	file, handler, err := r.FormFile("image")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid image file")
		return
	}
	defer file.Close()

	// 3. Generate a unique filename to prevent collisions.
	ext := filepath.Ext(handler.Filename)
	uniqueFilename := fmt.Sprintf("images/%s%s", uuid.NewString(), ext)

	// 4. Upload the file to S3.
	url, err := h.Uploader.UploadFile(r.Context(), uniqueFilename, file)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to upload image")
		return
	}

	// 5. Respond with the public URL of the uploaded image.
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"url": url})
}
