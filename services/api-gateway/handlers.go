package main

import (
	"net/http"

	internal "api-gateway.fincore-platform.isez.dev/internal"
)

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "ok",
		"env":    app.config.env,
	}

	if err := app.writeJSON(w, http.StatusOK, data); err != nil {
		app.logger.Error("could not write JSON response", "error", err)
	}
}

func (app *application) accountsHandler(w http.ResponseWriter, r *http.Request) {
	accounts := internal.AccountService()
	data := map[string]interface{}{
		"message": "accounts endpoint",
		"data":    accounts,
	}

	if err := app.writeJSON(w, http.StatusOK, data); err != nil {
		app.logger.Error("could not write JSON response", "error", err)
	}
}
