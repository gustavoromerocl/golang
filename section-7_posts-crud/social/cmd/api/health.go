package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if err := writeJSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version}); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
