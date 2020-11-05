package hashmap

import (
	"testing"
)

type HashFuncBenchKey struct {
	key Key
}
// Helpers
func benchmarkIntHashFunc(b *testing.B, blockSize int) {
	for i := 0; i < b.N; i++ {
		hashFunc(blockSize, i)
	}
}
func benchmarkStringHashFunc(b *testing.B, blockSize int) {
	for i := 0; i < b.N; i++ {
		hashFunc(blockSize, string(i))
	}
}
func benchmarkSliceHashFunc(b *testing.B, blockSize int) {
	for i := 0; i < b.N; i++ {
		hashFunc(blockSize, []int{i})
	}
}
func benchmarkMapHashFunc(b *testing.B, blockSize int) {
	for i := 0; i < b.N; i++ {
		hashFunc(blockSize, map[string]int{string(i):i})
	}
}
func benchmarkStructHashFunc(b *testing.B, blockSize int) {
	for i := 0; i < b.N; i++ {
		hashFunc(blockSize, HashFuncBenchKey{i})
	}
}
// Int
func BenchmarkIntHashFunc16(b *testing.B) {
	benchmarkIntHashFunc(b, 16)
}
func BenchmarkIntHashFunc64(b *testing.B) {
	benchmarkIntHashFunc(b, 64)
}
func BenchmarkIntHashFunc128(b *testing.B) {
	benchmarkIntHashFunc(b, 128)
}
func BenchmarkIntHashFunc1024(b *testing.B) {
	benchmarkIntHashFunc(b, 1024)
}
// String
func BenchmarkStringHashFunc16(b *testing.B) {
	benchmarkStringHashFunc(b, 16)
}
func BenchmarkStringHashFunc64(b *testing.B) {
	benchmarkStringHashFunc(b, 64)
}
func BenchmarkStringHashFunc128(b *testing.B) {
	benchmarkStringHashFunc(b, 128)
}
func BenchmarkStringHashFunc1024(b *testing.B) {
	benchmarkStringHashFunc(b, 1024)
}
// Slice
func BenchmarkSliceHashFunc16(b *testing.B) {
	benchmarkSliceHashFunc(b, 16)
}
func BenchmarkSliceHashFunc64(b *testing.B) {
	benchmarkSliceHashFunc(b, 64)
}
func BenchmarkSliceHashFunc128(b *testing.B) {
	benchmarkSliceHashFunc(b, 128)
}
func BenchmarkSliceHashFunc1024(b *testing.B) {
	benchmarkSliceHashFunc(b, 1024)
}
// Map
func BenchmarkMapHashFunc16(b *testing.B) {
	benchmarkMapHashFunc(b, 16)
}
func BenchmarkMapHashFunc64(b *testing.B) {
	benchmarkMapHashFunc(b, 64)
}
func BenchmarkMapHashFunc128(b *testing.B) {
	benchmarkMapHashFunc(b, 128)
}
func BenchmarkMapHashFunc1024(b *testing.B) {
	benchmarkMapHashFunc(b, 1024)
}
// Struct
func BenchmarkStructHashFunc16(b *testing.B) {
	benchmarkStructHashFunc(b, 16)
}
func BenchmarkStructHashFunc64(b *testing.B) {
	benchmarkStructHashFunc(b, 64)
}
func BenchmarkStructHashFunc128(b *testing.B) {
	benchmarkStructHashFunc(b, 128)
}
func BenchmarkStructHashFunc1024(b *testing.B) {
	benchmarkStructHashFunc(b, 1024)
}