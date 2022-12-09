package controllers

import (
	// "fmt"
	// "strconv"
	"Blog/models"
	// "fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func CreateArticle(c *gin.Context) {
	var article models.Article
	
	err := c.ShouldBindWith(&article, binding.Form)
	if err != nil {
		c.JSON(400, gin.H{"message": "could not parse the details"})
		return
	}

	err = models.DB.Where("title = ?", article.Title).First(&article).Error
	if err == nil {
		c.JSON(400, gin.H{"error": "article already exists."})
		return
	}

	// article.CreatedAt = time.Now()
	now := time.Now()
        date := now.Format("Jan 2, 2006")
    
     article.CreatedAt = date

	err = models.DB.Select("title", "created_at", "nick_name", "content").Create(&article).Error
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"id":   article.ID,
		"title": article.Title,
		"nickname": article.NickName,
		"content": article.Content,
		"created_at": article.CreatedAt, 	
		"message": "Succesfully Created Article",
	})

}

func ListArticles(c *gin.Context) {
	var articles []models.Article
	p := c.Query("page")
	page, _ := strconv.Atoi(p)
	order := c.Query("order")
	err := models.DB.Order(order).Limit(2).Offset((page - 1) * 2).Find(&articles).Error
	if err != nil {
		c.JSON(404, gin.H{"message": "no data found"})
		return
	}
	c.JSON(http.StatusOK, articles)
}


func ArticleContent(c *gin.Context) {
  var article models.Article
  
  err := models.DB.Where("id= ? ",c.Param("id")).First(&article).Error
  if err != nil {
	c.JSON(400, gin.H{
		"message": "cannot find article in database",
	})
	return
  }
  c.JSON(200, gin.H{
	"content": article.Content,
})
}

