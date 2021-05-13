//1402. Reducing Dishes

/*
A chef has collected data on the satisfaction level of his n dishes. Chef can cook any dish in 1 unit of time.

Like-time coefficient of a dish is defined as the time taken to cook that dish including previous dishes multiplied by its satisfaction level  i.e.  time[i]*satisfaction[i]

Return the maximum sum of Like-time coefficient that the chef can obtain after dishes preparation.

Dishes can be prepared in any order and the chef can discard some dishes to get this maximum value.

*/

package main

import (
	"fmt"
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	maxNumDishes := len(satisfaction)
	eachSatisfication := 0
	maxSatsification := 0

	for i := maxNumDishes; i > 0; i-- {
		eachSatisfication += satisfaction[i-1]
		if eachSatisfication > 0 {
			maxSatsification += eachSatisfication
		} else {
			break
		}
	}

	return maxSatsification
}

func main() {
	values := []int{-2, 5, -1, 0, 3, -3}
	fmt.Println(maxSatisfaction(values))
}
