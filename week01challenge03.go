package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DecodeSeatMap(block string) (map[int][]bool, []int) {
	result := make(map[int][]bool)
	invalids := []int{}
	lines := strings.Split(block, "\n")

	// if err != nil {
	// 	return result, invalids
	// }
	for i, val := range lines {
		trimmed := strings.TrimSpace(val)
		row, part, isFound := strings.Cut(trimmed, ":")
		if isFound != true {
			invalids = append(invalids, i+1)
			continue
		}
		rowInt, err := strconv.Atoi(row)
		if err != nil || rowInt < 1 {
			invalids = append(invalids, i+1)
			continue
		}
		decoded, isValid := decoder(part)
		if isValid {
			result[rowInt] = decoded
		} else {
			invalids = append(invalids, i+1)
			continue
		}
	}
	return result, invalids
}
func decoder(part string) ([]bool, bool) {
	result := []bool{}
	// count := 0
	countStr := ""
	// decodeRegexp := regexp.MustCompile('(\d+)([E,O])')
	for _, val := range part {
		if val >= '0' && val <= '9' {
			countStr = countStr + string(val)
		} else if val == 'E' {
			count, err := strconv.Atoi(countStr)
			if err != nil || count < 1 {
				return result, false
			}
			for range count {
				result = append(result, false)
				countStr = ""
			}
		} else if val == 'O' {
			count, err := strconv.Atoi(countStr)
			if err != nil {
				return result, false
			}
			for range count {
				result = append(result, true)
				countStr = ""
			}
		} else {
			return result, false
		}
	}
	return result, true
}

func main() {
	fmt.Println(DecodeSeatMap("12:3E2O1E\nbad line\n7:1O1E2O\n3:0E2O\n123:3E 6O2O\n12:1E4O2E"))
}
