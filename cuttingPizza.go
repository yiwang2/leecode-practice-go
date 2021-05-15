//1444. Number of Ways of Cutting a Pizza
/*
Given a rectangular pizza represented as a rows x cols matrix containing the following characters: 'A' (an apple) and '.' (empty cell) and given the integer k. You have to cut the pizza into k pieces using k-1 cuts.

For each cut you choose the direction: vertical or horizontal, then you choose a cut position at the cell boundary and cut the pizza into two pieces. If you cut the pizza vertically, give the left part of the pizza to a person. If you cut the pizza horizontally, give the upper part of the pizza to a person. Give the last piece of pizza to the last person.

Return the number of ways of cutting the pizza such that each piece contains at least one apple. Since the answer can be a huge number, return this modulo 10^9 + 7.

*/
package main

import (
	"fmt"
	"strings"
)

var cache [][][]int
var kMode int = 1e9 + 7

func ways(pizza []string, k int) int {
	matrix := createdMatrx(pizza)
	row := len(pizza)
	column := len(pizza[0])
	createCache(row, column, k)
	return dp(matrix, 0, 0, k-1)

}

func createCache(row int, column int, k int) {
	cache = make([][][]int, row)
	for i, _ := range cache {
		cache[i] = make([][]int, column)
		for j, _ := range cache[i] {
			cache[i][j] = make([]int, k)
			for z, _ := range cache[i][j] {
				cache[i][j][z] = -1
			}
		}
	}
}

func createdMatrx(pizza []string) [][]int {
	row := len(pizza)
	column := len(pizza[0])

	result := make([][]int, row)

	for index, rowString := range pizza {
		result[index] = make([]int, column)
		rowStringSplits := strings.Split(rowString, "")
		for i, value := range rowStringSplits {
			if value == "A" {
				result[index][i] = 1
			} else {
				result[index][i] = 0
			}

		}

	}

	return result
}

func dp(matrix [][]int, m int, n int, k int) int {
	row := len(matrix)
	column := len(matrix[0])

	if k == 0 {
		return hasApple(matrix, m, n, row-1, column-1)
	}

	ans := cache[m][n][k]
	if ans != -1 {
		return ans
	}

	ans = 0
	for currentRow := m; currentRow < row-1; currentRow++ {
		ans = (ans + hasApple(matrix, m, n, currentRow, column-1)*dp(matrix, currentRow+1, n, k-1)) % kMode
	}

	for currentColumn := n; currentColumn < column-1; currentColumn++ {
		ans = (ans + hasApple(matrix, m, n, row-1, currentColumn)*dp(matrix, m, currentColumn+1, k-1)) % kMode
	}
	cache[m][n][k] = ans
	return ans
}

func hasApple(matrix [][]int, x1 int, y1 int, x2 int, y2 int) int {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if matrix[i][j] == 1 {
				return 1
			}
		}
	}

	return 0
}

func main() {
	pizza := []string{"A..", "AA.", "..."}
	fmt.Println(ways(pizza, 3))
}
