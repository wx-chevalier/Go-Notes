package hashmap

import (
	"testing"
)
type HashKey struct {
	key interface{}
}

func testKeyHash(t *testing.T, availableHash uint, hashKey uint)  {
	if hashKey != availableHash {
		t.Errorf("Expected hash key %d, got %d", availableHash, hashKey)
	}
}

func testBucketIdx(t *testing.T, availableIdx uint, bucketIdx uint)  {
	if bucketIdx != availableIdx {
		t.Errorf("Expected bucket idx %d, got %d", availableIdx, bucketIdx)
	}
}

func TestIntHash16(t *testing.T) {
	hashKey, bucketIdx := hashFunc(16, 123112122)

	testKeyHash(t, uint(249811310353814468), hashKey)
	testBucketIdx(t, uint(4), bucketIdx)
}

func TestStringHash16(t *testing.T) {
	hashKey, bucketIdx := hashFunc(16, "test_value")

	testKeyHash(t, uint(11696809291008937669), hashKey)
	testBucketIdx(t, uint(5), bucketIdx)
}

func TestSliceHash16(t *testing.T) {
	hashKey, bucketIdx := hashFunc(16, []string{"test_value"})

	testKeyHash(t, uint(13905799525442369900), hashKey)
	testBucketIdx(t, uint(12), bucketIdx)
}

func TestMapHash16(t *testing.T) {
	hashKey, bucketIdx := hashFunc(16, map[string]string{"key":"value"})

	testKeyHash(t, uint(6432338246103250761), hashKey)
	testBucketIdx(t, uint(9), bucketIdx)
}

func TestStructHash16(t *testing.T) {
	hashKey, bucketIdx := hashFunc(16, HashKey{"key"})

	testKeyHash(t, uint(2074006382534894026), hashKey)
	testBucketIdx(t, uint(10), bucketIdx)
}