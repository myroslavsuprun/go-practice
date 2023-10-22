package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

func (app *Config) authenticateUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &credentials)
	if err != nil {
		app.writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(credentials.Email)
	if err != nil {
		app.writeJSONError(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	valid, err := user.PasswordMatches(credentials.Password)
	if err != nil || !valid {
		app.writeJSONError(w, errors.New("invalid credentials"), http.StatusUnauthorized)
		return
	}

	err = app.logRequestViaRPC("authentication", fmt.Sprintf("%s logged in", user.Email))
	if err != nil {
		app.writeJSONError(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Successfully authenticated",
		Data:    user,
	}

	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.writeJSONError(w, err)
		return
	}
}

func (app *Config) logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	request, err := http.NewRequest(http.MethodPost, "http://logger-service:8085/log", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}

type LogPayload struct {
	Name string
	Data string
}

func (app *Config) logRequestViaRPC(name, data string) error {
	client, err := rpc.Dial("tcp", "logger-service:5001")
	if err != nil {
		return err
	}

	rpcPayload := struct {
		Name string
		Data string
	}{
		Name: name,
		Data: data,
	}

	var result string
	client.Call("RPCServer.LogInfo", rpcPayload, &result)
	if err != nil {
		return err
	}

	fmt.Println("RPC result: ", result)
	return nil
}
