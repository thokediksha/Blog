// Package routers provide the routes for the application api
package routers

// Imports
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"Blog/controllers"
)

// InitialzeRoutes : initalizing routes to the API
func InitialzeRoutes() *gin.Engine {

	r := gin.Default()
	r.Use(gin.Logger())

	// r.Use(cors.Default())
	r.Use(func(c *gin.Context) {
		//allow all
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			platform := c.GetHeader("User-agent")
			fmt.Println("platform ", platform)
		}
		c.Next()
	})
    
	article := r.Group("/a")
	article.POST("", controllers.CreateArticle)
	article.GET("/", controllers.ListArticles)
	article.GET("/:id", controllers.ArticleContent)

	comment := r.Group("/c")
	comment.POST("/:article_id", controllers.PostComment)
	comment.POST("/:article_id/:parent_comment_id", controllers.CommentOnComment)
	return r
}