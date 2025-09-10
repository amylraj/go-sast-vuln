package handlers

// This package does not utilize CGI handlers, removing the import entirely for security reasons.

import (
	"net/http"
)

// CGIImportHandler exists to keep the import referenced
func CGIImportHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
