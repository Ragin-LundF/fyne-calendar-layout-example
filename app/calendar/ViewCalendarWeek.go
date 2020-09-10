package calendar

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"k8s-management-go/app/components"
)

// ViewCalendarWeek shows the current calendar as weekly view
func ViewCalendarWeek() fyne.CanvasObject {
	var mondayColumn = components.CreateDay()
	var tuesdayColumn = components.CreateDay()
	var wednesdayColumn = components.CreateDay()
	var thursdayColumn = components.CreateDay()
	var fridayColumn = components.CreateDay()

	var calendarHBox = widget.NewHBox(
		widget.NewVBox(
			widget.NewLabel("Monday"),
			mondayColumn),
		widget.NewVBox(
			widget.NewLabel("Tuesday"),
			tuesdayColumn),
		widget.NewVBox(
			widget.NewLabel("Wednesday"),
			wednesdayColumn),
		widget.NewVBox(
			widget.NewLabel("Thursday"),
			thursdayColumn),
		widget.NewVBox(
			widget.NewLabel("Friday"),
			fridayColumn),
	)

	return widget.NewVBox(
		widget.NewLabelWithStyle("Calendar - Week", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		components.CreateToolbar(),
		layout.NewSpacer(),
		calendarHBox,
		components.CreateStatusBar(),
	)
}
