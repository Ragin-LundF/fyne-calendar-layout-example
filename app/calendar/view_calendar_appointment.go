package calendar

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// ViewCalendarAppointment shows the appointment overview
func ViewCalendarAppointment() fyne.CanvasObject {
	return widget.NewVBox(
		widget.NewLabelWithStyle("Appointments", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
	)
}
