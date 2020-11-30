package main

import "examples.go/set/hashsets"

/*
### Description

Given an array of values. Determine if there is a duplicate one

### Example 1:

```
Input: {1, 2, 3}
Output: false
```

### Example 2:

```
Input: {"key1", "key2", "key1"}
Output: true
```
*/

// ContainsDuplicates checks if the list contains duplicates
func ContainsDuplicates(values ...interface{}) bool {
	set := hashsets.New()
	for _, value := range values {
		if set.Contains(value) {
			return true
		}
		set.Add(value)
	}
	return false
}
