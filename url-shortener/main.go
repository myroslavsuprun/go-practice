package main

import (
	"fmt"
	"net/http"

	"shortener/db"
	"shortener/flagReader"
	"shortener/handler"
	"shortener/parser"
)

func main() {
	Init()

	fileName, fileType := flagReader.Parse()
	fileData, err := parser.GetBytes(fileName)
	if err != nil {
		panic(err)
	}

	mux := defaultMux()

	var handlerF http.HandlerFunc

	if fileType == flagReader.None {
		dbInst, err := db.Connect()
		if err != nil {
			panic(err)
		}

		dbHandler, err := handler.DBHandler(dbInst, mux)
		if err != nil {
			panic(err)
		}

		handlerF = dbHandler
	}

	if fileType == flagReader.Yaml {
		yamlHandler, err := handler.YAMLHandler(fileData, mux)
		if err != nil {
			panic(err)
		}
		handlerF = yamlHandler
	}

	if fileType == flagReader.Json {
		jsonHandler, err := handler.JSONHandler(fileData, mux)
		if err != nil {
			panic(err)
		}
		handlerF = jsonHandler
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handlerF)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
