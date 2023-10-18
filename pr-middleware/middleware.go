package prmiddleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				log.Println("Recovered from an error: ", r)

				if env := os.Getenv("ENV"); env != "production" {
					log.Println(string(debug.Stack()))
				}

				w.WriteHeader(http.StatusInternalServerError)
				resp, err := json.Marshal(ErrorResponse{Message: "Internal server error"})
				if err != nil {
					w.Write([]byte("Something went wrong"))
				}

				w.Header().Set("Content-Type", "application/json")
				_, err = w.Write(resp)
				if err != nil {
					log.Println(err)
				}
			}
		}()

		rw := &responseWriter{ResponseWriter: w}
		next.ServeHTTP(rw, r)
		rw.flush()
	})
}
