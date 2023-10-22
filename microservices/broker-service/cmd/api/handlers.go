package main

import (
	"broker/event"
	"broker/logs"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
	Mail   MailPayload `json:"mail,omitempty"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	err := app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.writeJSONError(w, err)
	}
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type MailPayload struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func (app *Config) Handle(w http.ResponseWriter, r *http.Request) {
	var payload RequestPayload

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	if payload.Action == "authenticate" {
		app.authenticateUser(w, r, payload.Auth)
		return
	}

	if payload.Action == "log" {
		app.logItemViaRPC(w, r, payload.Log)
		return
	}

	if payload.Action == "mail" {
		app.sendMail(w, r, payload.Mail)
		return
	}

	app.writeJSONError(w, errors.New("Unsupported action"), http.StatusBadRequest)
}

func (app *Config) sendMail(w http.ResponseWriter, r *http.Request, payload MailPayload) {
	jsonData, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		app.writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	request, err := http.NewRequest("POST", "http://mail-service:8086/send", bytes.NewBuffer(jsonData))
	if err != nil {
		app.writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.writeJSONError(w, err, http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		app.writeJSONError(w, errors.New("Error calling mail service"), http.StatusBadRequest)
		return
	}

	var respPayload jsonResponse
	respPayload.Error = false
	respPayload.Message = "Successfully sent mail"

	err = app.writeJSON(w, http.StatusOK, respPayload)
	if err != nil {
		app.writeJSONError(w, err)
		return
	}
}

func (app *Config) logData(w http.ResponseWriter, r *http.Request, payload LogPayload) {
	jsonData, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		app.writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	request, err := http.NewRequest("POST", "http://logger-service:8085/log", bytes.NewBuffer(jsonData))
	if err != nil {
		app.writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		app.writeJSONError(w, err, http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		app.writeJSONError(w, errors.New("Error calling logger service"), http.StatusBadRequest)
		return
	}

	var respPayload jsonResponse
	respPayload.Error = false
	respPayload.Message = "Successfully logged data"
	err = app.writeJSON(w, response.StatusCode, respPayload)
	if err != nil {
		app.writeJSONError(w, err)
		return
	}

}

func (app *Config) authenticateUser(w http.ResponseWriter, r *http.Request, payload AuthPayload) {
	jsonData, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		app.writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	request, err := http.NewRequest("POST", "http://authentication-service:8084/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	fmt.Println("response: ", response, err)
	if err != nil {
		app.writeJSONError(w, err, http.StatusBadRequest)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		app.writeJSONError(w, errors.New("Invalid credentials"), http.StatusUnauthorized)
		return
	}

	var jsonFromService jsonResponse
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	if jsonFromService.Error {
		app.writeJSONError(w, errors.New(jsonFromService.Message), http.StatusUnauthorized)
		return
	}

	var authPayload jsonResponse
	authPayload.Error = false
	authPayload.Message = "Successfully authenticated"
	authPayload.Data = jsonFromService.Data

	err = app.writeJSON(w, http.StatusOK, authPayload)
	if err != nil {
		app.writeJSONError(w, err)
		return
	}

}

func (app *Config) logEventViaAmqp(w http.ResponseWriter, r *http.Request, payload LogPayload) {
	err := app.pushToQueue(payload.Name, payload.Data)
	if err != nil {
		app.writeJSONError(w, err)
	}

	var respPayload jsonResponse
	respPayload.Error = false
	respPayload.Message = "Successfully logged data using AMQP"

	err = app.writeJSON(w, http.StatusOK, respPayload)
	if err != nil {
		app.writeJSONError(w, err)
	}
}

func (app *Config) pushToQueue(name, msg string) error {
	emitter, err := event.NewEventEmitter(app.Rabbit)
	if err != nil {
		return err
	}
	payload := LogPayload{
		Name: name,
		Data: msg,
	}

	j, _ := json.Marshal(payload)
	err = emitter.Push(string(j), "log.INFO")
	if err != nil {
		return err
	}

	return nil

}

type RPCPayload struct {
	Name string
	Data string
}

func (app *Config) logItemViaRPC(w http.ResponseWriter, r *http.Request, payload LogPayload) {
	client, err := rpc.Dial("tcp", "logger-service:5001")
	if err != nil {
		app.writeJSONError(w, err)
		return
	}

	rpcPayload := RPCPayload{
		Name: payload.Name,
		Data: payload.Data,
	}

	var result string
	client.Call("RPCServer.LogInfo", rpcPayload, &result)
	if err != nil {
		app.writeJSONError(w, err)
		return
	}

	respPayload := jsonResponse{
		Error:   false,
		Message: result,
	}

	err = app.writeJSON(w, http.StatusOK, respPayload)
	if err != nil {
		app.writeJSONError(w, err)
		return
	}
}
func (app *Config) LogViaGRPC(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.writeJSONError(w, err)
		return
	}

	conn, err := grpc.Dial("logger-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		app.writeJSONError(w, err)
		return
	}
	defer conn.Close()

	c := logs.NewLogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.WriteLog(ctx, &logs.LogRequest{
		LogEntry: &logs.Log{
			Name: requestPayload.Log.Name,
			Data: requestPayload.Log.Data,
		},
	})
	if err != nil {
		app.writeJSONError(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged"

	app.writeJSON(w, http.StatusAccepted, payload)
}
