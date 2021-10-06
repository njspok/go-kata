package xunit

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func RunTestCase(t *testing.T, testable interface{}, name string) {
	NewTestCase(name, testable, t).Run()
}

func NewTestCase(name string, testable interface{}, t *testing.T) *TestCase {
	ts, ok := testable.(interface{ SetTestCase(*TestCase) })
	if !ok {
		panic("cant cast to SetTestCase")
	}

	tc := &TestCase{
		name:     name,
		testable: ts,
		t:        t,
	}

	ts.SetTestCase(tc)

	return tc
}

type TestCase struct {
	name     string
	testable interface{}
	t        *testing.T
}

func (tc *TestCase) Run() *TestResult {
	tc.t.Run(tc.name, func(t *testing.T) {

	})

	defer tc.runTearDown()

	result := NewTestResult()
	result.Start()

	tc.runSetUp()
	tc.run()

	return result
}

func (tc *TestCase) False(run bool) {
	require.False(tc.t, run)
}

func (tc *TestCase) True(run bool) {
	require.True(tc.t, run)
}

func (tc *TestCase) Equals(expected, actual interface{}) {
	require.Equal(tc.t, expected, actual)
}

func (tc *TestCase) Zero(value interface{}) {
	require.Zero(tc.t, value)
}

func (tc *TestCase) Panic(value interface{}, f func()) {
	require.PanicsWithValue(tc.t, value, f)
}

func (tc *TestCase) runSetUp() {
	if up, ok := tc.testable.(interface{ SetUp() }); ok {
		up.SetUp()
	}
}

func (tc *TestCase) run() {
	method := reflect.ValueOf(tc.testable).MethodByName(tc.name)
	if (method == reflect.Value{}) {
		panic("method not found")
	}

	method.Call([]reflect.Value{})
}

func (tc *TestCase) runTearDown() {
	if down, ok := tc.testable.(interface{ TearDown() }); ok {
		down.TearDown()
	}
}
