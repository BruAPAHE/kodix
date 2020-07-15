package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type Response struct {
	Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Version string `json:"version"`
	} `json:"meta"`
	Response interface{} `json:"response"`
}

func Send(body interface{}, e error, w http.ResponseWriter) {
	switch e {
	case ErrBadParams:
		create(body, http.StatusBadRequest, e, w)
		return
	case ErrServerError:
		create(body, http.StatusInternalServerError, e, w)
		return
	default:
		create(body, http.StatusOK, errors.New(""), w)
		return
	}
}

func create(body interface{}, status int, e error, w http.ResponseWriter) {
	r := &Response{
		Meta: struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
			Version string `json:"version"`
		}{
			Message: e.Error(),
			Code:    status,
			Version: os.Getenv("API_VERSION"),
		},
		Response: body,
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(r)
	return
}
