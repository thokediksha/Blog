package unittest

import (
	controller "Blog/controllers"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateArticle(t *testing.T) {

	gin.SetMode(gin.TestMode)
	body := bytes.NewBuffer([]byte("{\"ApiTest\":\"GetReq\"}"))
	req, err := http.NewRequest("GET", "a", body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Setup your router, just like you did in your main function, and
	// register your routes
	router := gin.Default()
	router.GET("a", controller.CreateArticle)

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	fmt.Println(w.Body)
}

func TestListArticles(t *testing.T) {

	gin.SetMode(gin.TestMode)
	body := bytes.NewBuffer([]byte("{\"ApiTest\":\"GetReq\"}"))
	req, err := http.NewRequest("GET", "a/", body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Setup your router, just like you did in your main function, and
	// register your routes
	router := gin.Default()
	router.GET("a/", controller.ListArticles)

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	fmt.Println(w.Body)
}

func TestArticleContent(t *testing.T) {

	gin.SetMode(gin.TestMode)
	body := bytes.NewBuffer([]byte("{\"ApiTest\":\"GetReq\"}"))
	req, err := http.NewRequest("GET", "a/:id", body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Setup your router, just like you did in your main function, and
	// register your routes
	router := gin.Default()
	router.GET("a/1", controller.ArticleContent)

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	fmt.Println(w.Body)
}
