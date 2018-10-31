package storage

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BufferTestSuite struct {
	suite.Suite
}

func (s *BufferTestSuite) TestPush() {
	buff := NewBuffer(3)
	buff.Push(1)

	buff.internal.Do(func(val interface{}) {
		if val == nil {
			return
		}

		s.Equal(1, val)
	})
}

func (s *BufferTestSuite) TestCopyLastValue() {
	buff := NewBuffer(3)
	buff.Push(1)
	buff.CopyLastValue()

	runs := 0
	buff.internal.Do(func(val interface{}) {
		if val == nil {
			return
		}

		s.Equal(1, val)
		runs++
	})

	s.Equal(2, runs)
}

func (s *BufferTestSuite) TestToSlice() {
	buff := NewBuffer(3)
	buff.Push(1)

	s.Equal([]interface{}{1}, buff.ToSlice())
}

func (s *BufferTestSuite) TestLen() {
	buff := NewBuffer(3)
	s.Equal(3, buff.Len())

	buff = NewBuffer(10)
	s.Equal(10, buff.Len())
}

func TestBuffer(t *testing.T) {
	suite.Run(t, new(BufferTestSuite))
}
