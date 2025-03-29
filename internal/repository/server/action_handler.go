package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kafkaphoenix/gosurf/internal/usecases"
)

type ActionHandler struct {
	service *usecases.ActionService
}

func NewActionHandler(as *usecases.ActionService) *ActionHandler {
	return &ActionHandler{service: as}
}

func (h *ActionHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/v1/users/{id}/actions/total", h.getActionsByID)
}

func (h *ActionHandler) getActionsByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	uid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
	}

	total, err := h.service.GetTotalActionsByID(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(total)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
