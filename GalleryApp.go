package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	// "strings"
	"io/ioutil"
	// "os"
	"path/filepath"
)

func showGalleryApp() {
	picIndex := 0
	var prevButton *widget.Button
	var nextButton *widget.Button
	var picsArr []string
	gapp := myapp.NewWindow("Gallery")
	imgCanvas := gapp.Canvas()
	gapp.Resize(fyne.NewSize(1280, 960))
	root_src := "PictureGallery"
	files, err := ioutil.ReadDir(root_src)
	spaceWidget := widget.NewLabelWithStyle(" ", fyne.TextAlignCenter, fyne.TextStyle{Monospace: true, Bold: true})
	spaceWidget1 := widget.NewLabelWithStyle(" ", fyne.TextAlignCenter, fyne.TextStyle{Monospace: true, Bold: true})
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			extension := filepath.Ext(file.Name())
			if extension == ".png" || extension == ".jpg" || extension == ".jpeg" {
				picsArr = append(picsArr, file.Name())
			}
		}
	}
	imageCanvas := canvas.NewImageFromFile(root_src + "/" + picsArr[picIndex])
	var imageLayout *fyne.Container
	textFileName := picsArr[0]
	prevButton = widget.NewButtonWithIcon(" ", theme.NavigateBackIcon(), func() {
		if picIndex > 0 {
			picIndex--
			imageCanvas = canvas.NewImageFromFile(root_src + "/" + picsArr[picIndex])
			imageLayout = container.New(layout.NewPaddedLayout(),
				imageCanvas,
				container.NewBorder(
					spaceWidget,
					spaceWidget1,
					prevButton,
					nextButton,
				),
			)
			imgCanvas.SetContent(imageLayout)
			spaceWidget1.SetText(picsArr[picIndex])
		}
	})
	nextButton = widget.NewButtonWithIcon(" ", theme.NavigateNextIcon(), func() {
		if picIndex < len(picsArr)-1 {
			picIndex++
			imageCanvas = canvas.NewImageFromFile(root_src + "/" + picsArr[picIndex])
			imageLayout = container.New(layout.NewPaddedLayout(),
				imageCanvas,
				container.NewBorder(
					spaceWidget,
					spaceWidget1,
					prevButton,
					nextButton,
				),
			)
			imgCanvas.SetContent(imageLayout)
			spaceWidget1.SetText(picsArr[picIndex])

		}
	})
	spaceWidget1.SetText(textFileName)
	// imageCanvas.Resize(fyne.NewSize(1280,720))
	// imageCanvas.FillMode = canvas.ImageFillContain
	// imageCanvas.Resize(fyne.NewSize(1080,640))
	// imageCanvas.FillMode = canvas.ImageFillOriginal
	imageLayout = container.New(layout.NewPaddedLayout(),
		imageCanvas,
		container.NewBorder(
			spaceWidget,
			spaceWidget1,
			prevButton,
			nextButton,
		),
	)
	fmt.Printf("%T", imageLayout)
	fmt.Printf("%T", prevButton)
	// w.SetContent(imageLayout)
	imgCanvas.SetContent(imageLayout)
	gapp.Show()
}
