package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewFilterbox(onChange func(string)) *tview.InputField {
	field := tview.NewInputField().
		SetLabel("Filters ").
		SetFieldWidth(0)
	field.SetLabelColor(tcell.ColorDefault)
	field.SetBackgroundColor(tcell.ColorDefault)
	field.SetChangedFunc(onChange)

	return field
}
