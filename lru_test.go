package lru

import (
	"fmt"
	"testing"
)

func Example() {
	cache := New[int, struct{}](8, func(key int, _ struct{}) {
		fmt.Println("evicted", key)
	})

	for i := 1; i <= 10; i++ {
		cache.Replace(i, struct{}{})
	}

	fmt.Printf("%d items in cache\n", cache.Len())

	for k, v := range cache.IterateAsc() {
		fmt.Println(k, v)
	}

	// Output:
	//
	// evicted 1
	// evicted 2
	// 8 items in cache
	// 3 {}
	// 4 {}
	// 5 {}
	// 6 {}
	// 7 {}
	// 8 {}
	// 9 {}
	// 10 {}
}

func Benchmark(b *testing.B) {
	cache := New[int, struct{}](8)
	b.ResetTimer()

	for i := 8; i <= 512; i *= 2 {
		b.Run(fmt.Sprintf("cap_%03d", i), func(b *testing.B) {
			cache.Resize(i)
			b.ResetTimer()

			for i := range b.N {
				cache.Set(i, struct{}{})
			}
		})
	}
}

func BenchmarkThreadsafe(b *testing.B) {
	cache := NewThreadSafe[int, struct{}](8)
	b.ResetTimer()

	for i := 8; i <= 512; i *= 2 {
		b.Run(fmt.Sprintf("cap_%03d", i), func(b *testing.B) {
			cache.Resize(i)
			b.ResetTimer()

			for i := range b.N {
				cache.Set(i, struct{}{})
			}
		})
	}
}
