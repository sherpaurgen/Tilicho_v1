package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) createApartmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "creates a new apartment")
}

func (app *application) getApartmentHandler(w http.ResponseWriter, r *http.Request) {
	id_parameter := chi.URLParam(r, "id")
	// here 10 is base number i,e decimal
	id, err := strconv.ParseInt(id_parameter, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "showing details for apartment id %d\n", id)
}
