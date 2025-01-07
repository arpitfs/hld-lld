package booking

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Booking() {
	SetUpSystem()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Book a room")
		fmt.Println("2. View Bookings")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			bookRoom(scanner)
		case "2":
			viewBookings()
		case "3":
			Print("Exiting the application.")
			return
		default:
			Print("Please try with valid choice.\n")
		}
	}
}
