package kvswiss2

import (
	"fmt"
	"testing"
)

func TestSetGet_Resize_KeyString(t *testing.T) {
	cache := New[string, string](2048, 32, true)
	cache.Set("unicorns", "are cool")

	if res := cache.Get("unicorns"); res != "are cool" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", res, "are cool")
	}
}

func TestSetGet_Resize_KeyInt(t *testing.T) {
	cache := New[int, string](2048, 32, true)
	cache.Set(1337, "leet haxiors")

	if res := cache.Get(1337); res != "leet haxiors" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", res, "leet haxiors")
	}
}

func TestFlush_Resize(t *testing.T) {
	cache := New[int, string](2048, 32, true)
	cache.SetOnEvicted(func(k int, v string){
		fmt.Println(k)
	})

	cache.Set(1337, "leet haxiors")
	cache.Set(1338, "leet haxiors1")
	cache.Set(3434, "leet haxiors2")
	cache.Set(5465, "leet haxiors3")

	cache.Flush()
}

//Tests for cache without resize
func TestSetGet_KeyString(t *testing.T) {
	cache := New[string, string](2048, 32, false)
	cache.Set("unicorns", "are cool")

	if res := cache.Get("unicorns"); res != "are cool" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", res, "are cool")
	}
}

func TestSetGet_KeyInt(t *testing.T) {
	cache := New[int, string](2048, 32, false)
	cache.Set(1337, "leet haxiors")

	if res := cache.Get(1337); res != "leet haxiors" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", res, "leet haxiors")
	}
}

func TestFlush(t *testing.T) {
	cache := New[int, string](2048, 32, false)
	cache.SetOnEvicted(func(k int, v string){
		fmt.Println(k)
	})

	cache.Set(1337, "leet haxiors")
	cache.Set(1338, "leet haxiors1")
	cache.Set(3434, "leet haxiors2")
	cache.Set(5465, "leet haxiors3")

	cache.Flush()
}