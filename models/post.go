package models

type post struct {
	Id      string
	Title   string
	Content string
}

func NewPost(id, title, content string) *post {
	return &post{id, title, content}
}