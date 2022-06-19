package main

// Реализация интерфейса, пример
// type vehicle interface {
// 	move()
// }

// type Car struct{}

// func (c Car) move() {
// 	fmt.Println("Автомобиль едет")
// }

// func main() {
// 	var bmw vehicle = Car{}

// 	bmw.move()
// }

//Работа со строками
// const valString = "Charlie, today is a good day"

// const refString = "Mary had a little lamb"

func main() {
	// items := []int{1, 3, 5, 9, 12}
	// fmt.Println(BinarySearch(items, 6))
	// x := math.Log2(256)
	// fmt.Println(x)
	// fmt.Println(isPaindrome(334))
	// words := regexp.MustCompile("[,%_ ]{1}").Split(refString, -1)
	// for idx, word := range words {
	// 	fmt.Printf("Word %d is: %s\n", idx+1, word)
	// }
}

// func BinarySearch(list []int, item int) bool {
// 	low := 0
// 	high := len(list) - 1
// 	// fmt.Println(len(list))
// 	for low <= high {

// 		mid := (low + high) / 2
// 		if list[mid] < item {
// 			low = mid + 1
// 		} else {
// 			high = mid - 1
// 		}
// 	}
// 	if low == len(list) || list[low] != item {
// 		return false
// 	}
// 	return true
// }

// func isPaindrome(x int) bool {
// 	//convert int to slice of runes
// 	runes := []rune(strconv.Itoa(x))

// 	//reverse the runes in the slice
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	//convert back to int, via string
// 	y, err := strconv.Atoi(string(runes))

// 	//check for an error (overflow) and equality
// 	if err == nil && x == y {
// 		return true
// 	}
// 	return false
// }

// words := strings.FieldsFunc(refString, unicode.IsSpace)
// for idx, word := range words {
// 	fmt.Printf("Word %d is: %s\n", idx+1, word)
// }
// splitFunc := func(r rune) bool {
// 	return strings.ContainsRune("*%,_ ", r)
// }

// words := strings.FieldsFunc(valString, splitFunc)
// for idx, word := range words {
// 	fmt.Printf("Word %d is: %s\n", idx+1, word)
// }
// }
