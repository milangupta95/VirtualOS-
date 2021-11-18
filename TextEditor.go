package main

import (
	// "fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/layout"
	"io/ioutil"
	"strconv"
	"strings"
)

func showTextEditor() {
	var Tabcount int = 1
	var TabName string = "tab" + strconv.Itoa(Tabcount)
	var tabSet = make(map[string]*widget.Entry)
	w := myapp.NewWindow("Text Editor")
	w.Resize(fyne.NewSize(640, 480))
	headerLabel := widget.NewLabelWithStyle("Text Editor", fyne.TextAlignCenter, fyne.TextStyle{Monospace: true, Bold: true})
	tabSet[TabName] = widget.NewMultiLineEntry()
	tabs := container.NewAppTabs(
		container.NewTabItem("NewFile "+strconv.Itoa(Tabcount), tabSet[TabName]),
	)
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			Tabcount = Tabcount + 1
			TabName = "tab" + strconv.Itoa(Tabcount)
			tabSet[TabName] = widget.NewMultiLineEntry()
			tabs.Append(container.NewTabItem("NewFile "+strconv.Itoa(Tabcount), tabSet[TabName]))
			tabSet[TabName].CreateRenderer()
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.FolderOpenIcon(), func() {
			openFileDialog := dialog.NewFileOpen(
				func(r fyne.URIReadCloser, _ error) {
					ReadData, _ := ioutil.ReadAll(r)
					output := fyne.NewStaticResource("New File", ReadData)
					viewData := widget.NewMultiLineEntry()
					viewData.SetText(string(output.StaticContent))
					nw := fyne.CurrentApp().NewWindow(string(output.StaticName))
					nw.SetContent(container.NewScroll(viewData))
					nw.Resize(fyne.NewSize(400, 400))
					nw.Show()
				}, w)
			openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
			openFileDialog.Show()
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			tabIndex := strings.Split(tabs.CurrentTab().Text, " ")[1]
			textToAppend := tabSet["tab"+string(tabIndex)].Text
			saveFiledialog := dialog.NewFileSave(
				func(uc fyne.URIWriteCloser, _ error) {
					bytesToSave := []byte(textToAppend)
					uc.Write(bytesToSave)
				}, w)
			saveFiledialog.SetFileName(tabs.CurrentTab().Text + ".txt")
			saveFiledialog.Show()
		}),
	)
	// render:= ((tabs.CurrentTab()).Content)
	//fmt.Printf("%T",render)
	// fmt.Println(render.Size)
	w.SetContent(container.NewVBox(
		container.NewGridWithColumns(6,
			headerLabel,
			toolbar,
		),
		container.NewVBox(tabs),
	),
	)

	w.Show()
}
