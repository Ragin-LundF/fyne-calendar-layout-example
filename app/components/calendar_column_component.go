package components

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"math/rand"
	"time"
)

type CalendarElement struct {
	RowName     string
	SubElements []string
}
type CalendarColumn struct {
	Elements []CalendarElement
}

func CreateCalendarColumn(columns CalendarColumn) fyne.CanvasObject {
	var columnAccordionContainer = widget.NewAccordionContainer()
	var accItem *widget.AccordionItem
	var accAppointment *widget.Label

	for _, element := range columns.Elements {
		if len(element.SubElements) > 0 {
			for _, subElement := range element.SubElements {
				accAppointment = widget.NewLabel(subElement)
				accItem = widget.NewAccordionItem(element.RowName, accAppointment)
			}
			columnAccordionContainer.Append(accItem)
		} else {
			accAppointment = widget.NewLabel(" ")
			accItem = widget.NewAccordionItem(element.RowName, accAppointment)
			columnAccordionContainer.Append(accItem)
		}
	}

	var scrollContainer = widget.NewScrollContainer(columnAccordionContainer)
	scrollContainer.SetMinSize(fyne.Size{
		Width:  200,
		Height: 400,
	})

	return scrollContainer
}

func CreateDay() fyne.CanvasObject {
	var calendarColumn = CalendarColumn{}
	var elements = calendarColumn.Elements
	for i := 0; i < 24; i++ {
		var element = CalendarElement{}
		var appointment = createAppointment()
		if appointment != nil {
			element.SubElements = []string{*appointment}
			element.RowName = fmt.Sprintf("%v:00   *", i)
		} else {
			element.RowName = fmt.Sprintf("%v:00", i)
		}
		elements = append(elements, element)
	}
	calendarColumn.Elements = elements

	return CreateCalendarColumn(calendarColumn)
}

func createAppointment() *string {
	appointments := []string{
		"fyne meeting",
		"private appointment",
		"shopping",
	}

	// simple rand to have random appointments
	rand.Seed(time.Now().UnixNano())
	var number = rand.Intn(len(appointments) + 5) // higher number for empty times
	if number < len(appointments) {
		return &appointments[number]
	}
	return nil
}
