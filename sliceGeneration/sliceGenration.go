package slicegeneration

import (
	"hash/maphash"
	"math/rand"
	"sort"
)

type SplitIndexs struct {
	Start int
	End   int
}

func RandomPopulatedSlice(listSize int) []int {
	mpSeed := rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))
	// 1 << n is equal to 2^n
	//listSize = 1 << 27
	randomNumbers := make([]int, listSize)
	for i := 0; i < listSize; i++ {
		// insert number between 0 to 1 << n i.e 2^n
		randomNumbers[i] = mpSeed.Intn(1 << 10)
	}

	//fmt.Println(randomNumbers[:1000])
	return randomNumbers
}
func Orderedslice(listSize int, c string) []int {
	var randomNumbers []int
	// if "r" a new randomly populated list will be sorted and returned else
	if c == "r" {
		randomNumbers = RandomPopulatedSlice(listSize)

	} else {
		// else a list with fixed seed is sorted
		randomNumbers = SameRandList(listSize)

	}
	sort.Ints(randomNumbers)
	return randomNumbers

}

func SplitSlice(orderedlist []int, k int) []SplitIndexs {
	n := len(orderedlist)

	minElementsPerPart := n / k
	partsWithExtraElement := n % k

	indexs := make([]SplitIndexs, k)

	startIndex := 0

	for i := 0; i < k; i++ {
		// Calculate the size of the current part
		partSize := minElementsPerPart
		if i < partsWithExtraElement {
			partSize++
		}

		// Store the starting and ending indices for the current part
		indexs[i] = SplitIndexs{Start: startIndex, End: startIndex + partSize - 1}

		// Move the start index for the next part
		startIndex += partSize
	}

	return indexs

}

func SameRandList(listSize int) []int {
	Seed := rand.New(rand.NewSource(int64(42)))
	randomNumbers := make([]int, listSize)
	for i := 0; i < listSize; i++ {
		// insert number between 0 to 1 << n i.e 2^n
		randomNumbers[i] = Seed.Intn(1 << 10)
	}
	return randomNumbers
}
