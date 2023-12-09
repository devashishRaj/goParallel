package main

import (
	slicegeneration "goParallel/sliceGeneration"
	"goParallel/sorting"
)

func main(){
	input := slicegeneration.SameRandList(1 << 20)
	sorting.QuickSortInPlace(input)

}