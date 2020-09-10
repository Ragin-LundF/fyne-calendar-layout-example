package main

import (
	"calendar-example/app/menu"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

func main() {
	calendarApp := app.NewWithID("go_calendar_example")

	calendarWindow := calendarApp.NewWindow("Calendar Example")
	mainMenu := menu.CreateMainMenu(calendarApp, calendarWindow)

	calendarWindow.SetMainMenu(mainMenu)
	calendarWindow.SetMaster()

	tabs := menu.CreateTabMenu(calendarApp, calendarWindow)
	calendarWindow.SetContent(tabs)

	calendarWindow.Resize(fyne.Size{
		Width:  900,
		Height: 400,
	})
	calendarWindow.ShowAndRun()
}
