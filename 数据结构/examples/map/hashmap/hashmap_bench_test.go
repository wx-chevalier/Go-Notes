package hashmap

import (
	"testing"
)

type BenchKey struct {
	key Key
}
type BenchValue struct {
	value interface{}
}
// Set helpers
func benchmarkSet(b *testing.B, blockSize int) {
	hashMap := NewHashMap(blockSize)
	for i := 0; i < b.N; i++ {
		hashMap.Set(i, i)
	}
}
func benchmarkStringSet(b *testing.B, blockSize int) {
	hashMap := NewHashMap(blockSize)
	for i := 0; i < b.N; i++ {
		hashMap.Set(string(i), string(i))
	}
}
func benchmarkSliceSet(b *testing.B, blockSize int) {
	hashMap := NewHashMap(blockSize)
	for i := 0; i < b.N; i++ {
		hashMap.Set([]int{i}, []int{i})
	}
}
func benchmarkMapSet(b *testing.B, blockSize int) {
	hashMap := NewHashMap(blockSize)
	for i := 0; i < b.N; i++ {
		hashMap.Set(map[int]int{i:i}, map[int]int{i:i})
	}
}
func benchmarkStructSet(b *testing.B, blockSize int) {
	hashMap := NewHashMap(blockSize)
	for i := 0; i < b.N; i++ {
		hashMap.Set(BenchKey{i}, BenchValue{i})
	}
}
// Get helpers
func benchmarkGet(b *testing.B, blockSize int) {
	b.StopTimer()

	hashMap := NewHashMap(blockSize)
	for i := 0; i < b.N; i++ {
		hashMap.Set(i, i)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err := hashMap.Get(i)
		if err != nil {
			b.Errorf("Inserted key %d not found", i)
		}
	}
}
func benchmarkStringGet(b *testing.B, blockSize int) {
	b.StopTimer()

	hashMap := NewHashMap(blockSize)
	for i := 0; i < b.N; i++ {
		hashMap.Set(string(i), string(i))
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err := hashMap.Get(string(i))
		if err != nil {
			b.Errorf("Inserted key %d not found", i)
		}
	}
}
func benchmarkSliceGet(b *testing.B, blockSize int) {
	b.StopTimer()

	hashMap := NewHashMap(blockSize)
	for i := 0; i < b.N; i++ {
		hashMap.Set([]int{i}, []int{i})
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err := hashMap.Get([]int{i})
		if err != nil {
			b.Errorf("Inserted key []int{%d} not found", i)
		}
	}
}
func benchmarkMapGet(b *testing.B, blockSize int) {
	b.StopTimer()

	hashMap := NewHashMap(blockSize)
	for i := 0; i < b.N; i++ {
		hashMap.Set(map[int]int{i:i}, map[int]int{i:i})
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err := hashMap.Get(map[int]int{i:i})
		if err != nil {
			b.Errorf("Inserted key map[int]int{%d:%d} not found", i, i)
		}
	}
}
func benchmarkStructGet(b *testing.B, blockSize int) {
	b.StopTimer()

	hashMap := NewHashMap(blockSize)
	for i := 0; i < b.N; i++ {
		hashMap.Set(BenchKey{i}, BenchValue{i})
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, err := hashMap.Get(BenchKey{i})
		if err != nil {
			b.Errorf("Inserted key BenchKey{%d} not found", i)
		}
	}
}
// Unset helpers
func benchmarkUnset(b *testing.B, blockSize int) {
	b.StopTimer()

	hashMap := NewHashMap(blockSize)
	for i := 0; i < b.N; i++ {
		hashMap.Set(i, i)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		err := hashMap.Unset(i)
		if err != nil {
			b.Errorf("Unset key %d error", i)
		}
	}
}
func setupNativeMap(b *testing.B) map[string]int {
	mapX := make(map[string]int)
	for i := 0; i < b.N; i++ {
		mapX[string(i)] = i
	}

	return mapX
}
//
// ================ Tests =================
//
// Set
func BenchmarkSet16(b *testing.B) {
	benchmarkSet(b, 16)
}

func BenchmarkSet64(b *testing.B) {
	benchmarkSet(b, 64)
}

func BenchmarkSet128(b *testing.B) {
	benchmarkSet(b, 128)
}

func BenchmarkSet1024(b *testing.B) {
	benchmarkSet(b, 1024)
}

func BenchmarkStringSet1024(b *testing.B) {
	benchmarkStringSet(b, 1024)
}

func BenchmarkSliceSet1024(b *testing.B) {
	benchmarkSliceSet(b, 1024)
}

func BenchmarkMapSet16(b *testing.B) {
	benchmarkMapSet(b, 16)
}

func BenchmarkStructSet16(b *testing.B) {
	benchmarkStructSet(b, 16)
}

// Get
func BenchmarkGet16(b *testing.B) {
	benchmarkGet(b, 16)
}

func BenchmarkGet64(b *testing.B) {
	benchmarkGet(b, 64)
}

func BenchmarkGet128(b *testing.B) {
	benchmarkGet(b, 128)
}

func BenchmarkGet1024(b *testing.B) {
	benchmarkGet(b, 1024)
}

func BenchmarkStringGet16(b *testing.B) {
	benchmarkStringGet(b, 16)
}

func BenchmarkSliceGet16(b *testing.B) {
	benchmarkSliceGet(b, 16)
}

func BenchmarkStructGet16(b *testing.B) {
	benchmarkStructGet(b, 16)
}

func BenchmarkMapGet16(b *testing.B) {
	benchmarkMapGet(b, 16)
}

// Unset
func BenchmarkUnset16(b *testing.B) {
	benchmarkUnset(b, 16)
}

func BenchmarkUnset64(b *testing.B) {
	benchmarkUnset(b, 64)
}

func BenchmarkUnset128(b *testing.B) {
	benchmarkUnset(b, 128)
}

func BenchmarkUnset1024(b *testing.B) {
	benchmarkUnset(b, 1024)
}

// Native map
func BenchmarkIntNativeMapSet(b *testing.B) {
	mapX := make(map[int]int)
	for i := 0; i < b.N; i++ {
		mapX[i] = i
	}
}
func BenchmarkStringNativeMapSet(b *testing.B) {
	mapX := make(map[string]string)
	for i := 0; i < b.N; i++ {
		mapX[string(i)] = string(i)
	}
}

var result int
func BenchmarkStringNativeMapGet(b *testing.B) {
	mapX := setupNativeMap(b)
	for i := 0; i < b.N; i++ {
		result = mapX[string(i)];
	}
}

func BenchmarkStringNativeMapDelete(b *testing.B) {
	mapX := setupNativeMap(b)
	for i := 0; i < b.N; i++ {
		delete(mapX, string(i))
	}
}