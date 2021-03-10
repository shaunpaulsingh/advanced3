package arrays

import (
	"fmt"
	"sort"
)

type SortedArray []int

func (arr *SortedArray)Sort(){
	sort.Ints(*arr)
}

func (arr *SortedArray)SelectionSort(){
	len := len(*arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if (*arr)[j] < (*arr)[minIndex] {
				(*arr)[j], (*arr)[minIndex] = (*arr)[minIndex], (*arr)[j]
			}
		}
	}
}

func (arr *SortedArray)InsertionSort(){

	len := len(*arr)
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if (*arr)[j] > (*arr)[i] {
				(*arr)[j], (*arr)[i] = (*arr)[i], (*arr)[j]
			}
		}
	}
}

func (arr *SortedArray)SequentialSearch(toMatch int) int{
	var numComparisons int = 0

	for i,v := range *arr{
		numComparisons = numComparisons + 1
		if v == toMatch {
			fmt.Println(numComparisons)
			return i
		}
	}
	return -1
}

func (arr *SortedArray)RecursiveSequentialSearch(toMatch int) int{
	return (*arr).rSeq(toMatch,0)
}

func (arr *SortedArray)rSeq(toMatch int, currentIndex int) int{
	if (*arr)[currentIndex] == toMatch{
		return currentIndex
	}else{
		return arr.rSeq(toMatch,currentIndex + 1)
	}
}

func (arr *SortedArray)RecursiveBinarySearch(toMatch int) int{
	return arr.rBinary(toMatch, 0, len(*arr))
}

func (arr *SortedArray)rBinary(toMatch int, first int, last int) int{
	mid := (first + last) / 2
	if (*arr)[mid] == toMatch{
		return mid
	}else{
		if toMatch < (*arr)[mid]{
			return arr.rBinary(toMatch,first,mid-1)
		}else{
			return arr.rBinary(toMatch, mid+1,last)
		}
	}
	return -1
}

func (arr *SortedArray)BinarySearch(toMatch int) int{
	var numComparions int = 0
	first := 0
	last := len(*arr)

	for first <= last {
		mid := (first + last) / 2

		numComparions = numComparions + 1
		if (*arr)[mid] == toMatch{
			fmt.Println(numComparions)
			return mid
		}else{
			numComparions = numComparions + 1
			if toMatch < (*arr)[mid]{
				last = mid - 1
			}else{
				first = mid + 1
			}
		}
	}
	return -1
}

func (arr *SortedArray)QuickSort(){
	arr.rQuickSort(0, len(*arr)-1)
}

func (arr *SortedArray)rQuickSort(start int , end int) {
	if (end - start) < 1 {
		return
	}

	pivot := (*arr)[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if (*arr)[i] < pivot {
			temp := (*arr)[splitIndex]

			(*arr)[splitIndex] = (*arr)[i]
			(*arr)[i] = temp

			splitIndex++
		}
	}

	(*arr)[end] = (*arr)[splitIndex]
	(*arr)[splitIndex] = pivot

	arr.rQuickSort(start, splitIndex-1)
	arr.rQuickSort(splitIndex+1,end)
}

func (arr *SortedArray)MergeSort(){
	*arr = rMergeSort(*arr)
}

func rMergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(rMergeSort(left), rMergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left) + len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

func (arr *SortedArray)HybridSort(){
	
}