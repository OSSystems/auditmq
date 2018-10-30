package storage

import "container/ring"

const samples = 10

type Buffer struct {
	internal *ring.Ring
}

func NewBuffer() *Buffer {
	return &Buffer{
		internal: ring.New(samples),
	}
}

func (b *Buffer) Do(f func(val interface{})) {
	b.internal.Do(f)
}

func (b *Buffer) Push(data interface{}) {
	b.internal.Value = data
	b.internal = b.internal.Next()
}

func (b *Buffer) CopyLastValue() {
	b.internal.Value = b.internal.Prev().Value
	b.internal = b.internal.Next()
}

func (b *Buffer) ToSlice() []interface{} {
	sample := []interface{}{}
	b.Do(func(val interface{}) {
		if val == nil {
			return
		}

		sample = append(sample, val)
	})

	return sample
}
