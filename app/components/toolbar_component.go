package components

import (
	"fmt"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func CreateToolbar() *widget.Toolbar {
	return widget.NewToolbar(widget.NewToolbarAction(theme.MailComposeIcon(), func() { CreateNewAppointment() }),
		widget.NewToolbarSeparator(),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.NavigateBackIcon(), func() { fmt.Println("Previous") }),
		widget.NewToolbarAction(theme.HomeIcon(), func() { fmt.Println("Home") }),
		widget.NewToolbarAction(theme.NavigateNextIcon(), func() { fmt.Println("Next") }),
	)
}
