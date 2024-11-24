package handlers

import "VysMax/DBManip"

type Handler struct {
	repo *DBManip.Repository
}

func NewHandler(repo *DBManip.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}
