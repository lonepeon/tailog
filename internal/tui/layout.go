package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/lonepeon/tailog/internal/decoding"
	"github.com/lonepeon/tailog/internal/filterlang"
	"github.com/rivo/tview"
)

type Layout struct {
	*tview.Grid

	errormsg  *tview.TextView
	filterbox *tview.InputField
	table     *Table
}

func NewLayout(buffer int, fieldNames []string) *Layout {
	table := NewTable(buffer, fieldNames)
	errormsg := tview.NewTextView()
	errormsg.SetTextColor(tcell.ColorRed).SetBackgroundColor(tcell.ColorDefault)

	filterbox := NewFilterbox(func(txt string) {
		errormsg.Clear()

		if txt == "" {
			table.FilterFunc(decoding.KeepAll)
			return
		}

		interpreter, err := filterlang.Parse(txt)
		if err != nil {
			errormsg.SetText(err.Error())
			return
		}

		table.FilterFunc(interpreter.Execute)
	})

	grid := tview.NewGrid().
		SetRows(1, 1, 0).
		SetBorders(true).
		AddItem(filterbox, 0, 0, 1, 3, 0, 0, true).
		AddItem(errormsg, 1, 0, 1, 3, 0, 0, false).
		AddItem(table, 2, 0, 1, 3, 0, 0, false)

	grid.SetBackgroundColor(tcell.ColorDefault)

	return &Layout{Grid: grid, errormsg: errormsg, filterbox: filterbox, table: table}
}

func (l *Layout) AddLogEntry(e decoding.Entry) {
	l.table.AddLogEntry(e)
}
