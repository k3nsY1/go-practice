package main

import "strconv"

//Бинарный поиск
func BinarySearch(list []int, item int) bool {
	low := 0
	high := len(list) - 1
	// fmt.Println(len(list))
	for low <= high {

		mid := (low + high) / 2
		if list[mid] < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low == len(list) || list[low] != item {
		return false
	}
	return true
}

//Проверка на палиндром к примеру 121, 543212345 и тд.
func isPalindrome(x int) bool {
	//convert int to slice of runes
	runes := []rune(strconv.Itoa(x))

	//reverse the runes in the slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	//convert back to int, via string
	y, err := strconv.Atoi(string(runes))

	//check for an error (overflow) and equality
	if err == nil && x == y {
		return true
	}
	return false
}
