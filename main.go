package main

import (
	"image/color"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fogleman/gg"
)

var (
	backgroundImageFilename = "template.png"
	outputFilename          = "out.png"
)

// UnifyNewline make newlines one.
func UnifyNewline(str string) string {
	rep := regexp.MustCompile(`\n+`)
	s := rep.ReplaceAllString(str, "\n")
	return s
}

// WrapText
func WrapText(str string, point int) []string {
	runes := []rune(str)
	var slice []string
	for i := 0; i < len(runes); i += point {
		if i+point < len(runes) {
			slice = append(slice, string(runes[i:(i+point)]))
		} else {
			slice = append(slice, string(runes[i:]))
		}
	}
	return slice
}

func FmtText(str string) []string {
	arr := strings.Split(str, "\n")
	var wrapped []string
	for _, v := range arr {
		wrapped = append(wrapped, WrapText(v, 21)...)
	}
	if len(wrapped) >= 9 {
		// 9行目の2文字以降を切り捨てる
		wrapped[8] = string([]rune(wrapped[8])[0])
		// 10行目以降を切り捨てる
		wrapped = wrapped[:9]
	}

	return wrapped
}

func main() {
	s := strings.Repeat("ブルータスお前もか", 19)
	s = UnifyNewline(s)
	lines := FmtText(s)

	backgroundImage, err := gg.LoadImage(backgroundImageFilename)
	if err != nil {
		panic(err)
	}
	dc := gg.NewContext(600, 315)
	dc.DrawImage(backgroundImage, 0, 0)

	fontPath := filepath.Join("Kosugi-Regular.ttf")
	if err := dc.LoadFontFace(fontPath, 24); err != nil {
		panic(err)
	}
	dc.SetColor(color.White)

	marginX := 30.0
	marginY := 5.0
	var x, y, padX, padY float64
	for i, v := range lines {
		textWidth, textHeight := dc.MeasureString(v)
		padX = (float64(dc.Width()) - marginX*2 - textWidth) / 2
		padY = (float64(dc.Height()) - marginY*2 - (textHeight+10)*float64(len(lines))) / 2
		x = marginX + padX
		y = marginY + padY + (textHeight+10)*float64(i)
		dc.DrawString(v, x, y)
	}

	if err := dc.SavePNG(outputFilename); err != nil {
		panic(err)
	}
}
