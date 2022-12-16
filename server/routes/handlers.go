package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.Write([]byte("Error Page Not Found"))
		return
	}
	files := []string{
		"../client/html/pages/home.tmpl",
		"../client/html/pages/base.tmpl",
		"../client/html/components/nav.tmpl",
	}
	tf, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = tf.ExecuteTemplate(w, "base", nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func ViewSnippetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err == nil && id >= 1 {
		fmt.Fprintf(w, "Data for id %d is %d", id, id*10)
		return
	}
	w.Write([]byte("Snippet Not Found..."))

}

func CreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not Allowed", 405)
		return
	}
	w.Write([]byte("Create Your Snippets..."))
}
