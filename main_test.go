package ITA_Cormen

import (
	"fmt"
	"testing"
)

// 1
func TestInsertionSort (t *testing.T) {
	A := []int{5,2,4,32,10,101,5}
	fmt.Printf("Unsorted array: %d\n", A)
	fmt.Printf("Array sorted increasingly: %d\n", insertionSortIncreasing(A))
	fmt.Printf("Array sorted decreasingly: %d\n", insertionSortDecreasing(A))
}

// 2
func TestLinearSearch (t *testing.T) {
	A := []int{24,23,1,5,6,78,7,8}
	fmt.Println(linearSearch(A, 5))
	fmt.Println(linearSearch(A, 2))
}

// 3
func TestAddBinary (t *testing.T) {
	A := [5]byte{1,0,1,1,1}
	B := [5]byte{1,1,0,0,0}
	fmt.Println(addBinary(A, B))
}

// 4
func TestSelectionSort (t *testing.T) {
	// A is mixed
	fmt.Println("SORT A \n")
	A := []int{5, 4, 6, 2, 1, 7, 3}
	fmt.Println(selectionSort(A))

	// B is sorted increasingly
	fmt.Println("SORT B \n")
	B := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(selectionSort(B))

	// C is sorted decreasingly
	fmt.Println("SORT C \n")
	C := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println(selectionSort(C))
}

// 5
func TestMergeSort (t *testing.T) {
	A := []int{1, 5, 7, 2, 4, 5, 7, 1, 2, 3, 6, 8, 9, 10}
	fmt.Println("merge algorithm....")
	fmt.Println(merge(A, 3, 6, 10))
	// Array B has n elements where n is a power of 2 but works for other n elements as-well
	B := []int{5, 2, 4, 7, 1, 3, 2, 6}
	fmt.Println("mergeSort algorithm....")
	fmt.Println(mergeSort(B, 0, 7))
}

