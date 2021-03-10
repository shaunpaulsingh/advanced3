package main

import (
	"advanced3/arrays"
	"fmt"
)

func main() {
	var testArray = make(arrays.SortedArray, 2000)
	for i := 0 ; i <= 1000 ; i ++{
		testArray[i] = (1000-i) * 2
	}

	//testArray.Sort()
	//testArray.InsertionSort()
	testArray.MergeSort()
	//testArray.QuickSort()

	fmt.Println(testArray)
	//testArray.SequentialSearch(1000)
	//testArray.BinarySearch(1000)
	fmt.Println(testArray.RecursiveSequentialSearch(1000))
	fmt.Println(testArray.RecursiveBinarySearch(1000))
}
