package prefectureservice

import (
	comment "github.com/KazuwoKiwame12/open-hack-u-2020-backend/DB/Model/Comment"
	commentservice "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Service/CommentService"
	emotionservice "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Service/EmotionService"
)

//PrefectureInfo ...クライアントに返すデータ
type PrefectureInfo struct {
	Prefecture string
	Color      string
}

//GetPrefectureInfoList ...各県の色と名前の情報取得
func GetPrefectureInfoList() []PrefectureInfo {
	prefectureInfoList := []PrefectureInfo{}
	for _, prefecture := range getAllPrefectures() {
		commentList := comment.GetListByPrefecture(prefecture)
		emotionIDList := commentservice.GetCommentEmotionIDList(commentList)
		prefectureColor := emotionservice.AVGColor(emotionIDList)

		prefectureInfo := PrefectureInfo{Prefecture: prefecture, Color: prefectureColor}
		prefectureInfoList = append(prefectureInfoList, prefectureInfo)
	}

	return prefectureInfoList
}

//getAllPrefectures ...全ての県名を取得
func getAllPrefectures() [47]string {
	prefectureList := [47]string{
		"HOKKAIDO",
		"AOMORI",
		"IWATE",
		"MIYAGI",
		"AKITA",
		"YAMAGATA",
		"FUKUSHIMA",
		"IBARAKI",
		"TOCHIGI",
		"GUNMA",
		"SAITAMA",
		"CHIBA",
		"TOKYO",
		"KANAGAWA",
		"NIIGATA",
		"TOYAMA",
		"ISHIKAWA",
		"FUKUI",
		"YAMANASHI",
		"NAGANO",
		"GIFU",
		"SHIZUOKA",
		"AICHI",
		"MIE",
		"SHIGA",
		"KYOTO",
		"OSAKA",
		"HYOGO",
		"NARA",
		"WAKAYAMA",
		"TOTTORI",
		"SHIMANE",
		"OKAYAMA",
		"HIROSHIMA",
		"YAMAGUCHI",
		"TOKUSHIMA",
		"KAGAWA",
		"EHIME",
		"KOCHI",
		"OKINAWA",
		"FUKUOKA",
		"SAGA",
		"NAGASAKI",
		"KUMAMOTO",
		"OITA",
		"MIYAZAKI",
		"KAGOSHIMA",
	}

	return prefectureList
}
