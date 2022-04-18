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

func (b *CircularBuffer) Push(e Entry) {
	b.entries[b.nextIndex] = e

	length := b.len + 1
	bufferCapacity := cap(b.entries)
	if length > bufferCapacity {
		length = bufferCapacity
	}
	b.len = length

	b.nextIndex = (b.nextIndex + 1) % bufferCapacity
	if b.len == bufferCapacity {
		b.startIndex = b.nextIndex
	}
}
