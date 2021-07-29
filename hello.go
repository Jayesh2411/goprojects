package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
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
		{Id: "3", Title: "Hello 3", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequest()

}

func handleRequest() {
	fmt.Println("this function handles the incoming request")
	fmt.Println("this function handles the incoming request")
	myrouter := gin.Default()

	myrouter.GET("/", sayHello)
	myrouter.GET("/about", about)

	myrouter.GET("/articles", getAllArticles)
	myrouter.GET("/article/:id", getArticleById)
	myrouter.POST("/article", createArticle)
	myrouter.DELETE("/article/:id", deleteArticle)
	myrouter.PATCH("/article/:id", updateArticle)
	myrouter.Run()
}

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func about(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Developer- Jayesh Sinha",
	})

}

func getAllArticles(c *gin.Context) {
	json.NewEncoder(c.Writer).Encode(Articles)
}

func getArticleById(c *gin.Context) {
	key := c.Param("id")
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(c.Writer).Encode(article)
		}
	}
}

func createArticle(c *gin.Context) {
	json.NewEncoder(c.Writer).Encode("Got it!!")
	var article Article
	c.BindJSON(&article)
	Articles = append(Articles, article)

	json.NewEncoder(c.Writer).Encode(article)

}

func deleteArticle(c *gin.Context) {
	key := c.Param("id")
	for index, article := range Articles {
		if article.Id == key {
			Articles = append(Articles[:index], Articles[index+1:]...)
			c.Writer.Write([]byte("Article with id " + key + " deleted!!"))
		}
	}
}

func updateArticle(c *gin.Context) {

	key := c.Param("id")

	body, _ := ioutil.ReadAll(c.Request.Body)

	var article Article
	json.Unmarshal(body, &article)
	for i, a := range Articles {
		if a.Id == key {
			Articles[i] = article
			c.Writer.Write([]byte("Article with id " + key + " updated!!"))
		}
	}
}
