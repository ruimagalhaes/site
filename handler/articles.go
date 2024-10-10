package handler

import (
	"errors"
	"fmt"
	"net/http"
	"site/model"
	"site/view"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticlesHandler struct{}

func (h ArticlesHandler) HandleGetArticleList(c echo.Context) error {
	articles, err := model.GetArticles()
	if err != nil {
		return err
	}
	return render(c, view.ArticleList(articles))
}

func (h ArticlesHandler) HandleGetArticle(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	article, err := model.GetArticle(id)
	if err != nil {
		return err
	}
	return render(c, view.Article(article))
}

func (h ArticlesHandler) HandleNewArticle(c echo.Context) error {
	emptyArticle := &model.Article{}
	return render(c, view.ArticleForm(emptyArticle))
}

func (h ArticlesHandler) HandleEditArticle(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	article, err := model.GetArticle(id)
	if err != nil {
		return err
	}
	return render(c, view.ArticleForm(&article))
}

func (h ArticlesHandler) HandlePostArticle(c echo.Context) error {
	title := c.FormValue("title")
	body := c.FormValue("body")
	id, err := model.CreateArticle(title, body)
	if err != nil {
		return err
	}
	article, err := model.GetArticle(id)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/articles/"+article.GetStrId())
}

func (h ArticlesHandler) HandlePutArticle(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	title := c.FormValue("title")
	body := c.FormValue("body")
	articleId, err := model.UpdateArticle(id, title, body)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/articles/"+fmt.Sprintf("%d", articleId))
}

func (h ArticlesHandler) HandleDeleteArticle(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	err = model.DeleteArticle(id)
	if err != nil {
		return errors.New("article not found")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
