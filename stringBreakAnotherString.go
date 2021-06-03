//1433. Check If a String Can Break Another String
package main

import (
	"fmt"
	"sort"
	"strings"
)

func checkIfCanBreak(s1 string, s2 string) bool {

	capacity := len(s1) // assuming s1 and s2 are same length

	s1Array := strings.Split(s1, "")
	s2Array := strings.Split(s2, "")

	sort.Strings(s1Array)
	sort.Strings(s2Array)

	alwaysSmallerOrEqual := true
	alwaysBiggerOrEqual := true

	for i := 0; i < capacity; i++ {
		if s1Array[i] > s2Array[i] {
			alwaysSmallerOrEqual = false
		}

		if s1Array[i] < s2Array[i] {
			alwaysBiggerOrEqual = false
		}

	}

	return alwaysSmallerOrEqual || alwaysBiggerOrEqual
}

func main() {
	s1 := "leetcodee"
	s2 := "interview"

	fmt.Println(checkIfCanBreak(s1, s2))
}
