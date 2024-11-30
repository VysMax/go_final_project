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

	parsedNow, err := time.Parse(internalfunc.Layout, now)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	newTime, err := internalfunc.NextDate(parsedNow, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	_, _ = w.Write([]byte(newTime))
}
