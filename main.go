package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxValue := 0
	for _, v := range candies {
		if v > maxValue {
			maxValue = v
		}
	}
	t := []bool{}
	for _, v := range candies {
		if v+extraCandies >= maxValue {
			t = append(t, true)
		} else {
			t = append(t, false)
		}
	}
	return t
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	t := root
	if t.Val == t.Left.Val+t.Right.Val {
		return true
	} else {
		return false
	}
}

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

// func qSort(a []int) []int {
// 	if len(a) < 2 {
// 		return a
// 	}
// 	low := a[0]
// 	higher := []int{}
// 	if low < a[1] {
// 		a[0], a[1] = a[1], a[0]
// 		highe
// 	}
// 	a = qSort(a) + qSort(low) + qSort(higher)
// 	return a
// }

// func qSort(a []int) []int {
// 	if len(a) < 2 {
// 		return a
// 	}

// }

// 	return nil
// }

// func decompressRLElist(nums []int) []int {
// 	newNums := []int{}
// 	for i := 0; i < len(nums); i++ {
// 		nums[i] +
// 	}
// 	return newNums
// }

func interpret(command string) string {
	command = strings.ReplaceAll(command, "(al)", "al")
	command = strings.ReplaceAll(command, "()", "o")
	return command
}

func numJewelsInStones(jewels string, stones string) int {
	count := 0
	for _, v := range jewels {
		for _, k := range stones {
			if v == k {
				count += 1
			}
		}
	}
	return count
}

func subtractProductAndSum(n int) int {
	strN := strconv.Itoa(n)
	arr := []int{}
	sum, mult, diff := 0, 1, 0

	for _, v := range strN {
		newInt, err := strconv.Atoi(string(v))
		if err != nil {
			fmt.Println(err)
		}
		arr = append(arr, newInt)
	}
	for _, v := range arr {
		sum += v
		mult *= v
	}
	diff = mult - sum
	fmt.Println(sum, mult)
	return diff
}

func maxInt(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if max < v {
			max = v
		}
	}
	return max
}

func sum1(nums []int) int {
	if len(nums) > 1 {
		return nums[0] + sum(nums[1:])
	}
	return nums[0]
}

func sum(nums []int) int {
	total := 0
	if len(nums) > 0 {
		for _, v := range nums {
			total += v
		}
	} else {
		total = 0
	}
	return total
}

func restoreString(s string, indices []int) string {
	newString := make([]rune, len(s))
	for i, v := range s {
		newString[indices[i]] = v
	}
	return string(newString)
}

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

func numberOfSteps(num int) int {
	steps := 0
	for num != 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}
		steps++
	}
	return steps
}

//	func sortSentence(s string) string {
//		idx := []
//		words := strings.Split(s, " ")
//		for _, i := range words {
//			fmt.Println(string(i[len(i)-1]))
//		}
//		return words[0]
//	}
func ReturnErrorPtr() *error {
	var err *error
	return err
}

type MyErr struct{}

func (me MyErr) Error() string {
	return "My error string"
}

func ReturnCustomErr() error {
	var customErr *MyErr
	return customErr
}
func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
	// sortSentence("is2 sentence4 This1 a3")
	// fmt.Println(numberOfSteps(14))
	// m := "My name is Tima"
	// lowerCase := strings.ToLower(m)
	// fmt.Println(lowerCase)
	// fmt.Println(interpret("(al)G(al)()()G"))
	// fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	// fmt.Println(interpret("G()(al)"))
	// fmt.Println(subtractProductAndSum(234))
	// fmt.Println(qSort([]int{12, 10}))
	// fmt.Println(maxInt([]int{20, 30, 50, 60, 40}))
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
