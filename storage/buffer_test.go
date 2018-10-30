package storage

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BufferTestSuite struct {
	suite.Suite
}

func (s *BufferTestSuite) TestPush() {
	buff := NewBuffer()
	buff.Push(1)

	buff.Do(func(val interface{}) {
		if val == nil {
			return
		}

		s.Equal(1, val)
	})
}

func (s *BufferTestSuite) TestCopyLastValue() {
	buff := NewBuffer()
	buff.Push(1)
	buff.CopyLastValue()

	runs := 0
	buff.Do(func(val interface{}) {
		if val == nil {
			return
		}

		s.Equal(1, val)
		runs++
	})

	s.Equal(2, runs)
}

func (s *BufferTestSuite) TestToSlice() {
	buff := NewBuffer()
	buff.Push(1)

	s.Equal([]interface{}{1}, buff)
}

func TestBuffer(t *testing.T) {
	suite.Run(t, &BufferTestSuite{})
}
