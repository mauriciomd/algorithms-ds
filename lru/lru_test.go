package lru

import "testing"

func TestLRU(t *testing.T) {
	t.Run("Method Set when length < capacity", func(t *testing.T) {
		lru := NewCache(4)
		lru.Set("foo", 1)

		got := lru.head.value
		want := 1

		if got != want {
			t.Errorf("got %d want :%d", got, want)
		}

		lru.Set("bar", 2)
		got = lru.tail.value
		want = 2

		if got != want {
			t.Errorf("got %d want :%d", got, want)
		}

		lru.Set("foo", 3)
		got = lru.tail.value
		want = 3

		if got != want {
			t.Errorf("got %d want :%d", got, want)
		}
	})

	t.Run("Method Set - Evict block", func(t *testing.T) {
		lru := NewCache(4)
		lru.Set("foo", 1)
		lru.Set("bar", 2)
		lru.Set("foobar", 3)
		lru.Set("bazinga", 4)
		lru.Set("xpto", 4)

		want := lru.head.value
		got := 2

		if got != want {
			t.Errorf("got %d want :%d", got, want)
		}

		want = lru.tail.value
		got = 4
		if got != want {
			t.Errorf("got %d want :%d", got, want)
		}
		_, exists := lru.lookup["foo"]

		if exists {
			t.Errorf("got %t want false", exists)
		}
	})

	t.Run("Method Get", func(t *testing.T) {
		lru := NewCache(4)
		lru.Set("foo", 1)
		lru.Set("bar", 2)
		lru.Set("foobar", 3)
		lru.Set("bazinga", 4)

		lru.Get("foo")
		want := lru.tail.value
		got := 1

		if got != want {
			t.Errorf("got %d want :%d", got, want)
		}

		lru.Get("foobar")
		want = lru.tail.value
		got = 3

		if got != want {
			t.Errorf("got %d want :%d", got, want)
		}
	})
}
