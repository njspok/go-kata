package xunit

import "testing"

func NewWasRun(name string, t *testing.T) *WasRun {
	wr := &WasRun{}
	wr.TestCase = NewTestCase(name, wr, t)
	return wr
}

type WasRun struct {
	*TestCase
	log []string
}

func (wr *WasRun) PanicMethod() {
	panic("PanicMethod")
}

func (wr *WasRun) TestMethod() {
	wr.log = append(wr.log, "TestMethod")
}

func (wr *WasRun) SetUp() {
	wr.log = append(wr.log, "SetUp")
}

func (wr *WasRun) TearDown() {
	wr.log = append(wr.log, "TearDown")
}
