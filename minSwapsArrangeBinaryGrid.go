//1536. Minimum Swaps to Arrange a Binary Grid
/*Given an n x n binary grid, in one step you can choose two adjacent rows of the grid and swap them.

A grid is said to be valid if all the cells above the main diagonal are zeros.

Return the minimum number of steps needed to make the grid valid, or -1 if the grid cannot be valid.

The main diagonal of a grid is the diagonal that starts at cell (1, 1) and ends at cell (n, n).*/
package main

import "fmt"

func minSwaps(grid [][]int) int {
	gridLength := len(grid)
	zerosForEachRow := make([]int, gridLength)

	for i := 0; i < gridLength; i++ {
		zerosForEachRow[i] = numOfTrailingZeros(grid, i)
	}

	result := 0
	for i := 0; i < gridLength; i++ {
		if zerosForEachRow[i] < gridLength-i-1 {
			find := false
			for j := i + 1; j < gridLength; j++ {
				if zerosForEachRow[j] >= gridLength-i-1 {
					for zerosForEachRow[i] < gridLength-i-1 {
						zerosForEachRow[j-1], zerosForEachRow[j] = zerosForEachRow[j], zerosForEachRow[j-1]
						result += 1
						j -= 1
					}

					find = true
				}
			}

			if find == false {
				result = -1
				break
			}
		}
	}

	return result
}

// num of trailing zeros at row
func numOfTrailingZeros(grid [][]int, rowIndex int) int {
	count := 0
	for j := 0; j < len(grid[rowIndex]); j++ {
		if grid[rowIndex][j] == 0 {
			count += 1
		} else {
			count = 0
		}
	}

	return count
}

func main() {
	grid := [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}}
	fmt.Println(minSwaps(grid))

}
