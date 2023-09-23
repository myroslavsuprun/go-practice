package handler

import (
	"fmt"
	"net/http"

	"database/sql"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	type pair struct {
		Path string
		Url  string
	}

	type pairs struct {
		Pairs []pair
	}

	var prs pairs

	err := yaml.Unmarshal(yml, &prs)
	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string, len(prs.Pairs))
	for _, entry := range prs.Pairs {
		pathsToUrls[entry.Path] = entry.Url
	}

	fmt.Println(pathsToUrls)
	return MapHandler(pathsToUrls, fallback), err
}

func JSONHandler(jsonB []byte, fallback http.Handler) (http.HandlerFunc, error) {
	type pair struct {
		Path string `json:"path"`
		Url  string `json:"url"`
	}

	var prs []pair

	err := json.Unmarshal(jsonB, &prs)
	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string, len(prs))
	for _, entry := range prs {
		pathsToUrls[entry.Path] = entry.Url
	}
	return MapHandler(pathsToUrls, fallback), nil
}

func DBHandler(db *sql.DB, fallback http.Handler) (http.HandlerFunc, error) {
	type pair struct {
		path string
		url  string
	}

	var prs []pair

	rows, err := db.Query("SELECT path, url FROM pairs")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p pair
		err = rows.Scan(&p.path, &p.url)
		if err != nil {
			return nil, err
		}
		prs = append(prs, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string, len(prs))
	for _, entry := range prs {
		pathsToUrls[entry.path] = entry.url
	}

	return MapHandler(pathsToUrls, fallback), nil
}
