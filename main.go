package main

import (
	"sort"
	"strconv"
)

// func Num(n int) {
// 	fmt.Println(n)
// }

// func smallestEventMultiple(n int) int {

// 	if n%2 == 0 {

// 		return n
// 	} else {
// 		return n * 2
// 	}

// }

func numIdenticalPairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				count += 1
			}
		}
	}
	return count
	// if nums[i] == nums[j] and i < j
}

func minimumSum(num int) int {
	var dig []int
	for _, v := range strconv.Itoa(num) {

		dig = append(dig, int(v)-48)
	}

	sort.Ints(dig)

	return dig[0]*10 + dig[1]*10 + dig[2] + dig[3]
	//return 77
}
func main() {
	// fmt.Println(minimumSum(2894))
	// arr := [4]int{
	// 	(2894 % 10000) / 1000,
	// 	(2894 % 1000) / 100,
	// 	(2894 % 100) / 10,
	// 	(2894 % 10),
	// }
	// fmt.Println(arr)
	// ch := make(chan int)
	// // ch <- 111
	// go func() {
	// 	ch <- 111
	// }()

	// val := <-ch
	// fmt.Printf("got %v\n", val)
	// for i := 0; i < 5; i++ {
	// 	ii := i
	// 	go func() {

	// 		Num(ii)
	// 	}()
	// }
	// time.Sleep(100 * time.Millisecond)
}

// func foobyval(n int) {
// 	fmt.Println(n)
// }

// func main() {
// 	for i := 0; i < 5; i++ {
// 		go foobyval(i)
// 	}

// 	time.Sleep(100 * time.Millisecond)
// }

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
// func Add(x int, y int) int {
// 	return x + y
// }
// func Swap(x, y string) (string, string) {
// 	return y, x
// }
// func fact(x int) int {
// 	if x == 1 {
// 		return 1
// 	} else {
// 		return x * fact(x-1)
// 	}
// }

// func main() {
// values := []string{"1, 2, 3"}
// for _, val := range values {
// 	go func(val string) {
// 		fmt.Println(val)
// 	}(val)
// }
// text := "hello my name is Tima"
// arr := strings.Fields(text)
// for _, v := range arr {
// 	fmt.Println(v)
// }
// var myGreeting map[string]string

// fmt.Println(myGreeting)
// fmt.Println(myGreeting == nil)
// fmt.Println(fact(4))
// CountDown(4)
// count := []int{}
// for i := 0; i < len(count); i++ {
// }s
// a := []int{14, 12, 10, 9}
// SelectionSort(a)
// fmt.Println(a)
// minPartitions("1234")
// a, b := Swap("world", "hello")
// fmt.Println(a, b)
// fmt.Println(Add(42, 13))
// fmt.Print(math.Pi)
// fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
//Random number
// fmt.Println("the number is", rand.Seed())
// fmt.Println("My favorite number is", rand.Intn(10))
// items := []int{1, 3, 5, 9, 12}
// fmt.Println(BinarySearch(items, 6))
// x := math.Log2(256)
// fmt.Println(x)
// fmt.Println(isPaindrome(334))
// words := regexp.MustCompile("[,%_ ]{1}").Split(refString, -1)
// for idx, word := range words {
// 	fmt.Printf("Word %d is: %s\n", idx+1, word)
// }
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
