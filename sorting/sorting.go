package sorting

//TODO: look into weight group or list size.
//TODO: debuging sequential and parallel control flow

import (
	"sync"
)

// func QuickSortInPlace(arr []int, low, high int) {
// 	if low < high {
// 		var pivot = partition(arr, low, high)
// 		QuickSortInPlace(arr, low, pivot)
// 		QuickSortInPlace(arr, pivot + 1, high)
// 	}
// }

// func partition(arr []int, low, high int) int {
// 	var pivot = arr[low]
// 	var i = low
// 	var j = high

// 	for i < j {
// 		for arr[i] <= pivot && i < high {
// 			i++;
// 		}
// 		for arr[j] > pivot && j > low {
// 			j--
// 		}

// 		if i < j {
// 			var temp = arr[i]
// 			arr[i] = arr[j]
// 			arr[j] = temp
// 		}
// 	}

// 	arr[low] = arr[j]
// 	arr[j] = pivot

// 	return j
// }

type StackFrame struct {
	low  int
	high int
}

// avg case time complexity is O(N * logN).
func QuickSortInPlace(arr []int) {
	//if true then arr is now sorted
	if len(arr) <= 1 {
		return
	}

	stack := make([]StackFrame, 0)
	// low an high of full slice
	stack = append(stack, StackFrame{low: 0, high: len(arr) - 1})

	for len(stack) > 0 {
		// frame stores the last element and stack remove it : pop
		frame := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		low, high := frame.low, frame.high
		pivotIndex := partition(arr, low, high)

		if pivotIndex-1 > low {
			// add stackfram with indices of low and element just before pivot
			stack = append(stack, StackFrame{low: low, high: pivotIndex - 1})
		}
		if pivotIndex+1 < high {
			// add stackframe with indices of element just after pivot and high
			stack = append(stack, StackFrame{low: pivotIndex + 1, high: high})
		}
	}
}

func partition(arr []int, low, high int) int {
	// Choose a random index between low and high (inclusive)
	//pivotIndex := rand.Intn(high-low+1) + low
	// Swap the randomly chosen pivot with the last element
	//arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			//swap
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// move pivot position where element at right a smaller or equal and at left side greater
	arr[i+1], arr[high] = arr[high], arr[i+1]
	// return pivot index
	return i + 1
}

type Worker struct {
	ID    int
	Job   chan StackFrame
	Wg    *sync.WaitGroup
	Array []int
}

// recursive
func MergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	var mid = len(list) / 2
	var a = MergeSort(list[:mid])
	var b = MergeSort(list[mid:])
	return merge(a, b)
}

func merge(a, b []int) []int {
	var k = make([]int, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {

		if a[i] <= b[j] {
			k[i+j] = a[i]
			i++
		} else {
			k[i+j] = b[j]
			j++
		}

	}

	for i < len(a) {
		k[i+j] = a[i]
		i++
	}
	for j < len(b) {
		k[i+j] = b[j]
		j++
	}

	return k
}

func QuickSortParallel(list []int) []int {
	var result []int

	return result
}
