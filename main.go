package ITA_Cormen

import "fmt"

func main() {
	
}

// 1. Insertion Sort Algorithm
func insertionSortIncreasing (n []int) []int {

	for j := 1; j <= (len(n)-1); j++ {
		// This is the selected key to which we will compare our values inside the sorted array
		key := n[j]
		// Insert n[j] into the sorted sequence sortedList[1...j-1]
		i := j-1
		for i >= 0 && n[i] > key {
			n[i+1] = n[i]
			i = i-1
		}
		n[i+1] = key
	}
	return n
}
// ex 2.1-2
func insertionSortDecreasing (n []int) []int {

	for j := 1; j <= (len(n)-1); j++ {
		// This is the selected key to which we will compare our values inside the sorted array
		key := n[j]
		// Insert n[j] into the sorted sequence sortedList[1...j-1]
		i := j-1
		for i >= 0 && n[i] < key {
			n[i+1] = n[i]
			i = i-1
		}
		n[i+1] = key
	}
	return n
}

// ex 2.1-3
// 2. Linear Search Algorithm
type result interface {}
// input parameters n an unsorted array and v the integer value to be searched inside the list
func linearSearch (n []int, v int) result {
	for i := 0; i<len(n); i++ {
		if n[i] == v {
			return i
		}
	}
	return nil
}

// ex 2.1-4
// 3. Add 2 Binary Integers
func addBinary(A,B [5]byte) [6]byte {
	var carry byte
	carry = 0

	var C [6]byte

	// A[0] ... A[n-1] (length = n)
	// B[0] ... B[n-1] (length = n)
	// C[0] ... C[n] (length = n+1)

	// Implying both byte arrays have the same length n
	for i := len(A)-1; i>=0; i--{
		C[i+1] = (A[i] + B[i] + carry)%2
		carry = (A[i] + B[i] + carry)/2
		C[0] = carry
	}
	return C
}

// ex 2.2-2
// 4. Selection Sort Algorithm
func selectionSort (a []int ) []int {
	n := len(a)
	swapCounter := 0
	for j := 0; j<n; j++ {
		// we assume the smallest element so far has always the last "j" index where we left off
		iMin := j

		// here we check the current "j" element with the next one and the next one and so on
		for i := j+1; i<n; i ++ {
			// replacing the iMin index only if it the elements are smaller then the current one being checked
			if a[i] < a[iMin] {
				iMin = i // in the end we will always have the iMin value the smallest number's index from the series
			}
		}
		// if index number is not the same, we swap the two slice elements
		if iMin != j {
			fmt.Printf("succes for index %d with value %d\n", iMin, a[iMin])
			fmt.Printf("swaping %d with %d \n\n", a[j], a[iMin])
			a[j], a[iMin] = a[iMin], a[j]
			swapCounter++
		}

	}
	fmt.Printf("Swapped %d times\n", swapCounter)
	return a
}

// 5. Merge Sort Algorithm
func merge (A []int, p, q, r int) []int {
	// Computes the length n1 of the subarray A[p...q]
	n1 :=  q - p + 1

	// Computes the length n2 of the subarray A[q+1...r]
	n2 := r - q

	// We create the arrays "L" and "R" ("left" and "right") of lengths "n1+1" and "n2+1"
	// the extra position in each array will hold the sentinel
	lengthL := n1 + 1; lengthR := n2 + 1

	L := make([]int, lengthL)
	R := make([]int, lengthR)

	// The for loop copies the subarray A[p...q] into L[0...(n1-1)]
	for i := 0; i < n1; i++ { //
		L[i] = A[p + i]
	}

	// The for loop copies the subarray A[q+1...r] into R[0...(n2-1)]
	for j := 0; j < n2; j++ {
		R[j] = A[q + 1 + j]
	}

	// Put the sentinels at the ends of the arrays L and R
	L[lengthL-1] = 123456789
	R[lengthR-1] = 123456789

	// n = r - p + 1 is the total number of items being merged
	i := 0
	j := 0

	// At the start of each iteration of the for loop the subarray A[p...(k-1)]
	// contains the (k-p) smallest elements of L[0...((n1-1)+1)] and R[0...((n2-1)+1)]
	// in sorted order. Moreover L[i] and R[j] are the smallest elements of their arrays
	// that have not been copied back into A.
	for k := p; k <= r; k++ {
		if L[i] <= R[j]{
			A[k] = L[i]
			i = i + 1
		} else {
			A[k] = R[j]
			j = j + 1
		}
	}
	// Returning array with sorted elements
	return A
}

func mergeSort (A []int, p, r int) []int {
	var q int
	if p < r {
		q = (p + r) / 2
		mergeSort(A, p, q)
		mergeSort(A, q+1, r)
		merge(A, p, q, r)
	}
	return A
}

// ex 2.3-4
// 6. Insertion Sort Recursive version
func recurseInsertionSort (A []int, n int) []int {
	if n >= 1 {
		recurseInsertionSort(A , n-1)
		insert(A, n)
	}
	return A
}

func insert(A []int, k int) {
	key := A[k]
	i := k-1
	for i >= 0 && A[i] > key {
		A[i+1] = A[i]
		i = i - 1
	A[i+1] = key
	}
}

// ex 2.3-5
// 7. Binary Search Algorithm (recursive + iterative, my version)
// (if A is a Sorted List)
type resultBinary interface {}
// input parameters n an unsorted array and v the integer value to be searched inside the list
func binarySearch (n []int, x int) resultBinary {
	for i := 0; i<len(n); i++ {
		if n[i] == x {
			return i
		}
	}
	return nil
}

