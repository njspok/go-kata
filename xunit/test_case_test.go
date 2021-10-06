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

func (t *TestCaseTest) SetTestCase(tc *TestCase) {
	t.TestCase = tc
}

func (t *TestCaseTest) TestResultMethod() {
	wr := NewWasRun("TestMethod", t.t)
	result := wr.Run()
	t.Equals("1 run, 0 failed", result.Summary())
}

func (t *TestCaseTest) TestTemplateMethod() {
	wr := NewWasRun("TestMethod", t.t)
	wr.Run()
	t.Equals([]string{
		"SetUp",
		"TestMethod",
		"TearDown",
	}, wr.log)
}

func (t *TestCaseTest) TestPanicMethod() {
	wr := NewWasRun("PanicMethod", t.t)

	t.Panic("PanicMethod", func() {
		wr.Run()
	})

	t.Equals([]string{
		"SetUp",
		"TearDown",
	}, wr.log)
}

func TestTestCaseTest(t *testing.T) {
	//NewTestCaseTest("TestTemplateMethod", t).Run()
	//NewTestCaseTest("TestPanicMethod", t).Run()
	//NewTestCaseTest("TestResultMethod", t).Run()

	tct := &TestCaseTest{}
	RunTestCase(t, tct, "TestTemplateMethod")
	RunTestCase(t, tct, "TestPanicMethod")
	RunTestCase(t, tct, "TestResultMethod")
}
