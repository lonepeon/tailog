package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/lonepeon/tailog/internal/decoding"
	"github.com/rivo/tview"
)

type Table struct {
	*tview.Table
	content *LogEntries
}

func NewTable(buffer int, fieldNames []string) *Table {
	tviewContent := NewLogEntries(buffer, fieldNames)
	tviewTable := tview.NewTable().
		SetBorders(true).
		SetFixed(1, len(fieldNames)).
		SetContent(tviewContent)
	tviewTable.SetBackgroundColor(tcell.ColorDefault)
	tviewTable.SetBordersColor(tcell.ColorDefault)

	table := Table{Table: tviewTable, content: tviewContent}
	table.setupInputCapture()

	return &table
}

func (t *Table) AddLogEntry(e decoding.Entry) {
	t.content.AddLogEntry(e)
}

func (t *Table) scrollUp() {
	row, col := t.GetOffset()
	t.SetOffset(row-1, col)
}

func (t *Table) scrollDown() {
	row, col := t.GetOffset()
	t.SetOffset(row+1, col)
}

func (t *Table) scrollLeft() {
	row, col := t.GetOffset()
	t.SetOffset(row, col-1)
}

func (t *Table) scrollRight() {
	row, col := t.GetOffset()
	t.SetOffset(row, col+1)
}

func (t *Table) setupInputCapture() {
	t.SetInputCapture(func(key *tcell.EventKey) *tcell.EventKey {
		switch key.Key() {
		case tcell.KeyDown:
			t.scrollDown()
		case tcell.KeyUp:
			t.scrollUp()
		case tcell.KeyLeft:
			t.scrollLeft()
		case tcell.KeyRight:
			t.scrollRight()
		default:
		}

		return key
	})
}
