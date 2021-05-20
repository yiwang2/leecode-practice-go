//731. My Calendar II  no tripel booking
package main

import (
	"fmt"
	"math"
)

type MyCalendarTwo struct {
	bookings            [][]int
	existingOverlapping [][]int
}

func Constructor() MyCalendarTwo {
	var calendar MyCalendarTwo
	calendar.bookings = [][]int{}
	calendar.existingOverlapping = [][]int{}
	return calendar

}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	if couldBook(this.existingOverlapping, start, end) {
		for _, value := range this.bookings {
			testOverLapping := hadOverlapping(value, []int{start, end})
			if testOverLapping {
				this.existingOverlapping = append(this.existingOverlapping, []int{int(math.Max(float64(value[0]), float64(start))), int(math.Min(float64(value[1]), float64(end)))})
			}
		}

		this.bookings = append(this.bookings, []int{start, end})
		return true
	}

	return false
}

func couldBook(existingOverlapping [][]int, start int, end int) bool {

	for _, value := range existingOverlapping {
		//if early or late totally, never booked
		testOverLapping := hadOverlapping(value, []int{start, end})
		if testOverLapping {
			return false
		}
	}

	return true
}

func hadOverlapping(a []int, b []int) bool {
	return !((a[0] > b[0] && a[0] >= b[1]) || (a[1] <= b[0] && a[1] < b[1]))
}

func main() {
	var two MyCalendarTwo = Constructor()
	two.Book(47, 50)
	two.Book(1, 10)
	two.Book(27, 36)
	two.Book(40, 47)
	two.Book(20, 27)
	two.Book(15, 23)
	two.Book(10, 18)
	two.Book(27, 36)
	two.Book(17, 25)
	fmt.Println("--------------")
	fmt.Println(two.Book(8, 17))
}
