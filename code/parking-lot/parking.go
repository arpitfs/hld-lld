package parkinglot

import (
	"fmt"
)

const (
	TotalSlots  = 8
	TotalLevels = 2
)

func StartParking() {
	parking := NewParking()
	fmt.Println(parking.Park(VehicleTypeSmall))
	fmt.Println(parking.Park(VehicleTypeSmall))
	fmt.Println(parking.Park(VehicleTypeSmall))
	fmt.Println(parking.Park(VehicleTypeLarge))
	fmt.Println(parking.Park(VehicleTypeSmall))
	fmt.Println(parking.Park(VehicleTypeSmall))
	fmt.Println(parking.Park(VehicleTypeSmall))
	fmt.Println(parking.Park(VehicleTypeMedium))
	parking.Leave(1, 1)
	parking.DisplaySlots()
}

func (p *Parking) DisplaySlots() {
	for _, level := range p.levels {
		fmt.Println()
		fmt.Println("Parking Level", level.levelNumber)
		for _, slot := range level.slots {
			if slot.isVacant {
				fmt.Printf("%d -> vacant, ", slot.slotNumber)
			} else {
				fmt.Printf(" %d -> not vacant,", slot.slotNumber)
			}
		}
	}
}

func (p *Parking) Park(vehicleType VehicleType) (string, string) {
	for _, level := range p.levels {
		for i := range level.slots {
			slot := &level.slots[i]
			if slot.isVacant && vehicleType == slot.slotType {
				slot.isVacant = false
				return fmt.Sprintf("Park at Level %d, Slot %d", level.levelNumber, slot.slotNumber), ""
			}
		}
	}
	return "", "No vacant space"
}

func (p *Parking) Leave(levelNumber, slotNumber int) {
	level := &p.levels[levelNumber-1]
	slot := &level.slots[slotNumber-1]
	slot.isVacant = true
}
