package token_ring

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	token := make(chan struct{}, 1)
	token <- struct{}{}

	counter := 0

	go func() {
		for range 10000 {
			<-token
			counter++
			token <- struct{}{}
		}
	}()

	go func() {
		for range 10000 {
			<-token
			counter++
			token <- struct{}{}
		}
	}()

	time.Sleep(2 * time.Second)

	require.Equal(t, 20000, counter)
}

//func Test(t *testing.T) {
//	counter := 0
//
//	p1 := NewProcess(func() { counter++ }, 1000)
//	p2 := NewProcess(func() { counter++ }, 1000)
//
//	ring := NewRing(p1, p2)
//
//	go p1.Run()
//	go p2.Run()
//
//	time.Sleep(1 * time.Second)
//
//	require.Equal(t, 2000, counter)
//}
