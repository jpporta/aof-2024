package p3

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func First(input string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAll([]byte(input), -1)
	sum := 0
	for _, mul := range matches {
		nums := strings.Split(string(mul[4:len(mul)-1]), ",")
		n1, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalf("Failed to convert first number")
		}
		n2, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatalf("Failed to convert second number")
		}
		sum += n1 * n2
	}
	return sum
}

func Second(input string) int {
	commands := strings.Split(input, "don't()")
	valid_commands := commands[0]
	re_do := regexp.MustCompile(`do()`)
	for _, section := range commands[1:] {
		do_idx := re_do.FindIndex([]byte(section))
		if do_idx == nil {
			continue
		}
		valid_commands += section[do_idx[1]:]
	}

	return First(valid_commands)
}
