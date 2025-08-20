package utils

import "fmt"

// Permutations returns all possible permutations of the input slice.
func Permutations(nums []int) [][]int {
	var res [][]int
	var helper func([]int, int)
	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			helper(arr, n-1)
			if n%2 == 1 {
				arr[0], arr[n-1] = arr[n-1], arr[0]
			} else {
				arr[i], arr[n-1] = arr[n-1], arr[i]
			}
		}
	}
	helper(nums, len(nums))
	return res
}

// SumSlice returns the sum of all elements in the slice.
func SumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

// PossibleMagicSquares checks how many magic squares can be formed with the given numbers.
func PossibleMagicSquares(nums []int) {
	magicSquaresFound := 0
	
	perms := Permutations(nums)	
	for _, perm := range perms {
		targetSum := perm[0] + perm[1] + perm[2]	
		count := 0

		// column 1 -> [0,3,6]
		col1 := []int{perm[0], perm[3], perm[6]}

		// column 2 -> [1,4,7]
		col2 := []int{perm[1], perm[4], perm[7]}

		// column 3 -> [2,5,8]
		col3 := []int{perm[2], perm[5], perm[8]}

		// main diagonal -> [0,4,8]
		diag1 := []int{perm[0], perm[4], perm[8]}

		// secondary diagonal -> [2,4,6]
		diag2 := []int{perm[2], perm[4], perm[6]}

		colsAndDiags := [][]int{col1, col2, col3, diag1, diag2}

		for _, col := range colsAndDiags {
			if SumSlice(col) == targetSum {
				count++
			}
		}

		if SumSlice(perm[3:6]) == targetSum && SumSlice(perm[6:9]) == targetSum && count == len(colsAndDiags) {
			fmt.Println("Magic square found:")
			fmt.Println(perm[0:3])
			fmt.Println(perm[3:6])
			fmt.Println(perm[6:9])
			fmt.Println()
			magicSquaresFound++
		}
	}

	if magicSquaresFound == 0 {
		fmt.Println("No magic square can be formed with these numbers.")	
	}
}
