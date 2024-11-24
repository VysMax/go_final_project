package handlers

import (
	"VysMax/internalfunc"
	"VysMax/models"
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) OneTaskHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet:

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
		w.Write(resp)

	case r.Method == http.MethodPost:

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
		w.Write(resp)

	case r.Method == http.MethodPut:

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
		w.Write(resp)

	case r.Method == http.MethodDelete:

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
		w.Write(resp)
	}

}
