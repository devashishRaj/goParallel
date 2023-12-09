package sorting

import (
	slicegeneration "goParallel/sliceGeneration"
	"sort"
	"testing"
)

func TestQuick_sort(t *testing.T) {
	input := []int{5, 6, 7, 8, 9, 1, 2, 3, 4, 10}
	arr := input
	QuickSortInPlace(arr)
	assertSort(t, arr, input)
	t.Logf(" %v ", input)
}

func BenchmarkQuick_sort(b *testing.B) {
	input := slicegeneration.SameRandList(1 << 27)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ip := input
		QuickSortInPlace(ip)
		//assertSort(b, ip, input)
	}
}

func TestParallelQuicksort(t *testing.T) {
	input := slicegeneration.SameRandList(1 << 27)
	arr := input
	//sort.Ints(arr) // takes 8.8
	QuickSortParallel(arr)
	assertSort(t, arr, input)
}

func assertSort(t testing.TB, sorted, orignal []int) {
	t.Helper()
	if !sort.SliceIsSorted(sorted, func(i, j int) bool { return sorted[i] < sorted[j] }) {
		// t.Logf("%v", sorted)
		// t.Logf("%v", orignal)
		t.Fatalf("not sorted")
	}
}

func Benchmark(b *testing.B) {
	input := slicegeneration.SameRandList(1 << 20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arr := input
		QuickSortParallel(arr)
		//assertSort(b, arr, input)
	}
}
