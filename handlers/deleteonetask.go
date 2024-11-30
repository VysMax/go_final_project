package handlers

import (
	"encoding/json"
	"net/http"

	"VysMax/models"
)

func (h *Handler) DeleteOneTask(w http.ResponseWriter, r *http.Request) {

	var result models.Response

	id := r.FormValue("id")

	err := h.repo.DeleteRow(id)

	if err != nil {
		result.Error = "Задача не найдена"
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}
