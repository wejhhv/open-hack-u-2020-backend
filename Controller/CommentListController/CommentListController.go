package commentlistcontroller

import (
	"net/http"
	"strconv"

	comment "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Comment"
	commentservice "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Service/CommentService"
	"github.com/labstack/echo/v4"
)

//CommentListInPrefecture ...県に存在するコメント一覧取得
func CommentListInPrefecture(c echo.Context) error {
	commentList := comment.GetListByPrefecture(c.Param("prefecture"))
	commentListForClient := commentservice.ConvertCtoCFC(commentList)

	return c.JSON(http.StatusOK, commentListForClient)
}

//MyCommentList ...自分のコメント一覧取得
func MyCommentList(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	pageNum, _ := strconv.Atoi(c.Param("page_num"))
	pageSize, _ := strconv.Atoi(c.Param("page_size"))

	commentList := comment.GetListByUserWithPaginate(userID, pageNum, pageSize)
	commentListForClient := commentservice.ConvertCtoCFC(commentList)

	return c.JSON(http.StatusOK, commentListForClient)
}
