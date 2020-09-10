package components

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func CreateToolbar() *widget.Toolbar {
	return widget.NewToolbar(widget.NewToolbarAction(theme.MailComposeIcon(), func() { createNewAppointment() }),
		widget.NewToolbarSeparator(),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() { fmt.Println("Previous") }),
		widget.NewToolbarAction(theme.HomeIcon(), func() { fmt.Println("Home") }),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() { fmt.Println("Next") }),
	)
}

func createNewAppointment() {
	var windowAppointment = fyne.CurrentApp().NewWindow("Create new appointment")
	var categories = []string{
		"private", "business", "customer",
	}

	var form = &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Title", Widget: widget.NewEntry()},
			{Text: "Category", Widget: widget.NewSelect(categories, func(changed string) {})},
			{Text: "Description", Widget: widget.NewMultiLineEntry()},
			{Text: "Duration (minutes)", Widget: createDurationWidgets()},
			{Text: "Private", Widget: widget.NewCheck(" ", func(b bool) {})},
			{Text: "Notification", Widget: createNotificationRadioBoxes()},
		},
		OnCancel: func() { windowAppointment.Close() },
		OnSubmit: func() { windowAppointment.Close() },
	}

	var box = widget.NewVBox(widget.NewToolbar(widget.NewToolbarAction(theme.CancelIcon(), func() { windowAppointment.Close() })), form)

	windowAppointment.SetContent(box)
	windowAppointment.Resize(fyne.NewSize(700, 500))
	windowAppointment.Show()
}

func createDurationWidgets() *widget.Box {
	var durationEntry = widget.NewEntry()
	durationEntry.SetText("15")
	var durationSlider = widget.NewSlider(15, 400)
	durationSlider.OnChanged = func(number float64) {
		durationEntry.SetText(fmt.Sprintf("%v", int(number)))
	}
	return widget.NewVBox(
		durationEntry,
		durationSlider)
}

func createNotificationRadioBoxes() *widget.Radio {
	var notification = widget.NewRadio([]string{"none", "15 minutes", "30 minutes"}, func(s string) {})
	notification.SetSelected("none")

	return notification
}
