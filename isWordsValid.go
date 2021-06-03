//1003. Check If Word Is Valid After Substitutions

package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {

	sArray := strings.Split(s, "")
	sPartten := ""
	distinct := make(map[string]bool)

	for _, value := range sArray {
		if _, found := distinct[value]; !found {
			distinct[value] = true
			sPartten += value
		}
	}

	if sPartten != "abc" {
		return false
	}

	fmt.Println(sPartten)
	cannotFound := true
	for len(s) != 0 {
		index := strings.Index(s, sPartten)
		//fmt.Println(index)
		if index < 0 {
			break
		} else if index == 0 {
			s = s[len(sPartten):]
		} else {
			//fmt.Println(s[0:index])
			//fmt.Println(s[index+len(sPartten):])
			s = s[0:index] + s[index+len(sPartten):]
		}
	}

	if len(s) == 0 {
		cannotFound = false
	}

	return !cannotFound
}

func main() {
	s := "cababc"
	fmt.Println(isValid(s))
}
