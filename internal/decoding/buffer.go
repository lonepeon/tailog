package decoding

type CircularBuffer struct {
	entries    []Entry
	startIndex int
	nextIndex  int
	len        int
}

func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		entries:    make([]Entry, size),
		len:        0,
		startIndex: 0,
		nextIndex:  0,
	}
}

func (b *CircularBuffer) Len() int {
	return b.len
}

func (b *CircularBuffer) Entries() []Entry {
	entries := make([]Entry, 0, b.len)
	for i := 0; i < b.len; i++ {
		index := (b.startIndex + i) % cap(b.entries)
		entries = append(entries, b.entries[index])
	}

	return entries
}

// Push adds a new entry to the buffer. It returns a boolean indicating if
// a valud had to be evicted
func (b *CircularBuffer) Push(e Entry) bool {
	var hasEvictedValue bool
	b.entries[b.nextIndex] = e

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
