package handlers

import (
	"VysMax/internalfunc"
	"VysMax/models"
	"encoding/json"
	"net/http"
	"time"
)

func (h *Handler) MarkAsDone(w http.ResponseWriter, r *http.Request) {

	var result models.Response
	id := r.FormValue("id")
	updatedTask, err := h.repo.GetSingle(id)
	if err != nil {
		result.Error = "Не указан идентификатор"
	}

	switch {
	case updatedTask.Repeat == "" && result.Error == "":
		err = h.repo.DeleteRow(id)
		if err != nil {
			result.Error = "Задача не найдена"
		}
	case updatedTask.Repeat != "" && result.Error == "":
		updatedTask.Date, err = internalfunc.NextDate(time.Now(), updatedTask.Date, updatedTask.Repeat)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = h.repo.UpdateDate(updatedTask)
		if err != nil {
			result.Error = "Задача не найдена"
		}
	default:
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
