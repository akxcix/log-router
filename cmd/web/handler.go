package main

import (
	"io/ioutil"
	"net/http"

	"github.com/akxcix/log-router/pkg/logstore"
	"github.com/akxcix/log-router/pkg/utils"
)

type Handler struct {
	logstore *logstore.Store
}

func NewHandler(store *logstore.Store) *Handler {
	return &Handler{
		logstore: store,
	}
}

func (h *Handler) notFound(w http.ResponseWriter, r *http.Request) {
	data := "Path Not Found"
	utils.RespondWithError(w, r, http.StatusNotFound, data)
}

func (h *Handler) handleHealthcheck(w http.ResponseWriter, r *http.Request) {
	data := "Inconceivable!"
	utils.RespondWithData(w, r, data)
}

func (h *Handler) handleLog(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.RespondWithError(w, r, http.StatusInternalServerError, "unable to read body")
	}

	dataString := string(data)

	err = h.logstore.Save(dataString)
	if err != nil {
		utils.RespondWithError(w, r, http.StatusInternalServerError, "unable to save log")
	}

	utils.RespondWithData(w, r, "received log")
}
