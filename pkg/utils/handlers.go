package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akxcix/log-router/pkg/dto"
	"github.com/rs/zerolog/log"
)

func RespondWithData(w http.ResponseWriter, r *http.Request, data interface{}) {
	res := &dto.Response{
		Status: http.StatusOK,
		Data:   data,
	}

	json, err := json.Marshal(res)
	if err != nil {
		msg := "Unable to marshall data json"
		log.Error().Err(err).Msg(msg)
		RespondWithInternalServerError(w, r, http.StatusInternalServerError, msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}

func RespondWithInternalServerError(w http.ResponseWriter, r *http.Request, s int, e string) {
	res := &dto.Response{
		Status: s,
		Data:   e,
	}

	json, err := json.Marshal(res)
	if err != nil {
		log.Error().Err(err).Msg("Something went wrong while marshalling json")
		json = []byte(fmt.Sprintf("{\"status\":%d,\"data\":\"%s\"}", http.StatusInternalServerError, "Something bad happened"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	w.Write(json)
}
