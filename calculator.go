package main

import (
	"strconv"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalc() {
	w := myapp.NewWindow("Calculator")
	input := widget.NewLabel("Do The Calculations Here")
	output := ""
	var historyArr []string
	historyStr := ""
	history := widget.NewLabel(historyStr)
	ishistoryClick := false
	historybtn := widget.NewButton("History", func() {
		if ishistoryClick {
			historyStr = ""
			history.SetText(historyStr)
		} else {
			for i := 0; i < len(historyArr); i++ {
				historyStr = historyStr + historyArr[i] + "\n"
				history.SetText(historyStr)
			}
		}
		ishistoryClick = !ishistoryClick
	})
	backbtn := widget.NewButton("Back", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})
	clearbtn := widget.NewButton("Clear", func() {
		output = ""
		input.SetText(output)
		historyStr = ""
		history.SetText(historyStr)
	})
	btnBB := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})
	btnSB := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})
	btnDiv := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})

	digSeven := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})
	digEight := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})
	digNine := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})
	btnMul := widget.NewButton("*", func() {
		output = output + "*"
		input.SetText(output)
	})
	digFour := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})
	digFive := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})
	digSix := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})
	btnPlus := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})
	digOne := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})
	digTwo := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})
	digThree := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})
	btnMinus := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})
	digZero := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})
	dotBtn := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})
	btnEqual := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				stringToappend := output + " = "
				output = strconv.FormatFloat(result.(float64), 'f', -1, 64)
				stringToappend = stringToappend + output
				input.SetText(output)
				historyArr = append(historyArr, stringToappend)
			} else {
				output = "There Is Some Error"
				input.SetText(output)
			}

		} else {
			output = "There Is Some Error"
			input.SetText(output)
		}
	})
	w.SetContent(container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historybtn,
				backbtn,
			),
			container.NewGridWithColumns(4,
				clearbtn,
				btnBB,
				btnSB,
				btnDiv,
			),
			container.NewGridWithColumns(4,
				digSeven,
				digEight,
				digNine,
				btnMul,
			),
			container.NewGridWithColumns(4,
				digFour,
				digFive,
				digSix,
				btnPlus,
			),
			container.NewGridWithColumns(4,
				digOne,
				digTwo,
				digThree,
				btnMinus,
			),
			container.NewGridWithColumns(3,
				digZero,
				dotBtn,
				btnEqual,
			),
		),
	))
	w.Show()
}
