package handlers

import "VysMax/database"

type Handler struct {
	repo *database.Repository
}

func NewHandler(repo *database.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}
