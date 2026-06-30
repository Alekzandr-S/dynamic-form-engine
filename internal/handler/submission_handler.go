package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alekzandr-s/dynamic-form-engine/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type SubmissionHandler struct {
	engine *service.TemplateEngineService
}

func NewSubmissionHandler(engine *service.TemplateEngineService) *SubmissionHandler {
	return &SubmissionHandler{
		engine: engine,
	}
}

func (h *SubmissionHandler) GetPublished(w http.ResponseWriter, r *http.Request) {
	definitionID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid definition id", http.StatusBadRequest)
		return
	}

	version, err := h.engine.GetPublishedForm(r.Context(), definitionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(version)
}

func (h *SubmissionHandler) Submit(w http.ResponseWriter, r *http.Request) {
	definitionID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid definition id", http.StatusBadRequest)
		return
	}

	var submission map[string]any

	if err := json.NewDecoder(r.Body).Decode(&submission); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.engine.Submit(r.Context(), definitionID, submission); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}
