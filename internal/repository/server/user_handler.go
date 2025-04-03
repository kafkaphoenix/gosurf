package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kafkaphoenix/gosurf/internal/usecases"
)

type UserHandler struct {
	service *usecases.UserService
}

func NewUserHandler(us *usecases.UserService) *UserHandler {
	return &UserHandler{service: us}
}

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	// ordered starting with the most specific
	mux.HandleFunc("/v1/users/{id}/actions/total", h.getTotalActionsByID)
	mux.HandleFunc("/v1/users/{id}", h.getUserByID)

	// to avoid overlapping path issues
	mux.HandleFunc("/v1/referral-index", h.getReferralIndex)
}

func (h *UserHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	uid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(user)
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

func (h *UserHandler) getTotalActionsByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	uid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
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

func (h *UserHandler) getReferralIndex(w http.ResponseWriter, _ *http.Request) {
	rIndex := h.service.GetReferralIndex()

	js, err := json.Marshal(rIndex)
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
