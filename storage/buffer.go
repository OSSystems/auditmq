package storage

import "container/ring"

type Buffer struct {
	internal *ring.Ring
}

func NewBuffer(samples int) *Buffer {
	return &Buffer{
		internal: ring.New(samples),
	}
}

func (b *Buffer) Len() int {
	return b.internal.Len()
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
	b.internal.Do(func(val interface{}) {
		if val == nil {
			return
		}

		sample = append(sample, val)
	})

	return sample
}
