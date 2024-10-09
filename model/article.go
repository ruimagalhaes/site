package model

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Article struct {
	Id    string
	Title string
	Body  string
}

var mockedArticles = []Article{
	{
		Id:    "1",
		Title: "First Article",
		Body:  "This is the content of the first article.",
	},
	{
		Id:    "2",
		Title: "Second Article",
		Body:  "Content for the second article goes here.",
	},
	{
		Id:    "3",
		Title: "Third Article",
		Body:  "And here's the content for the third article.",
	},
}

func GetArticles() []Article {
	return mockedArticles
}

func CreateArticle(title, body string) Article {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	article := Article{
		Id:    fmt.Sprintf("%d", r.Intn(10000)),
		Title: title,
		Body:  body,
	}
	mockedArticles = append(mockedArticles, article)
	return article
}

func GetArticle(id string) (Article, error) {
	for _, article := range mockedArticles {
		if article.Id == id {
			return article, nil
		}
	}
	return Article{}, errors.New("article not found")
}

func UpdateArticle(id, title, body string) (Article, error) {
	for i, article := range mockedArticles {
		if article.Id == id {
			mockedArticles[i].Title = title
			mockedArticles[i].Body = body
			return mockedArticles[i], nil
		}
	}
	return Article{}, errors.New("article not found")
}

func DeleteArticle(id string) error {
	for i, article := range mockedArticles {
		if article.Id == id {
			mockedArticles = append(mockedArticles[:i], mockedArticles[i+1:]...)
			return nil
		}
	}
	return errors.New("article not found")
}

func (a *Article) IsNew() bool {
	return a.Id == ""
}
