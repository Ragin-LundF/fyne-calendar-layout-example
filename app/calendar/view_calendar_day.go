package calendar

import (
	"calendar-example/app/components"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// ViewCalendarDay shows the current calendar as day view
func ViewCalendarDay() fyne.CanvasObject {
	var dayColumn = components.CreateDay()

	var calendarHBox = widget.NewVBox(
		widget.NewLabel("Monday"),
		dayColumn)

	return widget.NewVBox(
		widget.NewLabelWithStyle("Calendar - Day", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		components.CreateToolbar(),
		layout.NewSpacer(),
		calendarHBox,
		components.CreateStatusBar(),
	)
}
