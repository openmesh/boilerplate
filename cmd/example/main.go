package main

import (
	"github.com/openmesh/boilerplate"
)

func main() {
	// service := boilerplate.Service{
	// 	Name:        "Booking",
	// 	Methods:     []string{"FindBookingByID", "FindBookings", "CreateBooking", "UpdateBooking", "DeleteBooking"},
	// 	PackageName: "booking",
	// }
	service := boilerplate.Service{
		Name:        "Unavailability",
		Methods:     []string{"FindUnavailabilityByID", "FindUnavailabilities", "CreateUnavailability", "UpdateUnavailability", "DeleteUnavailability"},
		PackageName: "booking",
	}

	boilerplate.Generate(service)
}
