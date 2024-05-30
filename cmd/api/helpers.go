package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type envelope map[string]any

func (app *application) readIDParam(r *http.Request) (int64, error) {
	id_parameter := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(id_parameter, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// here "any" is equivalent to interface{} is used to accept any type of data
// the status is the status code written to the response
// the headers is map of key value pairs
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
