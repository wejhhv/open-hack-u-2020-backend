package prefecturelistcontroller

import (
	"net/http"

	prefectureservice "github.com/KazuwoKiwame12/open-hack-u-2020-backend/Service/PrefectureService"
	"github.com/labstack/echo/v4"
)

//PrefectureInfoList ...各県の感情の色を返す
func PrefectureInfoList(c echo.Context) error {
	prefectureInfoList := prefectureservice.GetPrefectureInfoList()
	return c.JSON(http.StatusOK, prefectureInfoList)
}
