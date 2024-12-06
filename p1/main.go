package p1

import (
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func splitListsSorted(input string) ([]int, []int) {
	var list1, list2 []int

	rows := strings.Split(input, "\n")
	for _, row := range rows {
		nums := strings.Split(row, " ")
		if len(nums) < 4 {
			continue
		}
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalf("num1 failed to be converted")
		}
		num2, err := strconv.Atoi(nums[len(nums)-1])
		if err != nil {
			log.Fatalf("num2 failed to be converted")
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	return list1, list2
}

func First(input string) int {
	list1, list2 := splitListsSorted(input)
	diff := 0
	for i := range list1 {
		diff += int(math.Abs(
			float64(list1[i] - list2[i]),
		))
	}
	return diff
}

func Second(input string) int {
	list1, list2 := splitListsSorted(input)
	log.Println(list1, list2)
	p2 := 0
	sim := 0
	encounters := 0
	target := 0

	for _, val := range list1 {
		if target != val {
			encounters = 0
		}
		if p2 >= len(list2) {
			break
		}
		for p2 < len(list2) && val >= list2[p2] {
			if val == list2[p2] {
				target = val
				encounters += 1
			}
			p2++
		}
		// Val == list2[p2]
		if target == val {
			sim += val * encounters
		}
		log.Println(sim)
	}
	return sim
}
