package main

import (
	"log"
	"strconv"
)

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

func minPartitions(n string) int {
	g := 0
	for _, v := range n {
		ch, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal(err)
		}
		for g < ch {
			g++
		}
	}
	return g
}

//Поиск самого маленького элемента массива
func Smallest(a []int) int {
	smallest := a[0]
	for _, v := range a {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}

//Сортировка выбором
func SelectionSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		min := i
		for j := i + 1; j < len(ar); j++ {
			if ar[min] > ar[j] {
				min = j
			}
		}
		if min != i {
			swap(ar, i, min)
		}
	}
}
func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}
