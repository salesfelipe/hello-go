// Package: hello-package
package main

import (
	"fmt"
	"hello-go/booking"
)

// Main do stuff
func main() {

	// fmt.Println(booking.Schedule("7/25/2019 13:45:00"))

	fmt.Println(booking.IsAfternoonAppointment("Thursday, July 25, 2019 10:45:00"))
	fmt.Println(booking.Description("7/25/2019 13:45:00"))
	fmt.Println(booking.AnniversaryDate())
}
