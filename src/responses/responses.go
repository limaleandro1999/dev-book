package responses

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	dataJson, err := json.Marshal(data)
	if err != nil {
		Error(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(dataJson)
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, ResponseError{Error: err.Error()})
}
