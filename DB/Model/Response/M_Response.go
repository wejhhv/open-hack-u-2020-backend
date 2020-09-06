package response

import (
	"time"

	db "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB"
)

//Response ...table: Responsesのモデル
type Response struct {
	ID        int
	UserID    int
	CommentID int
	Comment   string
	DateTime  time.Time
}

//Create ...Responsesモデルの保存
func Create(response Response) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&response)
	if result.Error != nil {
		return false
	}
	return true
}

//Get ...Responsesモデルの取得
func Get(id int) *Response {
	db := db.Connect()
	defer db.Close()

	response := Response{}
	db.First(&response, id)

	return &response
}

//Delete ...Responsesモデルの削除
func Delete(id int) bool {
	db := db.Connect()
	defer db.Close()

	response := Response{}
	response.ID = id
	result := db.Delete(&response)
	if result.Error != nil {
		return false
	}
	return true
}

//GetListByComment ...コメントに紐付いたRresponseモデルの取得
func GetListByComment(commentID int) []Response {
	db := db.Connect()
	defer db.Close()

	responseList := []Response{}
	db.Where("comment_id = ?", commentID).Find(&responseList)
	return responseList
}
