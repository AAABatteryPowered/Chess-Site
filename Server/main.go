package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hi")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../Client/index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Films", Director: "Me"},
			},
		}
		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("../Client/index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	home_redirect := func(w http.ResponseWriter, r *http.Request) {
		http.RedirectHandler("https://www.google.com/", http.StatusSeeOther)
	}

	http.HandleFunc("/home", h1)
	http.HandleFunc("/add-film/", h2)

	http.HandleFunc("/", home_redirect)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
