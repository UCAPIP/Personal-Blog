package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Returns the home page (where the articles are located)
func handlerHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	formattedJson, err := loadArticles(jsonPath)
	if err != nil {
		log.Printf("ERROR: Unable to load articles from %s: %v", jsonPath, err)
		http.Error(w, "Internal server error: Unable to load data", http.StatusInternalServerError)
		return
	}

	data := PageHome{
		Articles: formattedJson,
	}

	tmpl, err := template.ParseFiles("ui/index.html")
	if err != nil {
		http.Error(w, "Template not found: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)

	fmt.Println(data, "handlerGetJson - ok")
}

func handlerGetJson(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	formattedJson, err := loadArticles(jsonPath)
	if err != nil {
		log.Printf("ERROR: Unable to load articles from %s: %v", jsonPath, err)
		http.Error(w, "Internal server error: Unable to load data", http.StatusInternalServerError)
		return
	}

	data := PageHome{
		Articles: formattedJson,
	}

	tmpl, err := template.ParseFiles("ui/index.html")
	if err != nil {
		http.Error(w, "Template not found: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)

	fmt.Println(data, "handlerGetJson - ok")
}
