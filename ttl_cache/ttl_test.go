package ttl_cache

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTTLCache(t *testing.T) {
	t.Run("initialize", func(t *testing.T) {
		cache := NewTTLCache()
		require.Zero(t, cache.Size())
		require.Equal(t, Stats{Hits: 0, Miss: 0}, cache.Stats())
	})
	t.Run("put element", func(t *testing.T) {
		// Arrange
		cache := NewTTLCache()

		// Act
		cache.Set("one", 1, 0)

		// Assert
		require.EqualValues(t, 1, cache.Size())

		_, exist := cache.Get("ten")
		require.False(t, exist)
		require.Equal(t, Stats{Hits: 0, Miss: 1}, cache.Stats())

		v, exist := cache.Get("one")
		require.True(t, exist)
		require.EqualValues(t, 1, v)
		require.Equal(t, Stats{Hits: 1, Miss: 1}, cache.Stats())
	})
	t.Run("put another element", func(t *testing.T) {
		// Arrange
		cache := NewTTLCache()
		cache.Set("one", 1, 0)

		// Act
		cache.Set("two", 2, 0)

		// Assert
		require.EqualValues(t, 2, cache.Size())

		_, exist := cache.Get("ten")
		require.False(t, exist)
		require.Equal(t, Stats{Hits: 0, Miss: 1}, cache.Stats())

		v, exist := cache.Get("one")
		require.True(t, exist)
		require.EqualValues(t, 1, v)
		require.Equal(t, Stats{Hits: 1, Miss: 1}, cache.Stats())

		v, exist = cache.Get("two")
		require.True(t, exist)
		require.EqualValues(t, 2, v)
		require.Equal(t, Stats{Hits: 2, Miss: 1}, cache.Stats())
	})
	t.Run("delete not exist element", func(t *testing.T) {
		// Arrange
		cache := NewTTLCache()

		// Act
		res := cache.Delete("some")

		// Assert
		require.False(t, res)
		require.Equal(t, Stats{Hits: 0, Miss: 0}, cache.Stats())
		require.Zero(t, cache.Size())
	})
	t.Run("delete exist element", func(t *testing.T) {
		// Arrange
		cache := NewTTLCache()
		cache.Set("one", 1, 0)
		cache.Set("two", 2, 0)

		// Act
		res := cache.Delete("one")

		// Assert
		require.True(t, res)

		_, exist := cache.Get("one")
		require.False(t, exist)

		_, exist = cache.Get("two")
		require.True(t, exist)

		require.Equal(t, Stats{Hits: 1, Miss: 1}, cache.Stats())
		require.EqualValues(t, 1, cache.Size())
	})

}

func TestTTLCache_Get(t *testing.T) {
	cache := NewTTLCache()
	cache.Set("world", 1, 0)

	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		// получаем не существующее значение
		wg.Add(1)
		go func() {
			_, _ = cache.Get("hello")
			wg.Done()
		}()

		// получаем существующее значение
		wg.Add(1)
		go func() {
			_, _ = cache.Get("world")
			wg.Done()
		}()
	}
	wg.Wait()

	require.EqualValues(t, 1000, cache.Stats().Miss)
	require.EqualValues(t, 1000, cache.Stats().Hits)
}

func TestTTLCache_DataRace(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		cache := NewTTLCache()

		wg := sync.WaitGroup{}
		for i := 0; i < 1000; i++ {
			// получаем знаение из кшеа
			wg.Add(1)
			go func() {
				_, _ = cache.Get("hello")
				wg.Done()
			}()

			// получаем статистику
			wg.Add(1)
			go func() {
				cache.Stats()
				wg.Done()
			}()

			// сбрасываем статистику
			wg.Add(1)
			go func() {
				cache.ResetStats()
				wg.Done()
			}()
		}
		wg.Wait()
	})

	t.Run("two", func(t *testing.T) {
		cache := NewTTLCache()

		wg := sync.WaitGroup{}
		for i := 0; i < 10000; i++ {
			// читаем значение из кеша
			wg.Add(1)
			go func() {
				_, _ = cache.Get("hello")
				wg.Done()
			}()

			// кладем значение в кеш
			wg.Add(1)
			go func() {
				cache.Set("hello", "world", 0)
				wg.Done()
			}()
		}
		wg.Wait()
	})
}
