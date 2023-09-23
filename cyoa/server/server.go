package server

import (
	"cyoa/stories"
	"cyoa/templates"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func Init(tmpl *template.Template, story stories.Story) error {
	mux := getMux()
	addTmplHandlers(mux, tmpl, story)

	fmt.Println("Starting server on port 3333")
	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return err
	}
	return nil
}

func addTmplHandlers(mux *http.ServeMux, t *template.Template, st stories.Story) {
	for sName, chapter := range st {
		mux.HandleFunc(fmt.Sprintf("/%s", sName), templateHandler(t, chapter))
	}
}

func templateHandler(t *template.Template, ch stories.Chapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, templates.ParseTmplChapter(ch))
		if err != nil {
			fmt.Printf("Error executing template: %s\n", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
	}
}

func getMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	return mux
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is my website!\n")
}
