package server

import (
	"encoding/json"
	"net/http"

	"github.com/kafkaphoenix/gosurf/internal/usecases"
)

type ActionHandler struct {
	service *usecases.ActionService
}

func NewActionHandler(as *usecases.ActionService) *ActionHandler {
	return &ActionHandler{service: as}
}

func (h *ActionHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/v1/actions/next-probabilities", h.getNextActionProbabilities)
}

func (h *ActionHandler) getNextActionProbabilities(w http.ResponseWriter, r *http.Request) {
	actionType := r.URL.Query().Get("type")
	if actionType == "" {
		http.Error(w, "missing type parameter", http.StatusBadRequest)
		return
	}

	probabilities, err := h.service.GetNextActionProbabilities(actionType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(probabilities)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
