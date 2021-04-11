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
	method := reflect.ValueOf(tc.testable).MethodByName(tc.name)
	if (method == reflect.Value{}) {
		panic("method not found")
	}
	method.Call([]reflect.Value{})
}

func (t *TestCase) False(run bool) {
	require.False(t.t, run)
}

func (t *TestCase) True(run bool) {
	require.True(t.t, run)
}
