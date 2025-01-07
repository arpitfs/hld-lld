package booking

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func bookRoom(scanner *bufio.Scanner) {
	count := 0
	for _, hotel := range hotels {
		if hotel.IsAvailable {
			count++
		}
	}

	if count == 0 {
		fmt.Println("No rooms available in any hotel.")

	} else {
		fmt.Println("Available hotels:")
		for i, hotel := range hotels {
			for _, room := range hotel.Rooms {
				if room.isVacant {
					fmt.Printf("%d. %s - Room %d\n", i+1, hotel.Name, room.Number)
				}
			}
		}
	}

	fmt.Print("Enter the hotel ")
	scanner.Scan()
	hotelName := strings.TrimSpace(scanner.Text())
	fmt.Print("Enter the hotel number ")
	scanner.Scan()
	hotelNumber, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	for _, hotel := range hotels {
		if hotel.Name == hotelName {
			for i := range hotel.Rooms {
				if hotel.Rooms[i].Number == hotelNumber {
					hotel.Rooms[i].isVacant = false
					fmt.Printf("Room %d in %s has been booked.\n\n", hotelNumber, hotelName)
				}
			}
		}
	}

	for _, hotel := range hotels {
		count := 0
		for _, rooms := range hotel.Rooms {
			if rooms.isVacant {
				count++
			}
		}

		if count == 0 {
			hotel.IsAvailable = false
		}
	}
}

func viewBookings() {
	for _, hotel := range hotels {
		fmt.Printf("Hotel: %s\n", hotel.Name)
		for _, room := range hotel.Rooms {
			if room.isVacant {
				fmt.Printf("Room %d is vacant.\n", room.Number)
			} else {
				fmt.Printf("Room %d is booked.\n", room.Number)
			}
		}
	}
}

func Print(data string) {
	fmt.Println(data)
}
