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

func (t *TestCaseTest) TestRunning() {
	test := NewWasRun("TestMethod", nil)
	t.False(test.wasRun)
	test.Run()
	t.True(test.wasRun)
}

func TestTestCaseTest(t *testing.T) {
	NewTestCaseTest("TestRunning", t).Run()
}
