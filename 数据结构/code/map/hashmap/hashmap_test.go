package hashmap

import (
	"fmt"
	"testing"
)

const iterationCount = 1000000

// Set int: int
func testIntSet(t *testing.T, blockSize int) {
	hashMap := NewHashMap(blockSize, hashFunc)

	for i := 0; i < iterationCount; i++ {
		err := hashMap.Set(i, i)
		if err != nil {
			t.Errorf("Setting error for key %d", i)
		}
	}
}

func TestIntSet16(t *testing.T) {
	testIntSet(t, 16)
}

func TestIntSet64(t *testing.T) {
	testIntSet(t, 64)
}

func TestIntSet128(t *testing.T) {
	testIntSet(t, 128)
}

func TestIntSet1024(t *testing.T) {
	testIntSet(t, 1024)
}

// Get int
func testIntGet(t *testing.T, blockSize int) {
	hashMap := NewHashMap(blockSize)

	for i := 0; i < iterationCount; i++ {
		err := hashMap.Set(i, i)
		if err != nil {
			t.Errorf("Setting error for key %d", i)
		}
	}

	for i := 0; i < iterationCount; i++ {
		_, err := hashMap.Get(i)
		if err != nil {
			t.Errorf("Inserted key %d not found", i)
		}
	}
}

// Get int
func testStringGet(t *testing.T, blockSize int) {
	hashMap := NewHashMap(blockSize)

	for i := 0; i < iterationCount; i++ {
		err := hashMap.Set(string(i), string(i))
		if err != nil {
			t.Errorf("Setting error for key %d", i)
		}
	}

	for i := 0; i < iterationCount; i++ {
		_, err := hashMap.Get(string(i))
		if err != nil {
			t.Errorf("Inserted key %d not found", i)
		}
	}
}

func TestIntGet16(t *testing.T) {
	testIntGet(t, 16)
}

func TestIntGet64(t *testing.T) {
	testIntGet(t, 64)
}

func TestIntGet128(t *testing.T) {
	testIntGet(t, 128)
}

func TestIntGet1024(t *testing.T) {
	testIntGet(t, 1024)
}

func TestStringGet1024(t *testing.T) {
	testStringGet(t, 1024)
}

// Unset int
func testIntUnset(t *testing.T, blockSize int) {
	hashMap := NewHashMap(blockSize)

	for i := 0; i < iterationCount; i++ {
		err := hashMap.Set(i, i)
		if err != nil {
			t.Errorf("Setting error for key %d", i)
		}
	}
	for i := 0; i < iterationCount; i++ {
		err := hashMap.Unset(i)
		if err != nil {
			t.Errorf("Unset key %d error", i)
		}
	}
}

func TestIntUnset16(t *testing.T) {
	testIntUnset(t, 16)
}

func TestIntUnset64(t *testing.T) {
	testIntUnset(t, 64)
}

func TestIntUnset128(t *testing.T) {
	testIntUnset(t, 128)
}

func TestIntUnset1024(t *testing.T) {
	testIntUnset(t, 1024)
}

// Set string
func testStringSet(t *testing.T, blockSize int) {
	hashMap := NewHashMap(blockSize)

	for i := 0; i < iterationCount; i++ {
		err := hashMap.Set(string(i), string(i))
		if err != nil {
			t.Errorf("Setting error for key %d", i)
		}
	}
}

func TestStringSet1024(t *testing.T) {
	testStringSet(t, 1024)
}

// Iterate
func TestIterate(t *testing.T) {
	hashMap := NewHashMap(16)

	for i := 0; i < 100; i++ {
		err := hashMap.Set(i, i)
		if err != nil {
			t.Errorf("Setting error for key %d", i)
		}
	}

	for r := range hashMap.Iter() {
		fmt.Sprintf("%s: %s, ", r.key, r.value)
	}
}
