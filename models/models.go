package models

// Book ...
type Book struct {
	Name        string  `json:"name"`
	Publication string  `json:"publication"`
	Author      *Author `json:"author"`
	Price       int64   `json:"price"`
	PublishDate string  `json:"publish_date"`
}

// Books ...
type Books struct {
	Book []Book `json:"books"`
}

// Author ..
type Author struct {
	Name              string   `json:"name"`
	Age               int64    `json:"age"`
	PublishedBooks    *Books   `json:"published_books"`
	PublishedArticles *Article `json:"published_articles"`
}

// Article ..
type Article struct {
	Title         string  `json:"title"`
	Author        *Author `json:"author"`
	PublishedDate string  `json:"published_date"`
	Content       string  `json:"content"`
}

type Page struct {
	Title string `json:"title"`
	Body  []byte `json:"body"`
}
