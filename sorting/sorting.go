package sorting

//TODO: look into weight group or list size.
//TODO: debuging sequential and parallel control flow

// avg case time complexity is O(N * logN).
func QuickSortInPlace(arr []int) {
	quickSortRecursive(arr, 0, len(arr)-1)
}

func quickSortRecursive(arr []int, low, high int) {
	if low < high {
		// Use insertion sort for small subarrays
		if high-low+1 <= 20 {
			insertionSort(arr, low, high)
			return
		}

		pivotIndex := partition(arr, low, high)

		// Recursively sort the subarrays
		quickSortRecursive(arr, low, pivotIndex-1)
		quickSortRecursive(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func insertionSort(arr []int, low, high int) {
	for i := low + 1; i <= high; i++ {
		key := arr[i]
		j := i - 1

		for j >= low && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
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
