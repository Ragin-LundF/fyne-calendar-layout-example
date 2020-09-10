package menu

import (
	"calendar-example/app/calendar"
	"calendar-example/app/constants"
	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func CreateMainMenu(app fyne.App, window fyne.Window) *fyne.MainMenu {
	quitItem := fyne.NewMenuItem("Quit", func() { app.Quit() })

	darkThemeItem := fyne.NewMenuItem("Dark Theme", func() {
		app.Settings().SetTheme(theme.DarkTheme())
		app.Preferences().SetString(constants.PreferencesTheme, constants.PreferencesThemeDark)
	})
	lightThemeItem := fyne.NewMenuItem("Light Theme", func() {
		app.Settings().SetTheme(theme.LightTheme())
		app.Preferences().SetString(constants.PreferencesTheme, constants.PreferencesThemeLight)
	})

	mainMenu := fyne.NewMainMenu(
		// a quit item will be appended to our first menu
		fyne.NewMenu("Calendar", fyne.NewMenuItemSeparator(), quitItem),
		fyne.NewMenu("Theme", darkThemeItem, lightThemeItem),
	)

	return mainMenu
}

// CreateTabMenu creates the main tab menu
func CreateTabMenu(calendarApp fyne.App, calendarWindow fyne.Window) (tabs *widget.TabContainer) {
	tabs = widget.NewTabContainer(
		widget.NewTabItemWithIcon("Day", theme.FolderIcon(), calendar.ViewCalendarDay()),
		widget.NewTabItemWithIcon("Week", theme.FileIcon(), calendar.ViewCalendarWeek()),
		widget.NewTabItemWithIcon("Appointments", theme.DocumentIcon(), calendar.ViewCalendarAppointment()),
	)
	tabs.SetTabLocation(widget.TabLocationLeading)

	return tabs
}
