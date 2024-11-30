package handlers

import (
	"encoding/json"
	"net/http"

	"VysMax/models"
)

func (h *Handler) GetOneTask(w http.ResponseWriter, r *http.Request) {
	var (
		result models.Task
		resp   []byte
		err    error
	)

	id := r.FormValue("id")
	if id == "" {
		result.Error = "Не указан идентификатор"
	}

	if result.Error == "" {
		result, err = h.repo.GetSingle(id)
		if err != nil {
			result.Error = "Задача не найдена"
		}
	}

	resp, err = json.Marshal(result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}
