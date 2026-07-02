package main

import (
	"ascii-art-web-stylize/ascii"
	"html/template"
	"net/http"
)

func ArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong Method", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	file := "banner/" + banner + ".txt"
	Lban, err := ascii.LoadBanner(file)
	if err != nil {
		http.Error(w, "Error Loading Banner", http.StatusInternalServerError)
		return
	}

	genArt := ascii.GenerateArt(text, Lban)
	data := PageData{
		Result: genArt,
	}

	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/output.html",
	)

	if err != nil {
		http.Error(w, "Error Parsing Template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
