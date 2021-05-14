//744. Find Smallest Letter Greater Than Target

package main

import "sort"

func nextGreatestLetter(letters []byte, target byte) byte {

	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})

	smallestLetter := letters[len(letters)-1]

	if letters[0] > target || letters[len(letters)-1] <= target {
		return letters[0]
	}

	for _, letter := range letters {
		if letter > target && letter < smallestLetter {
			smallestLetter = letter
		}
	}

	return smallestLetter
}
