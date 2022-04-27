package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func NewFilterbox() *tview.InputField {
	field := tview.NewInputField().
		SetLabel("Filters ").
		SetFieldWidth(0)
	field.SetLabelColor(tcell.ColorDefault)
	field.SetBackgroundColor(tcell.ColorDefault)

	return field
}
