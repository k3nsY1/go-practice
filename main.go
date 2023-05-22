package main

import (
	"errors"
	"fmt"
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

func ReturnErrorPtr() *error {
	var err *error
	return err
}

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() function!")
		return err
	} else {
		return nil
	}
}
func main() {
	err := returnError(1, 2)
	if err == nil {
		fmt.Println("returnError() ended normally")
	} else {
		fmt.Println(err)
	}

	err = returnError(10, 10)
	if err == nil {
		fmt.Println("returnError() ended normally!")
	} else {
		fmt.Println(err)
	}

	if err.Error() == "Error in returnError() function!" {
		fmt.Println("!!")
	}
}
