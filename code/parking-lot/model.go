package parkinglot

const (
	VehicleTypeSmall VehicleType = iota
	VehicleTypeMedium
	VehicleTypeLarge
)

type VehicleType int

type Vehicle struct {
	vehicleNumber string
	vehicleType   VehicleType
}

type Level struct {
	levelNumber int
	slots       []Slot
}

type Slot struct {
	slotNumber int
	isVacant   bool
	slotType   VehicleType
}

type Parking struct {
	levels []Level
}

func NewSlots() []Slot {
	return []Slot{
		{slotNumber: 1, isVacant: true, slotType: VehicleTypeSmall},
		{slotNumber: 2, isVacant: true, slotType: VehicleTypeMedium},
		{slotNumber: 3, isVacant: true, slotType: VehicleTypeLarge},
		{slotNumber: 4, isVacant: true, slotType: VehicleTypeSmall},
		{slotNumber: 5, isVacant: true, slotType: VehicleTypeMedium},
		{slotNumber: 6, isVacant: true, slotType: VehicleTypeLarge},
		{slotNumber: 7, isVacant: true, slotType: VehicleTypeSmall},
		{slotNumber: 8, isVacant: true, slotType: VehicleTypeMedium},
	}
}

func NewLevels(slots1, slots2 []Slot) []Level {
	return []Level{
		{levelNumber: 1, slots: slots1}, {levelNumber: 2, slots: slots2},
	}
}

func NewVehicle(vehicleNumber string, vehicleType VehicleType) Vehicle {
	return Vehicle{
		vehicleNumber: vehicleNumber,
		vehicleType:   vehicleType,
	}
}

func NewParking() *Parking {
	slots1 := NewSlots()
	slots2 := NewSlots()
	levels := NewLevels(slots1, slots2)

	parking := &Parking{levels: levels}
	return parking
}
