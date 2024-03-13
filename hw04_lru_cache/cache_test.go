package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(4)

		c.Set("0", 100)
		c.Set("1", 101)
		c.Set("2", 102)
		c.Set("3", 103)

		c.Clear()

		for i := 1; i <= 4; i++ {
			value, ok := c.Get(Key(strconv.Itoa(i)))
			require.Nil(t, value)
			require.False(t, ok)
		}
	})

	t.Run("push out", func(t *testing.T) {
		c := NewCache(4)

		c.Set("0", 100)
		c.Set("1", 101)
		c.Set("2", 102)
		c.Set("3", 103)
		c.Set("4", 104)

		value, ok := c.Get("0")
		require.Nil(t, value)
		require.False(t, ok)
	})

	t.Run("push out oldest", func(t *testing.T) {
		c := NewCache(4)

		c.Set("0", 100)
		c.Set("1", 101)
		c.Set("2", 102)
		c.Set("3", 103)

		c.Get("2")
		c.Get("3")
		c.Get("0")
		c.Get("1")

		c.Set("4", 104)

		value, ok := c.Get("2")
		require.Nil(t, value)
		require.False(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
