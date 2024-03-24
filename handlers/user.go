package handlers

import (
	"github.com/dev-hack95/go-htmx/structs"
	"github.com/dev-hack95/go-htmx/utilities"
	"github.com/dev-hack95/go-htmx/view/user"
	"github.com/labstack/echo/v4"
)

type UserHanler struct{}

func (h UserHanler) HandleUserShow(c echo.Context) error {
	u := structs.User{
		Email: "test1@gmail.com",
	}
	return utilities.Render(c, user.Show(u))
}
