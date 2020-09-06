package comment

import (
	"time"

	db "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB"
	dateservice "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Service/DateService"
	"github.com/jinzhu/gorm"
)

//Comment ...table: Commentsのモデル
type Comment struct {
	ID         int
	UserID     int
	EmotionID  int
	Prefecture string
	Latitude   float32
	Longtitude float32
	Comment    string
	DateTime   time.Time
}

//Create ...Commentsモデルの保存
func Create(comment Comment) bool {
	db := db.Connect()
	defer db.Close()

	result := db.Create(&comment)
	if result.Error != nil {
		return false
	}
	return true
}

//Get ...Commentsモデルの取得
func Get(id int) *Comment {
	db := db.Connect()
	defer db.Close()

	comment := Comment{}
	db.First(&comment, id)

	return &comment
}

//Delete ...Commentsモデルの削除
func Delete(id int) bool {
	db := db.Connect()
	defer db.Close()

	comment := Comment{}
	comment.ID = id
	result := db.Delete(&comment)
	if result.Error != nil {
		return false
	}
	return true
}

//GetListByID ...指定したIDのコメントを返す
func GetListByID(commentID int) Comment {
	db := db.Connect()
	defer db.Close()

	now := time.Now()
	from := dateservice.StartOfDay(now)
	to := dateservice.StartOfNextDay(now)

	comment := Comment{}
	db.Where("id = ? AND date_time >= ? AND date_time < ?", commentID, from, to).First(&comment)
	return comment
}

//GetListByPrefecture ...指定した県のコメント一覧
func GetListByPrefecture(prefecture string) []Comment {
	db := db.Connect()
	defer db.Close()

	now := time.Now()
	from := dateservice.StartOfDay(now)
	to := dateservice.StartOfNextDay(now)

	comments := []Comment{}
	db.Where("prefecture = ? AND date_time >= ? AND date_time < ?", prefecture, from, to).Find(&comments)
	return comments
}

//GetListByUserWithPaginate ...指定したユーザのコメント一覧(ぺジネーションで)
func GetListByUserWithPaginate(userID int, pageNum int, pageSize int) []Comment {
	db := db.Connect()
	defer db.Close()

	comments := []Comment{}
	db.Scopes(paginate(pageNum, pageSize)).Where("user_id = ?", userID).Find(&comments)
	return comments
}

func paginate(num int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := num
		if page == 0 {
			page = 1
		}

		pageSize := size
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
