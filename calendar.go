package main

import (
	"calendar-example/app/menu"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

func main() {
	var calendarApp = app.NewWithID("go_calendar_example")

	var calendarWindow = calendarApp.NewWindow("Calendar Example")
	var mainMenu = menu.CreateMainMenu(calendarApp, calendarWindow)

	calendarWindow.SetMainMenu(mainMenu)
	calendarWindow.SetMaster()

	var tabs = menu.CreateTabMenu()
	calendarWindow.SetContent(tabs)

	calendarWindow.Resize(fyne.Size{
		Width:  900,
		Height: 400,
	})
	calendarWindow.ShowAndRun()
}
