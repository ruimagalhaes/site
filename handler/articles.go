package handler

import (
	"database/sql"
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
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	articles, err := model.GetArticles(database)
	if err != nil {
		return err
	}
	return render(c, view.ArticleList(articles))
}

func (h ArticlesHandler) HandleGetArticle(c echo.Context) error {
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	article, err := model.GetArticle(database, id)
	if err != nil {
		return err
	}
	err = contextHasUserSession(c)
	return render(c, view.Article(article, err == nil))
}

func (h ArticlesHandler) HandleNewArticle(c echo.Context) error {
	err := contextHasUserSession(c)
	if err != nil {
		return fmt.Errorf("user not logged in")
	}
	emptyArticle := &model.Article{}
	return render(c, view.ArticleForm(emptyArticle))
}

func (h ArticlesHandler) HandleEditArticle(c echo.Context) error {
	err := contextHasUserSession(c)
	if err != nil {
		return fmt.Errorf("user not logged in")
	}
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	article, err := model.GetArticle(database, id)
	if err != nil {
		return err
	}
	return render(c, view.ArticleForm(&article))
}

func (h ArticlesHandler) HandlePostArticle(c echo.Context) error {
	err := contextHasUserSession(c)
	if err != nil {
		return fmt.Errorf("user not logged in")
	}
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	title := c.FormValue("title")
	body := c.FormValue("body")
	id, err := model.CreateArticle(database, title, body)
	if err != nil {
		return err
	}
	article, err := model.GetArticle(database, id)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/articles/"+article.GetStrId())
}

func (h ArticlesHandler) HandlePutArticle(c echo.Context) error {
	err := contextHasUserSession(c)
	if err != nil {
		return fmt.Errorf("user not logged in")
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	title := c.FormValue("title")
	body := c.FormValue("body")
	articleId, err := model.UpdateArticle(database, id, title, body)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/articles/"+fmt.Sprintf("%d", articleId))
}

func (h ArticlesHandler) HandleDeleteArticle(c echo.Context) error {
	err := contextHasUserSession(c)
	if err != nil {
		return fmt.Errorf("user not logged in")
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	err = model.DeleteArticle(database, id)
	if err != nil {
		return errors.New("article not found")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
