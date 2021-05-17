//1687. Delivering Boxes from Storage to Ports
package main

import (
	"fmt"
	"math"
)

//boxes[i] = [ports​​​, weight]
func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	boxLength := len(boxes)
	boxes = append([][]int{{-1, 0}}, boxes...)
	dp := make([]int, boxLength+1)
	//init value and set as max
	for index, _ := range dp {
		if index == 0 {
			dp[0] = 0
		} else {
			dp[index] = int(math.Pow10(6))
		}
	}

	j := 0 // current dp[j]'s trip
	weightSum := 0
	tripSum := 0
	lastPort := 0
	lastj := 0

	for i := 1; i <= boxLength; i++ {

		for j+1 <= boxLength && weightSum+boxes[j+1][1] <= maxWeight && j+1-i+1 <= maxBoxes {
			j += 1
			weightSum += boxes[j][1]
			//if current port is different with last port, then trip + 1
			if boxes[j][0] != boxes[j-1][0] {
				tripSum += 1
			}

			//remember the last different port
			if boxes[j][0] != lastPort {
				lastPort = boxes[j][0]
				lastj = j
			}
		}

		//let's do the greedy, then improve it
		//box[j+1] = box[i-1] + trips(box[i:j]) + 1
		expectedTrip := dp[i-1] + tripSum + 1
		if expectedTrip <= dp[j] {
			dp[j] = expectedTrip
		}
		//if 2 ports are same
		if j+1 <= boxLength && boxes[j][0] == boxes[j+1][0] {
			expectedTrip1 := dp[i-1] + tripSum //-1 then +1
			if expectedTrip1 <= dp[lastj-1] {
				dp[lastj-1] = expectedTrip1
			}
		}

		//let pointer moving forward, then we shall remove it for next round
		weightSum -= boxes[i][1]
		if i+1 <= boxLength && boxes[i][0] != boxes[i+1][0] {
			tripSum -= 1
		}

	}

	return dp[boxLength]

}

func boxDelivering2(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {

	// dp[i] := min trips to deliver boxes[0..i) and return to the storage
	boxLength := len(boxes)
	dp := make([]int, boxLength+1)
	trips := 2
	weight := 0

	l := 0 // last index, r current index
	for r := 0; r < boxLength; r++ {
		weight += boxes[r][1]

		// current box is different from previous one, need to make one more trip
		//yes no matter if we could have it or not
		if r > 0 && boxes[r][0] != boxes[r-1][0] {
			trips += 1
		}

		//[i.......j]
		//d[r+1] = d[l] + trips(d[l+1:r]) + 1

		//if over load or over size,
		for r-l+1 > maxBoxes || weight > maxWeight ||
			// loading boxes[l] in the previous turn is always no bad than
			// loading it in this turn
			(l < r && dp[l+1] == dp[l]) {
			weight -= boxes[l][1]
			if boxes[l][0] != boxes[l+1][0] {
				trips -= 1
			}

			l += 1
		}

		//   min trips to deliver boxes[0..r]
		// = min trips to deliver boxes[0..l) + trips to deliver boxes[l..r]
		dp[r+1] = dp[l] + trips
	}

	return dp[boxLength]
}

func main() {
	boxes := [][]int{{2, 4}, {2, 5}, {3, 1}, {3, 2}, {3, 7}, {3, 1}, {4, 4}, {1, 3}, {5, 2}}
	portsCount := 5
	maxBoxes := 5
	maxWeight := 7

	fmt.Println(boxDelivering(boxes, portsCount, maxBoxes, maxWeight))

}
