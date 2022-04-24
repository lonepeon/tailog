package tui

import "github.com/rivo/tview"

func NewFilterbox() *tview.InputField {
	return tview.NewInputField().
		SetLabel("Filters ").
		SetFieldWidth(0)
}
