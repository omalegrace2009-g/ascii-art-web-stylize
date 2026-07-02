package main

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Wrong Method", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Wrong URL path", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/input.html",
	)

	if err != nil {
		http.Error(w, "Error Parsing Template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
