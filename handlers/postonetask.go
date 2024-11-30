package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"VysMax/internalfunc"
	"VysMax/models"
)

func (h *Handler) PostOneTask(w http.ResponseWriter, r *http.Request) {

	var (
		newTask models.Task
		result  models.Response
		buf     bytes.Buffer
		newID   int64
	)

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err = json.Unmarshal(buf.Bytes(), &newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	newTask, errMessage := internalfunc.CheckTaskFields(newTask)

	switch errMessage == "" {
	case false:
		result.Error = errMessage
	case true:
		newID, err = h.repo.AddRow(newTask)
		switch err == nil {
		case true:
			result.ID = newID
		case false:
			result.Error = "Не удалось добавить запись в базу данных"
		}

	}

	resp, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(resp)
}
