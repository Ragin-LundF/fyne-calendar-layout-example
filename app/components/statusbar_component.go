package components

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func CreateStatusBar() *widget.Box {
	var progressBar = widget.NewProgressBar()
	progressBar.Resize(fyne.NewSize(400, 1))
	progressBar.SetValue(100)

	return widget.NewHBox(
		widget.NewLabel("Status..."),
		layout.NewSpacer(),
		progressBar,
	)
}
