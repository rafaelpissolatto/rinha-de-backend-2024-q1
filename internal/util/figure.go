package util

import (
	"github.com/common-nighthawk/go-figure"
)

func Figure() {
	titleApiFigure := figure.NewColorFigure("RINHADEV2024/Q1", "", "purple", true)
	authorFigure := figure.NewColorFigure("Rafael Pissolatto Nunes", "", "white", true)
	titleApiFigure.Print()
	authorFigure.Print()
	println()
}
