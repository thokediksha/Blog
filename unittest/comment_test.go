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

func TestPostComment(t *testing.T) {

	gin.SetMode(gin.TestMode)
	body := bytes.NewBuffer([]byte("{\"ApiTest\":\"GetReq\"}"))
	req, err := http.NewRequest("POST", "c/:article_id", body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Setup your router, just like you did in your main function, and
	// register your routes
	router := gin.Default()
	router.POST("c/1", controller.PostComment)

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	fmt.Println(w.Body)
}

func TestCommentOnComment(t *testing.T) {

	gin.SetMode(gin.TestMode)
	body := bytes.NewBuffer([]byte("{\"ApiTest\":\"GetReq\"}"))
	req, err := http.NewRequest("POST", "c/:article_id/:parent_comment_id", body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Setup your router, just like you did in your main function, and
	// register your routes
	router := gin.Default()
	router.POST("c/1/1", controller.PostComment)

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	fmt.Println(w.Body)
}
