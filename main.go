package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//github.com/fyne-io/examples/blob/53445faa6dcc/clock/clock.go
var myapp fyne.App = app.New()
var myWindow fyne.Window = myapp.NewWindow("Milan OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var panelContent *fyne.Container

func main() {

	clockOn := false
	img := canvas.NewImageFromFile("Wallpaper.jpg")
	img.Resize(fyne.NewSize(1260, 980))
	myWindow.Resize(fyne.NewSize(1440, 1080))
	btn1 = widget.NewButton("Calculator", func() {
		showCalc()
	})
	btn2 = widget.NewButton("Weather", func() {
		showWeatherApp()
	})
	btn3 = widget.NewButton("Gallery", func() {
		showGalleryApp()
	})
	btn4 = widget.NewButton("Text Editor", func() {
		showTextEditor()
	})
	btn5 = widget.NewButton("C\nl\no\nc\nk", func() {
		if clockOn == false {
			Show()
			clockOn = true
		} else {
			clockClose()
			clockOn = false
		}
	})
	top := widget.NewLabelWithStyle(" ", fyne.TextAlignCenter, fyne.TextStyle{
		Bold:      true,
		Monospace: true,
	})
	bottom := container.NewGridWithColumns(
		5,
		btn1,
		btn2,
		btn3,
		btn4,
	)
	side2 := widget.NewLabel(" ")
	myWindow.SetContent(container.New(layout.NewMaxLayout(),
		img,
		container.NewBorder(
			top,
			bottom,
			side2,
			btn5,
		),
	))
	myWindow.ShowAndRun()
}
