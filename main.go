package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Article struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	PublishedAt string `json:"published_at"`
	Content     string `json:"content"`
}

func (a *Article) EditArticleBody(description, title, date string) {
	a.Content = description
	a.Title = title
	a.PublishedAt = date
}

type PageHome struct {
	Articles []Article
}

// JSON Path
var jsonPath = "articles.json"

// Loading articles from a File
func loadArticles(path string) ([]Article, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Article{}, fmt.Errorf("The file does not exist")
		}
		return nil, fmt.Errorf("File read error: %w", err)
	}

	var articles []Article
	if err := json.Unmarshal(data, &articles); err != nil {
		return nil, fmt.Errorf("JSON parsing error: %w", err)
	}
	fmt.Println(articles, "loadArticles - ok")
	return articles, nil
}

// Adding a New Article
func addArticle(data *[]Article, content, title, date string) {
	maxID := func(articles []Article) int {
		maxId := 0
		for _, a := range articles {
			if a.ID > maxId {
				maxId = a.ID
			}
		}
		return maxId
	}(*data)

	newArticle := Article{
		ID:          maxID + 1,
		Title:       title,
		Content:     content,
		PublishedAt: date,
	}

	*data = append(*data, newArticle)

	newData, err := json.MarshalIndent(*data, "", "  ")
	if err != nil {
		fmt.Println("serialization error:", err)
		return
	}

	err = os.WriteFile(jsonPath, newData, 0644)
	if err != nil {
		fmt.Println("JSON writing error:", err)
		return
	}

}

// Editing the article description
func editArticle(data *[]Article, id int, content, title, date string) {
	for i, article := range *data {
		if article.ID == id {
			(*data)[i].EditArticleBody(content, title, date)

			newData, err := json.MarshalIndent(*data, "", "  ")
			if err != nil {
				fmt.Println("serialization error:", err)
				return
			}

			err = os.WriteFile(jsonPath, newData, 0644)
			if err != nil {
				fmt.Println("JSON writing error:", err)
				return
			}

			return

		}
	}

}

func deleteArticle(data *[]Article, id int) {

	for i, article := range *data {
		if article.ID == id {
			*data = append((*data)[:i], (*data)[i+1:]...)

			newData, err := json.MarshalIndent(*data, "", "  ")
			if err != nil {
				fmt.Println("serialization error:", err)
				return
			}

			err = os.WriteFile(jsonPath, newData, 0644)
			if err != nil {
				fmt.Println("JSON writing error:", err)
				return
			}

			return

		}
	}
}

func findArticle(data *[]Article, id int) (Article, error) {
	for i, article := range *data {
		if article.ID == id {
			return (*data)[i], nil
		}
	}
	return (*data)[0], fmt.Errorf("Article doesn't exist")
}

func main() {

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/article/{id}", handlerArticle)
	http.HandleFunc("/admin", basicAuth(handlerAdmin))
	http.HandleFunc("/admin/edit/{id}", basicAuth(hadlerFormEdit))
	http.HandleFunc("POST /admin/edit/{id}/update", basicAuth(handlerUpdate))
	http.HandleFunc("/admin/new", basicAuth(hadlerFormNew))
	http.HandleFunc("POST /admin/new/publish", basicAuth(handlerPublish))
	http.HandleFunc("POST /admin/delete/{id}", basicAuth(handlerDelete))

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
