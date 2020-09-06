package responsecontroller

import (
	"net/http"
	"time"

	response "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Response"
	"github.com/labstack/echo/v4"
)

type requestRegister struct {
	UserID    int    `json:"user_id"`
	CommentID int    `json:"comment_id"`
	Comment   string `json:"comment"`
}

type requestDelete struct {
	ResponseID int `json:"response_id"`
}

//Register ...コメントに対する返信登録
func Register(c echo.Context) error {
	request := new(requestRegister)
	if err := c.Bind(request); err != nil {
		returnValue := map[string]bool{"hasSuccess": false}
		return c.JSON(http.StatusInternalServerError, returnValue)
	}

	//構造体宣言
	var res response.Response
	//構造体に入れる
	res.UserID = request.UserID
	res.CommentID = request.CommentID
	res.Comment = request.Comment
	res.DateTime = time.Now()

	//DB処理
	hasSuccess := response.Create(res)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}
	return c.JSON(http.StatusOK, returnValue)
}

//Delete ...返信の削除
func Delete(c echo.Context) error {
	request := new(requestDelete)
	if err := c.Bind(request); err != nil {
		returnValue := map[string]bool{"hasSuccess": false}
		return c.JSON(http.StatusInternalServerError, returnValue)
	}

	//DB処理
	hasSuccess := response.Delete(request.ResponseID)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}

	return c.JSON(http.StatusOK, returnValue)
}
