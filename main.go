package main

import (
	"fmt"
	"log"
	"net/http"
)

type PageData struct {
	Result string
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", ArtHandler)
	fmt.Println("Server Running:")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
