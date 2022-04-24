package cmd

import (
	"io"

	"github.com/lonepeon/tailog/internal/decoding"
	"github.com/lonepeon/tailog/internal/decoding/structuredjson"
	"github.com/lonepeon/tailog/internal/tui"
	"github.com/rivo/tview"
)

func StartTUI(appName string, stdin io.Reader, buffer int, fieldNames []string) error {
	layout := tui.NewLayout(buffer, fieldNames)
	app := tview.NewApplication().SetRoot(layout, true).EnableMouse(true)

	watcher := decoding.NewWatcher(structuredjson.NewDecoder(stdin))
	watcher.Notify(func(entry decoding.Entry, err error) {
		if err != nil {
			panic(err.Error())
		}

		app.QueueUpdateDraw(func() { layout.AddLogEntry(entry) })
	})

	watcher.Start()
	defer watcher.Stop()

	return app.Run()
}
