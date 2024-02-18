package main

import (
	"log"
	"html/template"
	"net/http"
)

type Film struct {
	Title string
	Director string
}

func main() {
	
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("cmd/templates/index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Shawshank Redemption", Director: "Frank Darabont"},
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "The Dark Knight", Director: "Christopher Nolan"},
			},
		}
		tmpl.Execute(w, films)
	}

	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
