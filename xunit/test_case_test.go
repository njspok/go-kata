package xunit

import (
	"testing"
)

func NewTestCaseTest(name string, t *testing.T) *TestCaseTest {
	tct := &TestCaseTest{}
	tct.TestCase = NewTestCase(name, tct, t)
	return tct
}

type TestCaseTest struct {
	*TestCase
}

func (t *TestCaseTest) TestTemplateMethod() {
	wr := NewWasRun("TestMethod", nil)
	wr.Run()
	t.Equals([]string{
		"SetUp",
		"TestMethod",
		"TearDown",
	}, wr.log)
}

func (t *TestCaseTest) TestPanicMethod() {
	wr := NewWasRun("PanicMethod", nil)

	t.Panic("PanicMethod", func() {
		wr.Run()
	})

	t.Equals([]string{
		"SetUp",
		"TearDown",
	}, wr.log)
}

func TestTestCaseTest(t *testing.T) {
	NewTestCaseTest("TestTemplateMethod", t).Run()
	NewTestCaseTest("TestPanicMethod", t).Run()

}
