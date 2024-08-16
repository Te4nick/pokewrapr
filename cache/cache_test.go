package cache

import (
	//	"fmt"
	"fmt"
	"testing"
	"time"
)

func TestGetOk(t *testing.T) {
	cases := []struct {
		key   string
		value string
	}{
		{
			key:   "https://example.com",
			value: "testdata",
		},
		{
			key:   "https://example.com/path",
			value: "moretestdata",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache()

			cache.mutex.Lock()
			cache.content[c.key] = c.value
			cache.mutex.Unlock()

			val := cache.Get(c.key)
			if val == nil {
				t.Errorf("expected to find key")
				return
			}
			if stringVal, ok := val.(string); !ok || stringVal != string(c.value) {
				t.Errorf("expected to find string value")
				return
			}
		})
	}
}

func TestGetNil(t *testing.T) {

	cache := NewCache()

	val := cache.Get("nonexist")
	if val != nil {
		t.Errorf("expected not to find key")
		return
	}

}

func TestSetDefault(t *testing.T) {
	cases := []struct {
		key   string
		value string
	}{
		{
			key:   "https://example.com",
			value: "testdata",
		},
		{
			key:   "https://example.com/path",
			value: "moretestdata",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache()
			cache.SetDefault(c.key, c.value)

			cache.mutex.RLock()
			val, ok := cache.content[c.key]
			cache.mutex.RUnlock()

			if val == nil || !ok {
				t.Errorf("expected to find key")
				return
			}
			if stringVal, ok := val.(string); !ok || stringVal != string(c.value) {
				t.Errorf("expected to find string value")
				return
			}
		})
	}
}

func TestSet(t *testing.T) {
	cases := []struct {
		key   string
		value string
	}{
		{
			key:   "https://example.com",
			value: "testdata",
		},
		{
			key:   "https://example.com/path",
			value: "moretestdata",
		},
	}

	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime * 4
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache()
			cache.Set(c.key, c.value, baseTime)

			cache.mutex.RLock()
			val1, ok := cache.content[c.key]
			cache.mutex.RUnlock()
			if val1 == nil || !ok {
				t.Errorf("expected to find key")
				return
			}

			if stringVal, ok := val1.(string); !ok || stringVal != string(c.value) {
				t.Errorf("expected to find string value")
				return
			}

			time.Sleep(waitTime)

			cache.mutex.RLock()
			_, ok = cache.content[c.key]
			cache.mutex.RUnlock()
			if ok {
				t.Errorf("expected to not find key")
				return
			}
		})
	}
}
