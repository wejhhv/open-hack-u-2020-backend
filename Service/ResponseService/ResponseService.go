package responseservice

import (
	"time"

	response "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Response"
	user "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/User"
)

//ResponseForClient ...クライアントに渡すデータ構造
type ResponseForClient struct {
	ID       int
	UserName string
	Comment  string
	DateTime time.Time
}

//ConvertRtoRFC ...[]Responseを[]ResponseForClientに変換する
func ConvertRtoRFC(responseList []response.Response) []ResponseForClient {
	responseListForClient := []ResponseForClient{}

	for _, response := range responseList {
		user := user.Get(response.UserID)

		responseForClient := ResponseForClient{}
		responseForClient.ID = response.ID
		responseForClient.UserName = user.Name
		responseForClient.Comment = response.Comment
		responseForClient.DateTime = response.DateTime

		responseListForClient = append(responseListForClient, responseForClient)
	}

	return responseListForClient
}
