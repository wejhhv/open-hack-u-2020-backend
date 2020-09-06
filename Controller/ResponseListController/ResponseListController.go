package responselistcontroller

import (
	"net/http"
	"strconv"

	response "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Response"
	commentservice "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Service/CommentService"
	responseservice "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Service/ResponseService"
	"github.com/labstack/echo/v4"
)

//CommentAndResponses ...クライアントに返すデータ
type CommentAndResponses struct {
	CommentContent commentservice.CommentContent
	Response       []responseservice.ResponseForClient
}

//ResponseListInComment ...コメントに紐付いた返信データを整形して返す
func ResponseListInComment(c echo.Context) error {
	commentID, _ := strconv.Atoi(c.Param("comment_id"))
	commentContent := commentservice.MakeCommentContent(commentID)

	if commentContent.UserName == "" {
		nullValue := []responseservice.ResponseForClient{}
		failed := CommentAndResponses{CommentContent: commentContent, Response: nullValue}
		return c.JSON(http.StatusOK, failed)
	}

	responseList := response.GetListByComment(commentID)
	responseListForClient := responseservice.ConvertRtoRFC(responseList)

	cars := CommentAndResponses{}
	cars.CommentContent = commentContent
	cars.Response = responseListForClient

	return c.JSON(http.StatusOK, cars)
}
