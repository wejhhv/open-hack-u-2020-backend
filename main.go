package main

import (
	"net/http"
	"os"

	commentcontroller "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Controller/CommentController"
	commentlistcontroller "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Controller/CommentListController"
	prefecturelistcontroller "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Controller/PrefectureListController"
	responsecontroller "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Controller/ResponseController"
	responselistcontroller "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Controller/ResponseListController"
	usercontroller "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Controller/UserController"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", helloWorld)

	//感情関連のデータ取得API
	e.GET("/emotion/prefectures", prefecturelistcontroller.PrefectureInfoList)
	e.GET("/emotion/:prefecture/comments", commentlistcontroller.CommentListInPrefecture)
	e.GET("/emotion/mycomments/:user_id/:page_num/:page_size", commentlistcontroller.MyCommentList)
	e.GET("/emotion/comments/:comment_id", responselistcontroller.ResponseListInComment)

	//コメント登録と削除API
	e.POST("/comment/register", commentcontroller.Register)
	e.DELETE("/comment/delete", commentcontroller.Delete)

	//コメントに対する返信登録と削除API
	e.POST("/comment/response/register", responsecontroller.Register)
	e.DELETE("/comment/response/delete", responsecontroller.Delete)

	//ユーザ登録と編集API
	e.POST("/user/register", usercontroller.Register)
	e.PATCH("/user/edit", usercontroller.Edit)

	e.Logger.Fatal(e.Start(":" + port))
}

func helloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World!!")
}
