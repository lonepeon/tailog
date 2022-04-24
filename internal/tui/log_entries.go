package tui

import (
	"github.com/lonepeon/tailog/internal/decoding"
	"github.com/rivo/tview"
)

type LogEntries struct {
	tview.TableContentReadOnly
	buffer     *decoding.CircularBuffer
	fieldNames []string
}

func NewLogEntries(size int, fieldNames []string) *LogEntries {
	return &LogEntries{
		buffer:     decoding.NewCircularBuffer(size),
		fieldNames: fieldNames,
	}
}

func (l *LogEntries) GetCell(row, column int) *tview.TableCell {
	if column >= len(l.fieldNames) {
		return tview.NewTableCell("")
	}

	fieldName := l.fieldNames[column]

	if row == 0 {
		return tview.NewTableCell(fieldName)
	}

	maxIndex := l.buffer.Len() - 1
	entry, ok := l.buffer.At(maxIndex - (row - 1))
	if !ok {
		return tview.NewTableCell("")
	}

	field, ok := entry.Field(fieldName)
	if !ok {
		return tview.NewTableCell("")
	}

	return tview.NewTableCell(field.Value())
}

func (l *LogEntries) GetRowCount() int {
	return l.buffer.Len() + 1
}

func (l *LogEntries) GetColumnCount() int {
	return len(l.fieldNames)
}

func (l *LogEntries) AddLogEntry(entry decoding.Entry) {
	l.buffer.Push(entry)
}
