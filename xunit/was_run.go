package xunit

import "testing"

func NewWasRun(name string, t *testing.T) *WasRun {
	wr := &WasRun{
		wasRun:   false,
		wasSetUp: false,
	}
	wr.TestCase = NewTestCase(name, wr, t)
	return wr
}

type WasRun struct {
	*TestCase
	wasRun   bool
	wasSetUp bool
}

func (wr *WasRun) TestMethod() {
	wr.wasRun = true
}

func (wr *WasRun) SetUp() {
	wr.wasRun = false
	wr.wasSetUp = true
}
