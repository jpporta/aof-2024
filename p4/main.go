package p4

import (
	"strings"
)

func First(input string) int {
	count := 0
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	for y := range lines {
		for x := range lines[y] {
			if len(lines[y]) == 0 {
				continue
			}
			if lines[y][x] != 'X' && lines[y][x] != 'S' {
				continue
			}
			// are there enough lines below - vertical
			if y <= len(lines)-4 {
				if lines[y][x] == 'X' && lines[y+1][x] == 'M' && lines[y+2][x] == 'A' && lines[y+3][x] == 'S' {
					count++
				} else if lines[y][x] == 'S' && lines[y+1][x] == 'A' && lines[y+2][x] == 'M' && lines[y+3][x] == 'X' {
					count++
				}
			}
			// are there enough lines to the right - horizontal
			if x <= len(lines[y])-4 {
				if lines[y][x] == 'X' && lines[y][x+1] == 'M' && lines[y][x+2] == 'A' && lines[y][x+3] == 'S' {
					count++
				} else if lines[y][x] == 'S' && lines[y][x+1] == 'A' && lines[y][x+2] == 'M' && lines[y][x+3] == 'X' {
					count++
				}
			}

			// both checks - diagonal bellow
			if y <= len(lines)-4 && x <= len(lines[y])-4 {
				if lines[y][x] == 'X' && lines[y+1][x+1] == 'M' && lines[y+2][x+2] == 'A' && lines[y+3][x+3] == 'S' {
					count++
				} else if lines[y][x] == 'S' && lines[y+1][x+1] == 'A' && lines[y+2][x+2] == 'M' && lines[y+3][x+3] == 'X' {
					count++
				}
			}
			// both checks - diagonal above
			if y >= 3 && x <= len(lines[y])-4 {
				if lines[y][x] == 'X' && lines[y-1][x+1] == 'M' && lines[y-2][x+2] == 'A' && lines[y-3][x+3] == 'S' {
					count++
				} else if lines[y][x] == 'S' && lines[y-1][x+1] == 'A' && lines[y-2][x+2] == 'M' && lines[y-3][x+3] == 'X' {
					count++
				}
			}

		}
	}
	return count
}

// func Second(input string) int {
// }
