//1406. Stone Game III
package main

import (
	"fmt"
	"math"
)

var ALICE string = "Alice"
var BOB string = "Bob"
var TIE string = "Tie"

// could alice win, no, then could alice tie, then alice lost
func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	dp := make([]int, n+1) // means if I pick stones from  i, The max value I may have
	dp[n-1] = stoneValue[n-1]
	sum := stoneValue[n-1]

	for i := n - 2; i >= 0; i-- {
		sum += stoneValue[i]
		min := math.MaxInt64
		for j := i + 1; j <= i+3 && j <= n; j++ {
			if min > dp[j] {
				min = dp[j]
			}
			dp[i] = sum - min
		}
	}

	fmt.Println(sum)
	fmt.Println(dp)

	//dp[0] means Alick picked up from 0 and the max value Alice could have, then sum - dp[0] means if Alice pick from 0, the max value bob could have
	if dp[0] > sum-dp[0] {
		return ALICE
	} else if dp[0] < sum-dp[0] {
		return BOB
	} else {
		return TIE
	}
}

func main() {
	values := []int{1, 2, 3, -9}

	fmt.Println(stoneGameIII(values))
}
