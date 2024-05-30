package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sherpaurgen/Tilicho_v1/internal/data"
)

func (app *application) createUnitHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "creates a new Unit")
}

func (app *application) getUnitHandler(w http.ResponseWriter, r *http.Request) {
	unitid, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	unit := data.Unit{
		UnitID:      int(unitid),
		BuildingID:  "B17",
		OwnerID:     int(unitid),
		Title:       "titelsdf",
		Description: "asdf",
		FloorNumber: 123,
		PricePerDay: 232,
		CreatedAt:   time.Now(),
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"unit": unit}, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "Ther server encountered an error", http.StatusInternalServerError)
	}
}
