package model

type Book struct {
	BookIdentity int64  `json:"book_identity"`
	BookName     string `json:"book_name"`
	Auther       string `josn:"auther"`
	Title        string `json:"title"`
	Content      string `json:"content"`
}
