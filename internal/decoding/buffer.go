package decoding

import "sync"

type FilterFunc func(Entry) bool

func KeepAll(e Entry) bool {
	return true
}

type FilteredEntry struct {
	Entry  Entry
	Hidden bool
}

type CircularBuffer struct {
	entries    []FilteredEntry
	startIndex int
	nextIndex  int
	len        int
	mux        sync.RWMutex
	filterFn   FilterFunc
}

func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		entries:    make([]FilteredEntry, size),
		len:        0,
		startIndex: 0,
		nextIndex:  0,
		filterFn:   KeepAll,
	}
}

func (b *CircularBuffer) At(askedIndex int) (Entry, bool) {
	if askedIndex >= b.Len() || askedIndex < 0 {
		return nil, false
	}

	b.mux.RLock()
	defer b.mux.RUnlock()

	var filteredIndex int
	for i := 0; i < b.len; i++ {
		entry := b.at(i)
		if entry.Hidden {
			continue
		}

		if filteredIndex == askedIndex {
			return entry.Entry, true
		}

		filteredIndex++
	}

	return nil, false
}

func (b *CircularBuffer) at(index int) FilteredEntry {
	return b.entries[(b.startIndex+index)%b.len]
}

func (b *CircularBuffer) set(index int, entry FilteredEntry) {
	b.entries[(b.startIndex+index)%b.len] = entry
}

func (b *CircularBuffer) Len() int {
	b.mux.RLock()
	defer b.mux.RUnlock()

	var length int
	for i := 0; i < b.len; i++ {
		entry := b.at(i)
		if entry.Hidden {
			continue
		}
		length++
	}

	return length
}

func (b *CircularBuffer) ReplaceFilter(fn FilterFunc) {
	b.filterFn = fn

	b.mux.Lock()
	defer b.mux.Unlock()

	for i := 0; i < b.len; i++ {
		entry := b.at(i)

		b.set(i, FilteredEntry{
			Entry:  entry.Entry,
			Hidden: !b.filterFn(entry.Entry),
		})
	}
}

// Push adds a new entry to the buffer. It returns a boolean indicating if
// a valud had to be evicted
func (b *CircularBuffer) Push(e Entry) bool {
	b.mux.Lock()
	defer b.mux.Unlock()

	var hasEvictedValue bool
	isShown := b.filterFn(e)
	b.entries[b.nextIndex] = FilteredEntry{Entry: e, Hidden: !isShown}

	length := b.len + 1
	bufferCapacity := cap(b.entries)
	if length > bufferCapacity {
		hasEvictedValue = true
		length = bufferCapacity
	}
	b.len = length

	b.nextIndex = (b.nextIndex + 1) % bufferCapacity
	if b.len == bufferCapacity {
		b.startIndex = b.nextIndex
	}

	return hasEvictedValue
}
