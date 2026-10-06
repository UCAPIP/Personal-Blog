package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

	fmt.Println(data, "handlerHome - ok")
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

func handlerArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handlerArticle - starts")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	formattedJson, err := loadArticles(jsonPath)
	if err != nil {
		log.Printf("ERROR: Unable to load articles from %s: %v", jsonPath, err)
		http.Error(w, "Internal server error: Unable to load data", http.StatusInternalServerError)
		return
	}

	targetArticle, err := findArticle(&formattedJson, id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("ui/templates/article.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, targetArticle)
	fmt.Println("handlerArticle - success")
}

func handlerAdmin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin" {
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

	tmpl, err := template.ParseFiles("ui/templates/admin.html")
	if err != nil {
		http.Error(w, "Template not found: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)

	fmt.Println(data, "handlerAdmin - ok")
}

type FormData struct {
	Heading     string // "New Article" или "Update Article"
	FormAction  string // "/admin/create" или "/admin/update/5"
	SubmitLabel string // "Publish" или "Update"
	Title       string
	Content     string
	PublishedAt string
}

func hadlerFormEdit(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	formattedJson, err := loadArticles(jsonPath)
	if err != nil {
		log.Printf("ERROR: Unable to load articles from %s: %v", jsonPath, err)
		http.Error(w, "Internal server error: Unable to load data", http.StatusInternalServerError)
		return
	}

	targetArticle, err := findArticle(&formattedJson, id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var editableArticle = FormData{
		Heading:     "Update Article",
		FormAction:  fmt.Sprintf("%d/update", targetArticle.ID),
		SubmitLabel: "Update",
		Title:       targetArticle.Title,
		Content:     targetArticle.Content,
		PublishedAt: targetArticle.PublishedAt,
	}

	tmpl, err := template.ParseFiles("ui/templates/form.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, editableArticle)
	fmt.Println("hadlerFormEdit - success")
}

func handlerUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "The method is not supported", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Form Data Processing Error", http.StatusBadRequest)
		return
	}

	title := r.PostFormValue("title")
	publishedAt := r.PostFormValue("published_at") // "YYYY-MM-DD"
	content := r.PostFormValue("content")

	if title == "" || publishedAt == "" || content == "" {
		http.Error(w, "All fields must be filled in", http.StatusBadRequest)
		return
	}

	log.Printf("Сохраняем статью: Title=%s, Date=%s", title, publishedAt)

	formattedJson, err := loadArticles(jsonPath)
	if err != nil {
		log.Printf("ERROR: Unable to load articles from %s: %v", jsonPath, err)
		http.Error(w, "Internal server error: Unable to load data", http.StatusInternalServerError)
		return
	}
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	editArticle(&formattedJson, id, content, title, publishedAt)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
	fmt.Printf("handlerUpdate id %d - success", id)
}
