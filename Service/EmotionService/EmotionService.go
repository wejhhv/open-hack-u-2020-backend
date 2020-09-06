package emotionservice

import (
	"image/color"
	"strconv"
)

//Emotion ...Emotionデータ
type Emotion struct {
	ID    int
	Color color.NRGBA
	Point int
}

//AlfaBlend ...α値を用いた色合成の計算
func AlfaBlend(emotionIDList []int) string {
	mixColor := color.NRGBA{R: 0, G: 0, B: 0, A: 160}
	emotionList := getEmotionList()

	for _, emotionID := range emotionIDList {
		alfaRate := float32(emotionList[emotionID].Color.A / 255)

		newR := float32(emotionList[emotionID].Color.R)*alfaRate + float32(mixColor.R)*(1-alfaRate)
		newG := float32(emotionList[emotionID].Color.G)*alfaRate + float32(mixColor.G)*(1-alfaRate)
		newB := float32(emotionList[emotionID].Color.B)*alfaRate + float32(mixColor.B)*(1-alfaRate)
		mixColor = color.NRGBA{R: uint8(newR), G: uint8(newG), B: uint8(newB), A: 160}
	}

	return "#" + convertToHexaDecimalNumColor(mixColor)
}

//AVGColor ...RGBの値をそれぞれの平均を取る
func AVGColor(emotionIDList []int) string {
	mixColor := color.NRGBA{R: 255, G: 255, B: 255, A: 160}
	emotionList := getEmotionList()

	for _, emotionID := range emotionIDList {
		newR := (emotionList[emotionID].Color.R + mixColor.R) / 2
		newG := (emotionList[emotionID].Color.G + mixColor.G) / 2
		newB := (emotionList[emotionID].Color.B + mixColor.B) / 2
		mixColor = color.NRGBA{R: uint8(newR), G: uint8(newG), B: uint8(newB), A: 160}
	}

	return "#" + convertToHexaDecimalNumColor(mixColor)
}

func getEmotionList() [12]Emotion {
	colorList := [4]color.NRGBA{}
	colorList[0] = color.NRGBA{R: 255, G: 102, B: 255, A: 160} //喜
	colorList[1] = color.NRGBA{R: 255, G: 51, B: 51, A: 160}   //怒
	colorList[2] = color.NRGBA{R: 153, G: 153, B: 255, A: 160} //哀
	colorList[3] = color.NRGBA{R: 204, G: 204, B: 51, A: 160}  //楽

	emotionList := [12]Emotion{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			emotionID := (i * 3) + j
			emotionList[emotionID] = Emotion{ID: emotionID, Color: colorList[i], Point: emotionID * 5}
		}
	}

	return emotionList
}

func convertToHexaDecimalNumColor(col color.NRGBA) string {
	stringR := strconv.FormatInt(int64(col.R), 16)
	stringG := strconv.FormatInt(int64(col.G), 16)
	stringB := strconv.FormatInt(int64(col.B), 16)

	if col.R < 16 {
		stringR = "0" + stringR
	}
	if col.G < 16 {
		stringG = "0" + stringG
	}
	if col.B < 16 {
		stringB = "0" + stringB
	}

	return stringR + stringG + stringB
}
