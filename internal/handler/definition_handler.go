package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alekzandr-s/dynamic-form-engine/internal/domain"
	"github.com/alekzandr-s/dynamic-form-engine/internal/service"
)

type DefinitionHandler struct {
	service *service.DefinitionService
}

func NewDefinitionHandler(service *service.DefinitionService) *DefinitionHandler {
	return &DefinitionHandler{
		service: service,
	}
}

func (h *DefinitionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var definition domain.FormDefinition

	if err := json.NewDecoder(r.Body).Decode(&definition); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Create(
		r.Context(), &definition); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
