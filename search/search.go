package search

import (
	slicegeneration "goParallel/sliceGeneration"
	"runtime"
	"sync"
)

func BinarySearch(arr []int, target int, low int, high int) int {

	//high = high -1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1 // target not found
}

func ParallelBinary(list []int, parts int, target int) int {
	resultChan := make(chan int, parts)
	//doneCh := make(chan int , 1)
	splitList := slicegeneration.SplitSlice(list, parts)
	var wg sync.WaitGroup

	//found := false // Flag to indicate if a successful result has been found
	for _, SplitIndexs := range splitList {
		wg.Add(1)
		start := SplitIndexs.Start
		end := SplitIndexs.End
		// if end == len(list){
		// 	end = SplitIndexs.End -1
		// }
		go func(start, end int) {
			defer wg.Done()
			if target >= start && target <= end {
				index := BinarySearch(list, target, start, end)
				if index != -1 {
					resultChan <- index
				}
			}
		}(start, end)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	// Collect results from the channel
	for result := range resultChan {
		if result != -1 {
			// If the target is found by any goroutine, break out of the loop and return the result
			return result
		}
	}
	// If no goroutine found the target
	return -1
}

func ListSearch(list []int, target int, start, end int) int {
	for i := start; i <= end; i++ {
		if list[i] == target {
			return i
		}
	}
	return -1
}

func ParallelListSearch(list []int, parts int, target int) int {
	runtime.GOMAXPROCS(10)
	resultChan := make(chan int, parts)
	splitList := slicegeneration.SplitSlice(list, parts)
	var wg sync.WaitGroup

	//found := false // Flag to indicate if a successful result has been found
	for _, SplitIndexs := range splitList {
		wg.Add(1)
		start := SplitIndexs.Start
		end := SplitIndexs.End
		if end == len(list) {
			end = SplitIndexs.End - 1
		}
		go func(start, end int) {
			defer wg.Done()
			index := ListSearch(list, target, start, end)
			if index != -1 {
				resultChan <- index
			}
		}(start, end)
	}

	wg.Wait()
	close(resultChan)

	// Collect results from the channel
	for result := range resultChan {
		if result != -1 {
			// If the target is found by any goroutine, break out of the loop and return the result
			return result
		}
	}
	// If no goroutine found the target
	return -1
}
