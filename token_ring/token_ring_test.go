package token_ring

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	ring := make(chan struct{}, 1)
	token := struct{}{}
	ring <- token

	counter := 0

	wg := sync.WaitGroup{}
	wg.Go(func() {
		for range 10000 {
			<-ring
			counter++
			ring <- token
		}
	})
	wg.Go(func() {
		for range 10000 {
			<-ring
			counter++
			ring <- token
		}
	})
	wg.Wait()

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
