package main

/** 交换数据 */
func swap(array []int, x, y uint) {
	temp := array[x]
	array[x] = array[y]
	array[y] = temp
}

// 将数组根据 pivotLocation 划分为两段
func partition(array []int, start uint, end uint, pivotLocation uint) uint {

	// 首先将枢轴交换到尾部
	pivot := array[pivotLocation]
	swap(array, pivotLocation, end)

	// 将所有小于枢轴值的移动到数组前部
	i := start
	for j := start; j < end; j++ {
		if array[j] <= pivot {
			swap(array, i, j)
			i++
		}
	}

	// 将枢轴值交换到中间位置
	swap(array, end, i)

	return i
}

// 快速排序
func QuickSort(array []int, start uint, end uint) {
	if start < end {
		pivot := (end + start) / 2
		r := partition(array, start, end, pivot)
		if r > start {
			QuickSort(array, start, r-1)
		}
		QuickSort(array, r+1, end)
	}
}

func main() {
	arr := []int{1, 2, 0}
	QuickSort(arr, 0, uint(len(arr)-1))
}
