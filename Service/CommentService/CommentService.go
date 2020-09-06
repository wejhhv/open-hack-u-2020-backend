package commentservice

import (
	"time"

	comment "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Comment"
	user "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/User"
)

//CommentForClient ...クライアントに渡すデータ構造
type CommentForClient struct {
	ID         int
	EmotionID  int
	Latitude   float32
	Longtitude float32
	UserName   string
	DateTime   time.Time
}

//CommentContent ...コメントの内容
type CommentContent struct {
	Comment   string
	EmotionID int
	UserName  string
	DateTime  time.Time
}

//ConvertCtoCFC ...[]commentを[]CommentForClientに変換する
func ConvertCtoCFC(commentList []comment.Comment) []CommentForClient {
	commentListForClient := []CommentForClient{}

	for _, comment := range commentList {
		user := user.Get(comment.UserID)
		commentForClient := CommentForClient{}

		commentForClient.ID = comment.ID
		commentForClient.EmotionID = comment.EmotionID
		commentForClient.Latitude = comment.Latitude
		commentForClient.Longtitude = comment.Longtitude
		commentForClient.UserName = user.Name
		commentForClient.DateTime = comment.DateTime

		commentListForClient = append(commentListForClient, commentForClient)
	}

	return commentListForClient
}

//MakeCommentContent ...CommentをCommentContentに整形する
func MakeCommentContent(id int) CommentContent {
	comment := comment.GetListByID(id)
	user := user.Get(comment.UserID)

	commentContent := CommentContent{}
	commentContent.Comment = comment.Comment
	commentContent.EmotionID = comment.EmotionID
	commentContent.UserName = user.Name
	commentContent.DateTime = comment.DateTime

	return commentContent
}

//GetCommentEmotionIDList ...受け取った[]CommentのEmotionのみを返す
func GetCommentEmotionIDList(commentList []comment.Comment) []int {
	emotionList := make([]int, 0)
	for _, comment := range commentList {
		emotionList = append(emotionList, comment.EmotionID)
	}

	return emotionList
}