func binarySearchRecurse (n []int, x int, counter int) resultBinary {
	limitInf := len(n)
	// We use this counter to see how many recurrences are made
	if x == 0 || x<0  {
		fmt.Println("not supported for non positive elements")
		return nil
	}
	counter ++
	if x <= limitInf {
		limitInf := len(n) / 2
		fmt.Printf("recurrence: %d\n", counter)
		binarySearchRecurse(n[:limitInf-1], x, counter)
		ret := binarySearch(n, x)
		return ret
	}
	return nil
}

// ex 2.3-5
// 8. Binary Search Algorithm (recursive only)
// n - sorted Array, x - searched value
// low - 0 index, high - n index
// T(n)=T((n−1)/2)+Θ(1)=Θ(lgn)
type resultRecursive interface {}
func recursiveBinarySearch (n []int, x, low, high int) resultRecursive {
	if low > high {
		return nil
	}
	mid := (low + high) / 2
	if x == n[mid] {
		fmt.Println("OK")
		return mid
	}
	if x > n[mid] {
		ret := recursiveBinarySearch(n, x, mid+1, high)
		return ret
	} else {
		ret := recursiveBinarySearch(n, x, low, mid-1)
		return ret
	}
}
// 9. Binary Search Algorithm (iterative)
type resultIterativeBinarySearch interface {}
func iterativeBinarySearch (A []int, x int) resultIterativeBinarySearch {
	length := len(A)
	low := A[0]
	high := A[length-1]
	for low <= high {
		mid := (low + high) / 2
		if x == A[mid] {
			return mid
		}
		if x > A[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

// ex 2.3-7
// Determines if two elements in S are the sum of X
func pairExists (S []int, X int) bool {
	A := mergeSort (S, 0, len(S)-1)
	for i:= 0; i<len(A); i++ {
		if binarySearch (A, X-S[i]) != nil {
			fmt.Printf("X element %d, S[i] element %d, X-S[i] element %d\n", X, S[i], X-S[i])
			return true
		}
 	}
	return false
}

// 10. Bubble Sort Algorithm
// this function swaps the two elements of the array
// A - Array; X,Y - index of elements being swapped
func swapElements (A[]int, X,Y int) {
	A[X], A[Y] = A[Y], A[X]
}

func bubbleSort (A []int) []int {
	lastIndex := len(A) - 1 // last element index
	for i:= 0; i < lastIndex - 1; i++ {
		for j := lastIndex; j > i; j-- {
			fmt.Printf("J %d J-a %d\n", j, j-1)
			if A[j] < A[j-1] {
				swapElements(A, j, j-1)
			}
		}
	}
	return A
}

// ex 2-3
// 11. Horner's Rule for evaluating a polynomial
// given the coefficients a0,a1,...an and a value for x
func hornerPolynomial (a[]int, x int) int {
	y := 0
	lastIndex := len(a) - 1
	for i := lastIndex ; i >= 0 ; i-- {
		y = a[i] + x * y
	}
	return y
}

// ex 2-4
// 12. Inversions
func inversion (a[]int) [][]int {
	var listInversions [][]int
	counter := 0
	for i:= 0; i < len(a); i++ {
		for j:= 0; j < len(a); j++ {
			if i < j && a[i]>a[j] {
				listInversions = append(listInversions,[]int{a[i], a[j]})
				counter ++
			}
		}
	}
	fmt.Printf("%d inversions\n", counter)
	return listInversions
}

// 13 The maximum-subarray problem

// Maximum Crossing Subarray Finding Algorithm
func findMaxCrossingSubarray(A[]int, low, mid, high int) []int {
	// Find a maximum subarray of the left half
	leftSum := -123456789
	maxLeft := 0
	sumOne := 0
	for i:= mid; i >= low; i-- {
		sumOne = sumOne + A[i]
		if sumOne > leftSum {
			leftSum = sumOne
			maxLeft = i
		}
	}

	// Find a maximum subarray of the right half
	rightSum := -123456789
	maxRight := 0
	sumTwo := 0
	for j := mid + 1; j <= high; j++ {
		sumTwo = sumTwo + A[j]
		if sumTwo > rightSum {
			rightSum = sumTwo
			maxRight = j
		}
	}

	// Combine the results (Left and Right Indexes and the sum of the integers)
	ret := []int{maxLeft, maxRight, leftSum + rightSum}
	return ret
}

// Maximum Subarray Finding Algorithm
func findMaximumSubarray (A[]int, low, high int) []int {
	if high == low { 									// Base case: only one element
		var ret []int
		ret = append(ret, low, high, A[low])
		return ret
	} else {
		// Middle index of array
		mid := (low + high) / 2

		// Left Maximum Subarray
		// (left-low, left-high, left-sum)
		leftArray := findMaximumSubarray(A, low, mid)
		leftLow := leftArray[0]
		leftHigh := leftArray[1]
		leftSum := leftArray[2]
		// Right Maximum Subarray
		// (right-low, right-high, right-sum)
		rightArray := findMaximumSubarray(A, mid+1, high)
		rightLow := rightArray[0]
		rightHigh := rightArray[1]
		rightSum := rightArray[2]

		// Crossing Maximum Subarray
		// (cross-low, cross-high, cross-sum)
		crossArray := findMaxCrossingSubarray(A, low, mid, high)
		crossLow := crossArray[0]
		crossHigh := crossArray[1]
		crossSum := crossArray[2]

		if leftSum >= rightSum && leftSum >= crossSum {
			ret := []int{leftLow, leftHigh, leftSum}
			return ret
		}
		if rightSum >= leftSum && rightSum >= crossSum {
			ret := []int{rightLow, rightHigh, rightSum}
			return ret
		} else {
			ret := []int{crossLow, crossHigh, crossSum}
			return ret
		}
	}
}
