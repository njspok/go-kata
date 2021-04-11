package xunit

import "testing"

func NewWasRun(name string, t *testing.T) *WasRun {
	wr := &WasRun{
		wasRun: false,
	}
	wr.TestCase = NewTestCase(name, wr, t)
	return wr
}

type WasRun struct {
	wasRun bool
	*TestCase
}

func (r *WasRun) TestMethod() {
	r.wasRun = true
}
