package game

import (
	"math/rand"
	"strconv"
	"time"
)

var (
	initializedSeed = false
)

// Shuffle slice
func Shuffle(a []int) {
	if !initializedSeed {
		rand.Seed(time.Now().UnixNano())
	}
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

// Equals compares guess and answer that are equal
func Equals(guess, answer []int) bool {
	if len(guess) != len(answer) {
		return false
	}
	for i, g := range guess {
		if answer[i] != g {
			return false
		}
	}
	return true
}

// Str2Int convert []string to []int
// string that was not int will be ignored (is not filled in)
func Str2Int(str []string) []int {
	res := make([]int, 0)
	for _, s := range str {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		res = append(res, i)
	}
	return res
}

// GetMooNum returns moo number with digit
func GetMooNum(digit int) []int {
	sl := make([]int, len(nums))
	copy(sl, nums)
	Shuffle(sl)
	answer := sl[:digit]
	return answer
}

// GetAllCandidates Reference: https://medium.com/weekly-webtips/step-by-step-guide-to-array-permutation-using-recursion-in-javascript-4e76188b88ff/*
func GetAllCandidates(digits int) [][]int {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var getPermute func(arr []int, n int) [][]int
	getPermute = func(arr []int, n int) [][]int {

		// If we need 0 more digits we need to return one "empty" permutation.
		// as this allows us to build up permutations of given length digit.
		if n == 0 {
			return [][]int{{}} // slice containing one empty slice
		}

		var res [][]int

		for i := 0; i < len(arr); i++ {
			curNum := arr[i]

			remainNums := make([]int, 0, len(arr)-1)
			remainNums = append(remainNums, arr[:i]...)
			remainNums = append(remainNums, arr[i+1:]...)

			// Recursively permute the remainder for n-1 times
			subPermutation := getPermute(remainNums, n-1)

			// Then for each subPermutation, we prepend currentNum
			for _, sp := range subPermutation {
				newPermutation := append([]int{curNum}, sp...)
				res = append(res, newPermutation)
			}
		}
		return res
	}
	return getPermute(nums, digits)
}
