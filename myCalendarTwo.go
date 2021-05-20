//731. My Calendar II  no tripel booking
package main

import "fmt"

type MyCalendarTwo struct {
	bookings [][]int
}

func Constructor() MyCalendarTwo {
	var calendar MyCalendarTwo
	calendar.bookings = [][]int{}
	return calendar

}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	if couldBook(this.bookings, start, end) {
		this.bookings = append(this.bookings, []int{start, end})
		return true
	}

	return false
}

func couldBook(bookings [][]int, start int, end int) bool {
	count := 0
	existingOverLapping := [][]int{}
	for _, value := range bookings {
		//if early or late totally, never booked
		if (value[0] > start && value[0] >= end) || (value[1] <= start && value[1] < end) {
			continue
		} else {
			count += 1
			existingOverLapping = append(existingOverLapping, value)
		}
	}

	for i := 0; i < len(existingOverLapping); i++ {
		for j := i + 1; j < len(existingOverLapping); j++ {
			if hadOverlapping(existingOverLapping[i], existingOverLapping[j]) {
				fmt.Println(existingOverLapping)
				fmt.Println(existingOverLapping[i])
				fmt.Println(existingOverLapping[j])
				return false
			}
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

	//[],[47,50],[1,10],[27,36],[40,47],[20,27],  [15,23],[10,18],[27,36],[17,25],[8,17]
}

/*

["MyCalendarTwo","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book","book"]
[[],[47,50],[1,10],[27,36],[40,47],[20,27],[15,23],[10,18],[27,36],[17,25],[8,17],[24,33],[23,28],[21,27],[47,50],[14,21],[26,32],[16,21],[2,7],[24,33],[6,13],[44,50],[33,39],[30,36],[6,15],[21,27],[49,50],[38,45],[4,12],[46,50],[13,21]]

*/
