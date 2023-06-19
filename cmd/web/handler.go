package main

import (
	"net/http"

	"github.com/akxcix/log-router/pkg/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) handleHealthcheck(w http.ResponseWriter, r *http.Request) {
	data := "Inconceivable!"
	utils.RespondWithData(w, r, data)
}
