package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	app.logger.Printf("Error parsing id from url: %v", id)
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

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dest interface{}) error {
	//dest destination
	err := json.NewDecoder(r.Body).Decode(dest)
	if err != nil {
		// while reading json from reqest body there is likely to happen errors below
		var syntaxError *json.SyntaxError
		var unmarshallTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("bad json body at character %d", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly formed json")
		case errors.As(err, &unmarshallTypeError):
			if unmarshallTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect json type %q", unmarshallTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect %d", unmarshallTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
	return nil
}
