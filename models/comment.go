package models

// import "time"

type Comment struct {
	ID              int64  `json:"id"`
	ArticleId       int64  `json:"article_id"`
	NickName        string `form:"nick_name" json:"nick_name" binding:"required"`
	Content         string `form:"content" json:"content" binding:"required"`
	ParentCommentId int64  `json:"parent_comment_id"`
	CreatedAt       string `json:"created_at"`
}


