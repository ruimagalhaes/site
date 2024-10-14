package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64
	Username string
	Password string
}

func HasUser(db *sql.DB) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user").Scan(&count)
	if err != nil {
		return false, fmt.Errorf("HasUser: %v", err)
	}
	return count > 0, nil
}

func CreateUser(db *sql.DB, username, password string) (int64, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}

	result, err := db.Exec("INSERT INTO user (username, password) VALUES (?, ?)", username, passwordHash)
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	return id, nil
}

func AuthenticateUser(db *sql.DB, username, password string) (int64, error) {
	var id int64
	var passwordHash []byte
	row := db.QueryRow("SELECT id, password FROM user WHERE username = ?", username)
	err := row.Scan(&id, &passwordHash)
	if err != nil {
		return 0, fmt.Errorf("loginUser: %v", err)
	}
	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(password))
	if err != nil {
		return 0, fmt.Errorf("loginUser: %v", err)
	}
	return id, nil
}

func GetUser(db *sql.DB, id int64) (User, error) {
	var u User
	row := db.QueryRow("SELECT * FROM user WHERE id = ?", id)
	if err := row.Scan(&u.Id, &u.Username, &u.Password); err != nil {
		if err == sql.ErrNoRows {
			return u, fmt.Errorf("userById %d: no such user", id)
		}
		return u, fmt.Errorf("userById %d: %v", id, err)
	}
	return u, nil
}
