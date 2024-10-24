package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"site/handler"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	database, err := startDB()
	if err != nil {
		fmt.Errorf("database: %v", err)
		return
	}

	session := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	session.Options.HttpOnly = true
	session.Options.SameSite = http.SameSiteLaxMode

	e := echo.New()
	e.Static("/static", "./static")
	e.Use(middleware.Logger())
	e.Use(sessionMiddleware(session))
	e.Use(databaseMiddleware(database))

	articlesHandler := handler.ArticlesHandler{}
	e.GET("/", articlesHandler.HandleGetArticleList)
	e.GET("/articles/:id", articlesHandler.HandleGetArticle)
	e.PUT("/articles/:id", articlesHandler.HandlePutArticle)
	e.POST("/articles/:id", articlesHandler.HandlePutArticle) //TODO: change to proper verb
	e.GET("/articles/:id/edit", articlesHandler.HandleEditArticle)
	e.GET("/articles/new", articlesHandler.HandleNewArticle)
	e.POST("/articles", articlesHandler.HandlePostArticle)
	e.DELETE("/articles/:id", articlesHandler.HandleDeleteArticle)
	e.GET("/articles/:id/delete", articlesHandler.HandleDeleteArticle) //TODO: change to proper verb

	authHandler := handler.AuthHandler{}
	e.GET("/admin", authHandler.HandleGetAdmin)
	e.POST("/user", authHandler.HandlePostUser)
	e.POST("/login", authHandler.HandlePostLogin)
	e.POST("/logout", authHandler.HandlePostLogout)

	e.Logger.Fatal(e.Start(":8080"))
}

func startDB() (*sql.DB, error) {
	// cfg := mysql.Config{
	// 	User:   os.Getenv("DBUSER"),
	// 	Passwd: os.Getenv("DBPASS"),
	// 	Net:    "tcp",
	// 	Addr:   "127.0.0.1:3306",
	// 	DBName: "site",
	// }
	// Get a database handle.
	var err error

	// db, err = sql.Open("mysql", cfg.FormatDSN())
	// db, err := sql.Open("mysql", "user:password@/dbname")
	var db *sql.DB
	db, err = sql.Open("mysql", os.Getenv("DBUSER")+":"+os.Getenv("DBPASS")+"@/site?parseTime=true")
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, err
	}

	fmt.Println("Connected to the database!")
	return db, nil
}

func sessionMiddleware(store sessions.Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := store.Get(c.Request(), "session")
			c.Set("session", session)
			return next(c)
		}
	}
}

func databaseMiddleware(database *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("database", database)
			return next(c)
		}
	}
}
