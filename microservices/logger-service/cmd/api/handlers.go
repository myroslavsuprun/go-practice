package main

import (
	"log-service/data"
	"net/http"
)

type JsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var payload JsonPayload
	_ = app.readJSON(w, r, &payload)

	err := app.Models.LogEntry.Insert(data.LogEntryModel{
		Name: payload.Name,
		Data: payload.Data,
	})
	if err != nil {
		app.writeJSONError(w, err)
		return
	}

	respPayload := jsonResponse{
		Error:   false,
		Message: "Log entry created successfully",
	}

	err = app.writeJSON(w, http.StatusCreated, respPayload)
	if err != nil {
		app.writeJSONError(w, err)
		return
	}
}
