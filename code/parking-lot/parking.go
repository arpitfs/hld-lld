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
	fmt.Println(parking.Park(Small))
	fmt.Println(parking.Park(Small))
	fmt.Println(parking.Park(Small))
	fmt.Println(parking.Park(Small))
	fmt.Println(parking.Park(Small))
	fmt.Println(parking.Park(Small))
	fmt.Println(parking.Park(Small))
	fmt.Println(parking.Park(Medium))
	parking.Leave(1, 1)
	parking.DisplaySlots()
}

func (p *Parking) DisplaySlots() {
	for _, levels := range p.levels {
		fmt.Println()
		fmt.Println("Parking Level", levels.levelNumber)
		for _, slots := range levels.slots {
			if slots.isVacant {
				fmt.Printf("%d -> vacant, ", slots.slotNumber)

			} else {
				fmt.Printf(" %d -> not vacant,", slots.slotNumber)

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
