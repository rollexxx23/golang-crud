package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rollexxx23/crud-app/routes"
)

func main() {
	fmt.Println("Hello World")
	fileServer := http.FileServer(http.Dir("../client/static/"))
	mux := http.NewServeMux() // creates the server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/snippet/get-list", routes.ViewSnippetById)
	mux.HandleFunc("/snippet/create", routes.CreateSnippet)
	mux.HandleFunc("/", routes.Home) // routing handler
	log.Println("Starting server on :6969")
	err := http.ListenAndServe(":6969", mux)
	log.Fatal(err)

}
