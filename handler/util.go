package handler

import (
	"database/sql"
	"fmt"
	"site/model"

	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func contextHasUserSession(c echo.Context) error {
	database, ok := c.Get("database").(*sql.DB)
	if !ok {
		return fmt.Errorf("unable to get database")
	}

	hasUser, err := model.HasUser(database)
	if err != nil {
		return err
	}
	if !hasUser {
		return fmt.Errorf("unable to get user")
	} else {
		session, ok := c.Get("session").(*sessions.Session)
		if !ok {
			return fmt.Errorf("unable to get session")
		}
		userID, ok := session.Values["userId"]
		if !ok {
			return fmt.Errorf("user not logged in")
		}
		userIDInt64, ok := userID.(int64)
		if !ok {
			return fmt.Errorf("user not logged in")
		}
		_, err := model.GetUser(database, userIDInt64)
		if err != nil {
			return fmt.Errorf("user session not matching up")
		}

		return nil
	}
}
