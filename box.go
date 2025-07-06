package chaisay

import (
	"fmt"
	"strings"
)

const (
	corner         string = "+"
	horizontalLine string = "-"
	verticalLine   string = "|"
)

func PrintTextBox(s string) {
	newStr := strings.Split(s, " ")
	finalStr := []string{}
	str := ""
	for _, val := range newStr {
		if len(str)+len(val) >= 50 {
			finalStr = append(finalStr, str)
			str = ""
		}
		str = str + " " + val
	}
	if len(str) > 0 {
		finalStr = append(finalStr, str)
		str = ""
	}
	fmt.Println(corner + strings.Repeat(horizontalLine, 2) + strings.Repeat(horizontalLine, maximumLength(finalStr)) + strings.Repeat(horizontalLine, 2) + corner)
	for _, val := range finalStr {
		fmt.Printf("%v  %v %v %v\n", verticalLine, val, strings.Repeat(" ", maximumLength(finalStr)-len(val)), verticalLine)
	}
	fmt.Println(corner + strings.Repeat(horizontalLine, 2) + strings.Repeat(horizontalLine, maximumLength(finalStr)) + strings.Repeat(horizontalLine, 2) + corner)
}

func maximumLength(arr []string) int {
	max := 0
	for _, val := range arr {
		if len(val) > max {
			max = len(val)
		}
	}
	return max
}
