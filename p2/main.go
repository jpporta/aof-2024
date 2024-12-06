package p2

import (
	"log"
	"strconv"
	"strings"
)

func getLines(input string) [][]int {
	var list [][]int
	rows := strings.Split(input, "\n")
	for _, row := range rows {

		numsStr := strings.Split(row, " ")
		var nums []int
		if len(numsStr) <= 1 {
			continue
		}
		for _, val := range numsStr {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatalf("num failed to be converted")
			}
			nums = append(nums, num)
		}
		list = append(list, nums)
	}
	return list
}
func absInt(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
func First(input string) int {
	safe := 0
	list := getLines(input)
	for _, line := range list {
		safeLine := true
		inc := true
		prev := -1
		for ind, num := range line {
			if prev == -1 {
				prev = num
				continue
			}
			diff := absInt(prev, num)
			if diff < 1 || diff > 3 {
				safeLine = false
				break
			}
			if ind == 1 {
				inc = prev < num
			}
			if (inc && num < prev) || (!inc && num > prev) {
				safeLine = false
				break
			}
			prev = num
		}
		if safeLine {
			safe++
		}
	}
	return safe
}
func checkRange(i int) bool { return i < 1 || i > 3 }

func checkLine(line []int, skip int) bool {
	inc := 0
	prev := -1
	for ind, num := range line {
		// If skip continue
		if ind == skip {
			continue
		}
		// If no prev / first element continue
		if prev == -1 {
			prev = num
			continue
		}
		// check if the abs diff is outside safe range
		diff := absInt(prev, num)
		if checkRange(diff) {
			if skip != -1000 {
				return false
			}
			return checkLine(line, ind-2) || checkLine(line, ind-1) || checkLine(line, ind) || checkLine(line, ind+1)
		}
		// Check if it is increasing or decreasing
		if inc == 0 {
			if prev < num {
				inc = 1
			} else {
				inc = -1
			}
		}
		// Has it changed from prior
		if (inc == 1 && num < prev) || (inc == -1 && num > prev) {
			if skip != -1000 {
				return false
			}
			return checkLine(line, ind-2) || checkLine(line, ind-1) || checkLine(line, ind) || checkLine(line, ind+1)
		}
		prev = num
	}
	return true
}

func Second(input string) int {
	safe := 0
	list := getLines(input)
	for _, line := range list {
		if checkLine(line, -1000) {
			safe++
		}
	}
	return safe
}
