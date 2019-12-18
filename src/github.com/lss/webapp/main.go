package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
)

func main() {
	templates := populdateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request) {
		requestedFile := r.URL.Path[1:]
		fmt.Println(r.URL.Path)
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
	// http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
}

func populdateTemplates() * template.Template {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}