package booking

var hotels []Hotel

type Hotel struct {
	Name        string
	Rooms       []Room
	IsAvailable bool
}

func SetUpSystem() []Hotel {
	tajHotel := Hotel{
		Name: "Taj",
		Rooms: []Room{
			{Number: 101, isVacant: true},
			{Number: 102, isVacant: true},
			{Number: 103, isVacant: true},
		},
		IsAvailable: true,
	}
	marriotHotel := Hotel{
		Name: "Marriot",
		Rooms: []Room{
			{Number: 101, isVacant: true},
			{Number: 102, isVacant: true},
			{Number: 103, isVacant: true},
		},
		IsAvailable: true,
	}
	hotels = []Hotel{tajHotel, marriotHotel}
	return hotels
}
