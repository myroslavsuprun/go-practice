package main

import "net/http"

func (app *Config) sendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var reqPayload mailMessage

	err := app.readJSON(w, r, &reqPayload)
	if err != nil {
		app.writeJSONError(w, err)
		return
	}

	msg := &Message{
		From:    reqPayload.From,
		To:      reqPayload.To,
		Subject: reqPayload.Subject,
		Data:    reqPayload.Message,
	}

	err = app.Mailer.SendSMTP(msg)
	if err != nil {
		app.writeJSONError(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Email sent successfully",
	}

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.writeJSONError(w, err)
		return
	}
}
