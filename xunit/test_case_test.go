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

func (t *TestCaseTest) TestRunning() {
	t.obj.Run()
	t.True(t.obj.wasRun)
}

func (t *TestCaseTest) TestSetUp() {
	t.obj.Run()
	t.True(t.obj.wasSetUp)
}

func TestTestCaseTest(t *testing.T) {
	NewTestCaseTest("TestRunning", t).Run()
	NewTestCaseTest("TestSetUp", t).Run()
}
