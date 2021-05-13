//1292. Maximum Side Length of a Square with Sum Less than or Equal to Threshold
package main

import (
	"fmt"
	"math"
)

func maxSideLength(mat [][]int, threshold int) int {

	minSize := 2
	sizeColumn := len(mat[0])
	sizeRow := len(mat)
	maxSize := int(math.Min(float64(sizeColumn), float64(sizeRow)))

	if maxSize < minSize {
		return 0
	}

	initSize := 1

	for i := minSize; i <= maxSize; i++ {
		if seachSideLength(mat, i, threshold) {
			if i > initSize {
				initSize = i
			}
		}
	}

	if initSize == 1 {
		return 0
	}
	return initSize
}

func seachSideLength(mat [][]int, sizeLimit int, threshold int) bool {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if !checkIfAllValuesSumLessOrEqualThreshold(mat, i, j, sizeLimit, threshold) {
				continue
			} else {
				return true
			}
		}
	}

	return false
}

func checkIfAllValuesSumLessOrEqualThreshold(mat [][]int, startRow int, startColumn int, sizeLimit int, threshold int) bool {
	sizeColumn := len(mat[0])
	sizeRow := len(mat)

	if startRow+sizeLimit > sizeRow || startColumn+sizeLimit > sizeColumn {
		return false
	}

	initValue := 0
	for i := startRow; i < startRow+sizeLimit; i++ {
		for j := startColumn; j < startColumn+sizeLimit; j++ {
			initValue = initValue + mat[i][j]
			if initValue > threshold {
				return false
			}
		}
	}

	return true

}

func main() {
	mat := [][]int{{18, 70}, {61, 1}, {25, 85}, {14, 40}, {11, 96}, {97, 96}, {63, 45}}
	fmt.Println(maxSideLength(mat, 40184))
}
