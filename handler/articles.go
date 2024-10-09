package handler

import (
	"errors"
	"net/http"
	"site/model"
	"site/view"

	"github.com/labstack/echo/v4"
)

type ArticlesHandler struct{}

func (h ArticlesHandler) HandleGetArticleList(c echo.Context) error {
	articles := model.GetArticles()
	return render(c, view.ArticleList(articles))
}

func (h ArticlesHandler) HandleGetArticle(c echo.Context) error {
	id := c.Param("id")
	article, err := model.GetArticle(id)
	if err != nil {
		return errors.New("article not found")
	}
	return render(c, view.Article(article))
}

func (h ArticlesHandler) HandleNewArticle(c echo.Context) error {
	emptyArticle := &model.Article{}
	return render(c, view.ArticleForm(emptyArticle))
}

func (h ArticlesHandler) HandleEditArticle(c echo.Context) error {
	id := c.Param("id")
	article, err := model.GetArticle(id)
	if err != nil {
		return errors.New("article not found")
	}
	return render(c, view.ArticleForm(&article))
}

func (h ArticlesHandler) HandlePostArticle(c echo.Context) error {
	title := c.FormValue("title")
	body := c.FormValue("body")
	article := model.CreateArticle(title, body)
	c.Redirect(http.StatusOK, "/articles/"+article.Id)
	return c.Redirect(http.StatusSeeOther, "/articles/"+article.Id)
}

func (h ArticlesHandler) HandlePutArticle(c echo.Context) error {
	id := c.Param("id")
	title := c.FormValue("title")
	body := c.FormValue("body")
	article, err := model.UpdateArticle(id, title, body)
	if err != nil {
		return errors.New("article not found")
	}
	return c.Redirect(http.StatusSeeOther, "/articles/"+article.Id)
}

func (h ArticlesHandler) HandleDeleteArticle(c echo.Context) error {
	id := c.Param("id")
	err := model.DeleteArticle(id)
	if err != nil {
		return errors.New("article not found")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
