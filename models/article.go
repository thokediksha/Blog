package models


type Article struct {
	ID        int64     `json:"id"`
	Title     string    `form:"title" json:"title" binding:"required"`
	NickName  string    `form:"nick_name" json:"nick_name" binding:"required"`
	Content   string    `form:"content" json:"content" binding:"required"`
	CreatedAt string `json:"created_at"`
}