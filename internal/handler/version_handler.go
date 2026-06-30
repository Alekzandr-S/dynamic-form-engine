package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alekzandr-s/dynamic-form-engine/internal/domain"
	"github.com/alekzandr-s/dynamic-form-engine/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type VersionHandler struct {
	service *service.VersionService
}

func NewVersionHandler(service *service.VersionService) *VersionHandler {
	return &VersionHandler{
		service: service,
	}
}

func (h *VersionHandler) Create(w http.ResponseWriter, r *http.Request) {
	definitionID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid definition id", http.StatusBadRequest)
	}

	var version domain.FormTemplateVersion

	if err := json.NewDecoder(r.Body).Decode(&version); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	version.DefinitionID = definitionID

	if err := h.service.CreateDraft(r.Context(), &version); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	version.DefinitionID = definitionID

	w.WriteHeader(http.StatusCreated)
}

func (h *VersionHandler) Publish(w http.ResponseWriter, r *http.Request) {
	definitionID, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid definition id", http.StatusBadRequest)
		return
	}

	if err := h.service.Publish(r.Context(), definitionID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
