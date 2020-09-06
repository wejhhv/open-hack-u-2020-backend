package usercontroller

import (
	"net/http"

	user "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/User"
	"github.com/labstack/echo/v4"
)

type requestRegister struct {
	Name string `json:"name"`
}

type requestEdit struct {
	ID      int    `json:"id"`
	NewName string `json:"newName"`
}

//Register ...ユーザ情報登録
func Register(c echo.Context) error {
	request := new(requestRegister)
	if err := c.Bind(request); err != nil {
		failed := user.User{}
		return c.JSON(http.StatusInternalServerError, failed)
	}

	user := user.Create(request.Name)
	return c.JSON(http.StatusOK, user)
}

//Edit ...ユーザ情報編集
func Edit(c echo.Context) error {
	request := new(requestEdit)
	if err := c.Bind(request); err != nil {
		failed := user.User{}
		return c.JSON(http.StatusInternalServerError, failed)
	}

	user := user.Update(request.ID, request.NewName)
	return c.JSON(http.StatusOK, user)
}
