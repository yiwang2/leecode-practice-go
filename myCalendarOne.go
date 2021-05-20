//729. My Calendar I
package main

type MyCalendar struct {
	bookings [][]int
}

func Constructor() MyCalendar {
	var calendar MyCalendar
	calendar.bookings = [][]int{}
	return calendar
}

func (this *MyCalendar) Book(start int, end int) bool {
	if couldBook(this.bookings, start, end) {
		this.bookings = append(this.bookings, []int{start, end})
		return true
	}

	return false
}

func couldBook(bookings [][]int, start int, end int) bool {
	result := true
	for _, value := range bookings {
		//if early or late totally
		if (value[0] > start && value[0] >= end) || (value[1] <= start && value[1] < end) {
			continue
		} else {
			result = false
			break
		}
	}

	return result
}
