package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rollexxx23/crud-app/routes"
)

func main() {
	fmt.Println("Welcome Welcome Welcome...")
	addr := flag.String("addr", ":6969", "HTTP network address")
	flag.Parse()
	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fileServer := http.FileServer(http.Dir("../client/static/"))
	mux := http.NewServeMux() // creates the server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/snippet/get-list", routes.ViewSnippetById)
	mux.HandleFunc("/snippet/create", routes.CreateSnippet)
	mux.HandleFunc("/", routes.Home) // routing handler
	log.Printf("Starting server on %s", *addr)

	err = http.ListenAndServe(*addr, mux)
	log.Fatal(err)

}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://aamjujdd:OeHi-7z8o0ptdAs69NJqVRH14ms-zb6R@tiny.db.elephantsql.com/aamjujdd")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Connected to DB...")
	return db, nil
}
