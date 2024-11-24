package handlers

import (
	"net/http"
	"time"

	"VysMax/internalfunc"
)

func (h *Handler) GetNextDate(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	parsedNow, err := time.Parse("20060102", now)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	newTime, err := internalfunc.NextDate(parsedNow, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write([]byte(newTime))
}
