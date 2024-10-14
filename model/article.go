package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id    int64
	Title string
	Body  string
}

func GetArticles(db *sql.DB) ([]Article, error) {
	var articles []Article
	rows, err := db.Query("SELECT * FROM article")
	if err != nil {
		return nil, fmt.Errorf("articles: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var a Article
		if err := rows.Scan(&a.Id, &a.Title, &a.Body); err != nil {
			return nil, fmt.Errorf("articles: %v", err)
		}
		articles = append(articles, a)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("articles: %v", err)
	}
	return articles, nil
}

func CreateArticle(db *sql.DB, title, body string) (int64, error) {
	result, err := db.Exec("INSERT INTO article (title, body) VALUES (?, ?)", title, body)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func GetArticle(db *sql.DB, id int64) (Article, error) {
	var a Article
	row := db.QueryRow("SELECT * FROM article WHERE id = ?", id)
	if err := row.Scan(&a.Id, &a.Title, &a.Body); err != nil {
		if err == sql.ErrNoRows {
			return a, fmt.Errorf("articlesById %d: no such article", id)
		}
		return a, fmt.Errorf("articlesById %d: %v", id, err)
	}
	return a, nil
}

func UpdateArticle(db *sql.DB, id int64, title, body string) (int64, error) {
	result, err := db.Exec("UPDATE article SET title = ?, body = ? WHERE id = ?", title, body, id)
	if err != nil {
		return 0, fmt.Errorf("updateArticle %d: %v", id, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("updateArticle %d: %v", id, err)
	}
	if rowsAffected == 0 {
		return 0, fmt.Errorf("updateArticle %d: no such article", id)
	}
	return id, nil
}

func DeleteArticle(db *sql.DB, id int64) error {
	result, err := db.Exec("DELETE FROM article WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("deleteArticle %d: %v", id, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("deleteArticle %d: %v", id, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("deleteArticle %d: no such article", id)
	}
	return nil
}

func (a *Article) IsNew() bool {
	return a.Id == 0
}

func (a *Article) GetStrId() string {
	return fmt.Sprintf("%d", a.Id)
}
