package main

import (
	"site/handler"
	"site/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	model.StartDB()

	e := echo.New()
	e.Use(middleware.Logger())

	articlesHandler := handler.ArticlesHandler{}
	e.GET("/", articlesHandler.HandleGetArticleList)
	e.GET("/articles/:id", articlesHandler.HandleGetArticle)
	e.PUT("/articles/:id", articlesHandler.HandlePutArticle)
	e.POST("/articles/:id", articlesHandler.HandlePutArticle) //change to proper verb
	e.GET("/articles/:id/edit", articlesHandler.HandleEditArticle)
	e.GET("/articles/new", articlesHandler.HandleNewArticle)
	e.POST("/articles", articlesHandler.HandlePostArticle)
	e.DELETE("/articles/:id", articlesHandler.HandleDeleteArticle)
	e.GET("/articles/:id/delete", articlesHandler.HandleDeleteArticle) //change to proper verb

	e.Logger.Fatal(e.Start(":8080"))
}
