package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"site/model"
	view "site/view/admin"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct{}

func (h AuthHandler) HandleGetAdmin(c echo.Context) error {
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	hasUser, err := model.HasUser(database)
	if err != nil {
		return err
	}
	if !hasUser {
		return render(c, view.UserForm())
	} else {
		err := contextHasUserSession(c)
		if err != nil {
			return render(c, view.LoginForm())
		} else {
			return render(c, view.Admin())
		}
	}
}

func (h AuthHandler) HandlePostUser(c echo.Context) error {
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	username := c.FormValue("username")
	password := c.FormValue("password")
	_, err := model.CreateUser(database, username, password)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/admin")
}

func (h AuthHandler) HandlePostLogin(c echo.Context) error {
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}
	username := c.FormValue("username")
	password := c.FormValue("password")
	id, err := model.AuthenticateUser(database, username, password)
	if err != nil {
		return err
	}
	session, ok := c.Get("session").(*sessions.Session)
	if !ok {
		return fmt.Errorf("unable to get session")
	}
	session.Values["userId"] = id
	err = session.Save(c.Request(), c.Response().Writer)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/admin")
}

func (h AuthHandler) HandlePostLogout(c echo.Context) error {
	session, ok := c.Get("session").(*sessions.Session)
	if !ok {
		return fmt.Errorf("unable to get session")
	}
	delete(session.Values, "userId")
	err := session.Save(c.Request(), c.Response().Writer)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
