//38. Count and Say
package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	return cs(n)
}

func cs(n int) string {

	result := "1"

	for i := 1; i < n; i++ {
		result = countNum(result)
	}

	return result
}

type statisticNode struct {
	value byte
	count int
}

func countNum(input string) string {
	value := input[0]
	count := 1
	statistics := []statisticNode{}

	for i := 1; i < len(input); i++ {
		if value == input[i] {
			count += 1
		} else {
			statistics = append(statistics, statisticNode{value, count})
			count = 1
			value = input[i]
		}
	}

	statistics = append(statistics, statisticNode{value, count})

	result := ""
	for i := 0; i < len(statistics); i++ {
		result = result + strconv.FormatInt(int64(statistics[i].count), 10) + string(statistics[i].value)
	}

	return result
}

func main() {
	fmt.Println(countAndSay(10))
}
