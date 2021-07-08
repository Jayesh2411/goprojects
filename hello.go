package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {
	fmt.Println("Hello, world!!")

	Articles = []Article{
		{Id: "1", Title: "Hello 1", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequest()

}

func handleRequest() {
	myrouter := mux.NewRouter().StrictSlash(true)

	myrouter.HandleFunc("/", sayHello)
	myrouter.HandleFunc("/about", about)
	myrouter.HandleFunc("/articles", getAllArticles)
	myrouter.HandleFunc("/article/{id}", getArticleById).Methods("GET")
	myrouter.HandleFunc("/article", createArticle).Methods("POST")
	myrouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myrouter.HandleFunc("/article/{id}", updateArticle).Methods("PATCH")
	http.ListenAndServe(":9000", myrouter)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Developer Jayesh Sinha"))

}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All articles"))
	json.NewEncoder(w).Encode(Articles)
}

func getArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(body, &article)
	Articles = append(Articles, article)
	w.Write([]byte("Article added"))
	json.NewEncoder(w).Encode(article)

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for index, article := range Articles {
		if article.Id == key {
			Articles = append(Articles[:index], Articles[index+1:]...)
			w.Write([]byte("Article with id " + key + " deleted!!"))
		}
	}

}

func updateArticle(w http.ResponseWriter, r *http.Request) {

	key := mux.Vars(r)["id"]

	body, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(body, &article)
	for _, a := range Articles {
		if a.Id == key {
			a = article
			w.Write([]byte("Article with id " + key + " updated!!"))
		}
	}
}
