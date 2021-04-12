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
	obj *WasRun
}

func (t *TestCaseTest) SetUp() {
	t.obj = NewWasRun("TestMethod", nil)
}

func (t *TestCaseTest) TestTemplateMethod() {
	t.obj.Run()
	t.Equals([]string{"SetUp", "TestMethod"}, t.obj.log)
}

func TestTestCaseTest(t *testing.T) {
	NewTestCaseTest("TestTemplateMethod", t).Run()
}
