package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"VysMax/internalfunc"
	"VysMax/models"
)

func (h *Handler) PutOneTask(w http.ResponseWriter, r *http.Request) {

	var (
		updatedTask models.Task
		result      models.Response
		buf         bytes.Buffer
		errMessage  string
	)

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err = json.Unmarshal(buf.Bytes(), &updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = strconv.Atoi(updatedTask.ID)

	switch {
	case err != nil:
		result.Error = "Порядковый номер задачи указан в неверном формате"
	default:
		updatedTask, errMessage = internalfunc.CheckTaskFields(updatedTask)
		if errMessage != "" {
			result.Error = errMessage
		}
	}

	if result.Error == "" {
		err = h.repo.UpdateTask(updatedTask)
		if err != nil {
			result.Error = "Задача не найдена"
		}
	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}
