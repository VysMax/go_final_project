package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	var (
		resp []byte
		err  error
	)
	search := r.FormValue("search")
	tasks := h.repo.GetMultiple(search)

	switch {
	case tasks.Tasks == nil:
		emptySlice := []string{}
		emptyMap := map[string][]string{
			"tasks": emptySlice,
		}
		resp, err = json.Marshal(emptyMap)
	default:
		resp, err = json.Marshal(tasks)

	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
