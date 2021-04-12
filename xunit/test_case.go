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
	return &TestCase{
		name:     name,
		testable: testable,
		t:        t,
	}
}

type TestCase struct {
	name     string
	testable interface{}
	t        *testing.T
}

func (tc *TestCase) Run() {
	if setUp, ok := tc.testable.(interface{ SetUp() }); ok {
		setUp.SetUp()
	}

	method := reflect.ValueOf(tc.testable).MethodByName(tc.name)
	if (method == reflect.Value{}) {
		panic("method not found")
	}
	method.Call([]reflect.Value{})
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
