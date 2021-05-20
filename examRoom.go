//855. Exam Room
package main

import (
	"fmt"
	"math"
	"sort"
)

type ExamRoom struct {
	seats       map[int]int
	maxCapacity int
}

func Constructor(n int) ExamRoom {
	var room ExamRoom
	room.seats = make(map[int]int) //seats[position] = 1
	room.maxCapacity = n

	return room
}

func (this *ExamRoom) Seat() int {
	return findSeat(this.seats, this.maxCapacity)
}

func (this *ExamRoom) Leave(p int) {
	_, ok := this.seats[p]
	if ok {
		delete(this.seats, p)
	}
}

func findSeat(seats map[int]int, maxCapacity int) int {
	position := 0
	currentCollection := collectCurrentSeizedInfo(seats)
	//fmt.Println(seats)
	//fmt.Println(currentCollection)
	if len(currentCollection) == 0 {
		seats[0] = 1
		return 0
	} else if len(currentCollection) == 1 {
		position = currentCollection[0]

		// in the middle
		if position != 0 && position != maxCapacity-1 {
			position = (maxCapacity - 1) / 2
			if position > maxCapacity-1-position {
				position = 0
			} else {
				position = maxCapacity - 1
			}
		} else if position == 0 {
			position = maxCapacity - 1
		} else {
			position = 0
		}
		seats[position] = 1
		return position
	} else {
		includingLeftWall := currentCollection[0] != 0
		includingRightWall := currentCollection[len(currentCollection)-1] != maxCapacity-1
		distanceRecord := make(map[int]int) // postion and max distance with neighbors

		for index, seizedIndex := range currentCollection {
			if index == 0 {
				if includingLeftWall {
					distanceRecord[0] = seizedIndex
				} else {
					tempPosition := (currentCollection[index+1]-currentCollection[index])/2 + currentCollection[index]
					distanceRecord[tempPosition] = int(math.Min(float64(tempPosition-currentCollection[index]), float64(currentCollection[index+1]-tempPosition)))
				}
				continue
			}
			if index == len(currentCollection)-1 {
				if includingRightWall {
					distanceRecord[maxCapacity-1] = maxCapacity - 1 - seizedIndex
				} else {
					tempPosition := (currentCollection[index]-currentCollection[index-1])/2 + currentCollection[index-1]
					distanceRecord[tempPosition] = int(math.Min(float64(tempPosition-currentCollection[index-1]), float64(currentCollection[index]-tempPosition)))
				}
				continue
			}

			tempPosition := (currentCollection[index+1]-currentCollection[index])/2 + currentCollection[index]
			distanceRecord[tempPosition] = int(math.Min(float64(tempPosition-currentCollection[index]), float64(currentCollection[index+1]-tempPosition)))

		}

		//min key with max value
		fmt.Println(distanceRecord)
		minKey := maxCapacity
		var maxValue int

		for _, v := range distanceRecord {
			if v > maxValue {
				maxValue = v
			}
		}

		for k, v := range distanceRecord {
			if v == maxValue && k < minKey {
				minKey = k
			}
		}
		position = minKey
		seats[position] = 1

		return position
	}
}

func collectCurrentSeizedInfo(seats map[int]int) []int {
	result := make([]int, 0, len(seats))

	for index, _ := range seats {
		result = append(result, index)
	}
	sort.Ints(result)
	return result
}

func (this *ExamRoom) DisplaySeat() {
	fmt.Println(this.seats)
}

func main() {
	room := Constructor(8)
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	room.Leave(0)
	room.Leave(7)
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
	fmt.Println(room.Seat())
}
