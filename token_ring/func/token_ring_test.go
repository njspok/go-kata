package _func

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	counter := 0

	start := make(chan struct{}, 1)
	start <- struct{}{}

	p1 := NewProc(func() { counter++ }, start)
	p2 := NewProc(func() { counter++ }, p1)
	p3 := NewProc(func() { counter++ }, p2)

	// замыкаем круг
	// делаем N round
	for range 1000 {
		<-p3
		start <- struct{}{}
	}

	close(start)

	require.Equal(t, 3000, counter)
}
