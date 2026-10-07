package main

import (
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/article/{id}", handlerArticle)
	http.HandleFunc("/admin", basicAuth(handlerAdmin))
	http.HandleFunc("/admin/edit/{id}", basicAuth(handlerFormEdit))
	http.HandleFunc("POST /admin/edit/{id}/update", basicAuth(handlerUpdate))
	http.HandleFunc("/admin/new", basicAuth(handlerFormNew))
	http.HandleFunc("POST /admin/new/publish", basicAuth(handlerPublish))
	http.HandleFunc("POST /admin/delete/{id}", basicAuth(handlerDelete))

	log.Println("The server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
