package minmeetingrooms

import "github.com/Dev-Snippets/algorithm-go-snippets/datastructures/sets/hashmultisets"

/**
### Description

Given an array of meeting time intervals consisting of start and end times, find the minimum number of conference rooms required.

### Example 1:

```
Input: [[0, 30],[5, 10],[15, 20]]
Output: 2
```

### Example 2:

```
Input: [[7,10],[2,4]]
Output: 1
```
**/

// MinMeetingRooms returns minimum meeting rooms required
func MinMeetingRooms(intervals [][]int) int {
	multiset := hashmultisets.New()
	for _, interval := range intervals {
		for i := interval[0]; i <= interval[1]; i++ {
			multiset.Add(i)
		}
	}

	multiSetPairs := multiset.GetTopValues()
	if len(multiSetPairs) == 0 {
		return 0
	}
	return multiSetPairs[0].Count
}
