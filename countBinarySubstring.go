//696. Count Binary Substrings

package main

import (
	"fmt"
	"math"
	"strings"
)

func countBinarySubstrings(s string) int {
	ans := 0
	prev := 0
	cur := 1

	sLength := len(s)
	splitsValues := strings.Split(s, "")

	for i := 1; i < sLength; i++ {
		if splitsValues[i-1] != splitsValues[i] {
			//this means, if we found sth like 00001
			//then pre = cur = 4, cur =1
			//then if we found st like 0000111'0', pre =4, cur = 3. then ans = ans +3
			ans += int(math.Min(float64(prev), float64(cur)))
			prev = cur
			cur = 1
		} else {
			cur++
		}
	}
	//this is the last 1
	return ans + int(math.Min(float64(prev), float64(cur)))
}

func countBinarySubstringsSlow(s string) int {

	sLength := len(s)
	result := 0
	if sLength == 0 || sLength == 1 {
		return result
	}

	splitsValues := strings.Split(s, "")
	firstvalue := splitsValues[0]
	secondValue := ""
	count := 1

	for i := 1; i < sLength; i++ {
		//initally, we expect the values are equal to 1st value and count ++
		if splitsValues[i] == firstvalue && secondValue == "" {
			count += 1
			continue
		}

		if splitsValues[i] != firstvalue && secondValue == "" {
			secondValue = splitsValues[i]
			count -= 1
			if count == 0 {
				break
			}
			continue
		}

		if secondValue != "" && splitsValues[i] == secondValue {
			count -= 1
			if count == 0 {
				break
			}
			continue
		}

		if count == 0 || (count > 0 && secondValue != "" && splitsValues[i] != secondValue) {
			break
		}
	}

	if count == 0 {
		result += 1
	}

	return result + countBinarySubstringsSlow(s[1:sLength])
}

func main() {
	s := "10101"
	fmt.Println(countBinarySubstringsSlow(s))
}
