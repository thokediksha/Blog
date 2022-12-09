package controllers

import (
	"Blog/models"
	"fmt"
	"strconv"

	// "fmt"
	// "net/http"
	// "strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func PostComment(c *gin.Context) {
	var comment models.Comment
	var article models.Article

	err := c.ShouldBindWith(&comment, binding.Form)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "could not parse the details"})
		return
	}
	now := time.Now()
	date := now.Format("Jan 2, 2006")

	comment.CreatedAt = date

	idStr := c.Param("article_id")
	id, _ := strconv.Atoi(idStr)

	comment.ArticleId = int64(id)

	if err = models.DB.Where("id = ?", comment.ArticleId).First(&article).Error; err != nil {
		c.JSON(400, gin.H{"message": "Article Does not exists"})
		return
	}

	err = models.DB.Select("created_at", "nick_name", "content", "article_id").Create(&comment).Error
	if err != nil {
		// fmt.Println(err)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"id":         comment.ID,
		"article_id": comment.ArticleId,
		"nickname":   comment.NickName,
		"content":    comment.Content,
		"created_at": comment.CreatedAt,
		"message":    "Succesfully posted comment",
	})
}

func CommentOnComment(c *gin.Context) {
	var comment models.Comment
	var article models.Article

	err := c.ShouldBindWith(&comment, binding.Form)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "could not parse the details"})
		return
	}

	now := time.Now()
	date := now.Format("Jan 2, 2006")

	comment.CreatedAt = date

	idStr := c.Param("article_id")
	id, _ := strconv.Atoi(idStr)

	comment.ArticleId = int64(id)

	if err = models.DB.Where("id = ?", comment.ArticleId).First(&article).Error; err != nil {
		c.JSON(400, gin.H{"message": "Article Does not exists"})
		return
	}

	idStr1 := c.Param("parent_comment_id")
	id1, _ := strconv.Atoi(idStr1)

	
	fmt.Println("id", comment.ParentCommentId)

	if err = models.DB.Where("id = ?", id1).First(&comment).Error; err != nil {
		c.JSON(400, gin.H{"message": "Comment Does not exists"})
		return
	}
	comment.ParentCommentId = int64(id1)
	fmt.Println("id", comment.ParentCommentId)

	err = models.DB.Select("created_at", "nick_name", "content", "article_id", "parent_comment_id").Create(&comment).Error
	if err != nil {
		// fmt.Println(err)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"id":                comment.ID,
		"article_id":        comment.ArticleId,
		"nickname":          comment.NickName,
		"content":           comment.Content,
		"created_at":        comment.CreatedAt,
		"parent_comment_id": comment.ParentCommentId,
		"message":           "Succesfully posted comment",
	})
}


