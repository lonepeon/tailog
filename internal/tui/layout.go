package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/lonepeon/tailog/internal/decoding"
	"github.com/rivo/tview"
)

type Layout struct {
	*tview.Grid

	filterbox *tview.InputField
	table     *Table
}

func NewLayout(buffer int, fieldNames []string) *Layout {
	filterbox := NewFilterbox()
	table := NewTable(buffer, fieldNames)

	grid := tview.NewGrid().
		SetRows(1, 0).
		SetBorders(true).
		AddItem(filterbox, 0, 0, 1, 3, 0, 0, true).
		AddItem(table, 1, 0, 1, 3, 0, 0, false)

	grid.SetBackgroundColor(tcell.ColorDefault)

	return &Layout{Grid: grid, filterbox: filterbox, table: table}
}

func (l *Layout) AddLogEntry(e decoding.Entry) {
	l.table.AddLogEntry(e)
}
