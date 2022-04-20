package decoding

import (
	"errors"
	"io"
	"time"
)

type WatchFunc func(Entry, error)

type Watcher struct {
	decoder    Decoder
	notifiers  []WatchFunc
	stopSignal chan interface{}
}

func NewWatcher(decoder Decoder) *Watcher {
	return &Watcher{
		decoder:    decoder,
		stopSignal: make(chan interface{}),
	}
}

func (w *Watcher) Start() {
	go func() {
		for {
			select {
			case <-w.stopSignal:
				return
			default:
			}

			if !w.decoder.More() {
				time.Sleep(1 * time.Second)
				continue
			}

			entry, err := w.decoder.Decode()
			if errors.Is(err, io.EOF) {
				return
			}

			for _, fn := range w.notifiers {
				fn(entry, err)
			}
		}
	}()
}

func (w *Watcher) Stop() {
	w.stopSignal <- nil
}

func (w *Watcher) Notify(fn WatchFunc) {
	w.notifiers = append(w.notifiers, fn)
}
