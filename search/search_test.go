package search

import (
	slicegeneration "goParallel/sliceGeneration"
	"sort"
	"testing"
)

func TestSeqBinary(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	sort.Ints(input)
	index :=BinarySearch(input, target, 0, len(input))
	assertSearch(t , target, index)
	
}

func TestParallelSrch(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	// Print input array for debugging
	t.Logf("Input array: %v", input)

	targetIndex := ParallelBinary(input, 4, target)
	// Print additional debug information
	assertSearch(t , target, targetIndex)
	t.Logf("Target found at index: %d", targetIndex)
	t.Logf("Value at target index: %d", input[targetIndex])
	if input[targetIndex] != target {
		t.Fatalf("Found incorrect value at index %d: got %d, expected %d", targetIndex, input[targetIndex], target)
	}
}

func BenchmarkSeqBinary(b *testing.B) {
	targetLimit := 1 << 10
	list := slicegeneration.Orderedslice(1 << 27 , "s")

	for i := 0; i < b.N; i++ {
		target := i % targetLimit
		index := BinarySearch(list, target, 0, len(list))
		if index == -1 {
			b.Fatal("Value not found")
		}
	}
}

func BenchmarkParallelBinary(b *testing.B) {
	targetLimit := 1 << 10
	list := slicegeneration.Orderedslice(1 << 27 , "s")
	for i := 0; i < b.N; i++ {
		t := i % targetLimit
		target := list[t]
		index := ParallelBinary(list, 100000, target)
		if index == -1 {
			b.Fatal("Value not found")
		}
	}
}

func TestListSearch(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	index := ListSearch(input, target, 0, len(input))
	assertSearch(t , target, index)
	t.Logf("Input array: %v", input)
	t.Logf("Target found at index: %d", index)
	t.Logf("Value at target index: %d", input[index])
}

func TestParallelListSearch(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7
	index := ParallelListSearch(input, 4, target)
	assertSearch(t , target, index)
	t.Logf("Input array: %v", input)
	t.Logf("Target found at index: %d", index)
	t.Logf("Value at target index: %d", input[index])
}

func BenchmarkListSearch(b *testing.B) {
	targetLimit := 1 << 22
	list := slicegeneration.RandomPopulatedSlice(1 << 27)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target := i % targetLimit
		index := ListSearch(list, target, 0, len(list))
		if index == -1 {
			b.Fatal("Value not found")
		}
	}
}

func BenchmarkParrallelListSearch(b *testing.B) {
	targetLimit := 1 << 22
	list := slicegeneration.RandomPopulatedSlice(1 << 27)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := i % targetLimit
		target := list[t]
		index := ParallelListSearch(list, 8, target)
		if index == -1 {
			b.Fatal("Value not found")
		}
	}
}

func BenchmarkListCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slicegeneration.RandomPopulatedSlice(1 << 27)
	}
}

func BenchmarkListCreationAndSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slicegeneration.Orderedslice(1 << 27 , "r")
	}
}

func assertSearch(t testing.TB ,target , index int ) {
	t.Helper()
	if index == -1 {
		t.Fatalf("Target %d not found , -1 returned" ,target)
	}

}
