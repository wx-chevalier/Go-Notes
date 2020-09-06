// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The hashmap package
package hashmap

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	growLoadFactor   float32 = 0.75
	shrinkLoadFactor float32 = 0.25
)

// Key interface
type Key interface{}

// HashMaper interface
type HashMaper interface {
	Set(key Key, value interface{}) error
	Get(key Key) (value interface{}, err error)
	Unset(key Key) error
	Count() int
	Iter() <-chan KeyValue
}

// HashMap
type HashMap struct {
	hashFunc         func(blockSize int, key Key) (hashKey uint, bucketIdx uint) // hash func
	defaultBlockSize int                                                         // buckets block size
	buckets          []*Bucket                                                   // buckets for chains
	size             int                                                         // size of hash map
	shrinked         bool                                                        // is shrinked
	halfSlice        bool                                                        // half slice used in buckets
}

// Bucket
type Bucket struct {
	hashKey uint
	key     Key
	value   interface{}
	next    *Bucket
}

// KeyValue
type KeyValue struct {
	key   Key
	value interface{}
}

// New HashMap.
func NewHashMap(blockSize int, fn ...func(blockSize int, key Key) (hashKey uint, bucketIdx uint)) HashMaper {
	//	log.Printf("New\n")
	hashMap := new(HashMap)
	hashMap.defaultBlockSize = blockSize
	hashMap.buckets = make([]*Bucket, hashMap.defaultBlockSize)
	hashMap.size = 0
	hashMap.shrinked = false
	hashMap.halfSlice = true

	if len(fn) > 0 && fn[0] != nil && isFunc(fn[0]) {
		//fmt.Println(isFunc(fn[0]))
		hashMap.hashFunc = fn[0]
	} else {
		hashMap.hashFunc = hashFunc
	}

	return hashMap
}

// Get
func (self *HashMap) Get(key Key) (value interface{}, err error) {
	hashKey, bucketIdx := self.hashFunc(len(self.buckets), key)
	bucket := self.buckets[bucketIdx]
	for bucket != nil {
		if bucket.hashKey == hashKey && reflect.DeepEqual(key, bucket.key) {
			return bucket.value, nil
		}

		bucket = bucket.next
	}

	return nil, errors.New("Key not found!")
}

// Set
func (self *HashMap) Set(key Key, value interface{}) error {
	if self.loadFactor() >= growLoadFactor {
		//log.Printf("grow %d %d %d\n", self.loadFactor(), len(self.buckets), self.size)
		self.grow()
	}

	hashKey, bucketIdx := self.hashFunc(len(self.buckets), key)
	head := self.buckets[bucketIdx]
	self.buckets[bucketIdx] = &Bucket{hashKey, key, value, head}
	self.size++

	return nil
}

// Unset
func (self *HashMap) Unset(key Key) error {
	hashKey, bucketIdx := self.hashFunc(len(self.buckets), key)
	bucket := self.buckets[bucketIdx]
	if bucket == nil {
		return errors.New("Unset key not found")
	}

	var prev *Bucket
	for bucket != nil {
		if bucket.hashKey == hashKey && reflect.DeepEqual(key, bucket.key) {
			if prev == nil && bucket.next == nil {
				self.buckets[bucketIdx] = nil
				self.size--
			} else if prev == nil {
				self.buckets[bucketIdx] = bucket.next
			} else {
				prev.next = bucket.next
			}
		}
		prev = bucket
		bucket = bucket.next
	}

	if self.loadFactor() <= shrinkLoadFactor && self.shrinked {
		self.shrink()
	}

	return nil
}

// Count
func (self *HashMap) Count() int {
	return self.size
}

// Shrinking
func (self *HashMap) Shrinking(mode bool) {
	self.shrinked = mode
}

// Function for calculate load factor
func (self *HashMap) loadFactor() float32 {
	return float32(self.size) / float32(len(self.buckets))
}

// Rehash buckets
func (self *HashMap) rehash(blockSize int) {
	//	log.Printf("rehashBuckets %d\n", len(buckets))
	buckets := make([]*Bucket, blockSize)
	for i, bucket := range self.buckets {
		for bucket != nil {
			bucketIdx := bucket.hashKey % uint(blockSize)
			head := buckets[bucketIdx]
			buckets[bucketIdx] = &Bucket{bucket.hashKey, bucket.key, bucket.value, head}
			bucket = bucket.next
		}
		self.buckets[i] = nil
	}
	self.buckets = buckets
}

// Grow
func (self *HashMap) grow() {
	//log.Printf("grow\n")
	blockSize := len(self.buckets) * 2
	if self.defaultBlockSize >= blockSize {
		blockSize = self.defaultBlockSize
	}
	self.rehash(blockSize)
}

// Shrink
func (self *HashMap) shrink() {
	//	log.Printf("shrink\n")
	blockSize := len(self.buckets) / 2
	if self.defaultBlockSize <= blockSize {
		blockSize = self.defaultBlockSize
	}
	self.rehash(blockSize)
}

// iterate
func (self *HashMap) iterate(c chan<- KeyValue) {
	//	log.Printf("Iterate %s\n", c)
	for _, b := range self.buckets {
		for n := b; n != nil; n = n.next {
			c <- KeyValue{n.key, n.value}
		}
	}
	close(c)
}

// Iter
func (self *HashMap) Iter() <-chan KeyValue {
	//log.Printf("Iter\n")
	c := make(chan KeyValue)
	go self.iterate(c)
	return c
}

// toString
func (self *HashMap) String() string {
	s := "{"
	for r := range self.Iter() {
		s = s + fmt.Sprintf("%s: %s, ", r.key, r.value)
	}
	s = s + "}"
	return s
}

func isFunc(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}
