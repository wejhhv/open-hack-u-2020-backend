package commentcontroller

import (
	"net/http"
	"time"

	comment "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Comment"
	"github.com/labstack/echo/v4"
)

type requestRegister struct {
	UserID     int     `json:"user_id"`
	EmotionID  int     `json:"emotion_id"`
	Comment    string  `json:"comment"`
	Latitude   float32 `json:"latitude"`
	Longtitude float32 `json:"longtitude"`
	Prefecture string  `json:"prefecture"`
}

type requestDelete struct {
	CommentID int `json:"comment_id"`
}

//Register ...コメント登録
func Register(c echo.Context) error {
	request := new(requestRegister)
	if err := c.Bind(request); err != nil {
		returnValue := map[string]bool{"hasSuccess": false}
		return c.JSON(http.StatusInternalServerError, returnValue)
	}

	var com comment.Comment
	//構造体に入れる
	com.UserID = request.UserID
	com.EmotionID = request.EmotionID
	com.Comment = request.Comment
	com.Latitude = request.Latitude
	com.Longtitude = request.Longtitude
	com.Prefecture = request.Prefecture
	com.DateTime = time.Now()

	//DB処理
	hasSuccess := comment.Create(com)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}
	return c.JSON(http.StatusOK, returnValue)
}

//Delete ...コメントの削除
func Delete(c echo.Context) error {
	request := new(requestDelete)
	if err := c.Bind(request); err != nil {
		returnValue := map[string]bool{"hasSuccess": false}
		return c.JSON(http.StatusInternalServerError, returnValue)
	}

	//DB処理
	hasSuccess := comment.Delete(request.CommentID)
	returnValue := map[string]bool{"hasSuccess": hasSuccess}
	return c.JSON(http.StatusOK, returnValue)
}
