package prmiddleware_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	prmiddleware "pr-middleware"
	"testing"
)

func TestRecover1(t *testing.T) {
	mux := http.NewServeMux()

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("Panicing test recover 1")
	})

	mux.Handle("/", prmiddleware.RecoverMiddleware(h))

	go http.ListenAndServe(":3000", mux)

	handle := handleResp(t)
	handle(http.Get("http://localhost:3000"))
}

func TestRecover2(t *testing.T) {
	mux := http.NewServeMux()

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/pdf")
		panic("Panicing test recover 2")
	})

	mux.Handle("/", prmiddleware.RecoverMiddleware(h))

	go http.ListenAndServe(":3002", mux)

	handle := handleResp(t)
	handle(http.Get("http://localhost:3002"))
}

func TestRecover3(t *testing.T) {
	mux := http.NewServeMux()

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
		w.Write([]byte("How are you doing?"))
		panic("Panicing test recover 3")
	})

	mux.Handle("/", prmiddleware.RecoverMiddleware(h))

	go http.ListenAndServe(":3003", mux)

	handle := handleResp(t)
	handle(http.Get("http://localhost:3003"))
}

func handleResp(t *testing.T) func(*http.Response, error) {
	return func(r *http.Response, err error) {
		if err != nil {
			t.Error("Failed with:", err)
		}

		var data struct {
			Message string `json:"message"`
		}
		defer r.Body.Close()
		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			t.Error("Failed with error: ", err)
		}

		fmt.Printf("%+v \n", data)

	}
}
