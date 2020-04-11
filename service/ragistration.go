package service

import (
	"errors"
	"fmt"
)

// Ragistration car parking muster
type Ragistration struct {
	Slot           int    `json:"slot"`
	RagistrationNo string `json:"carNumber"`
	Colour         string `json:"carColour"`
	isNonEmpty     bool
}

var (
	ragistrationMuster        = []Ragistration{}
	carNumbersByColour        = make(map[string][]Ragistration, 0)
	slotsByRagistrationNumber = make(map[string]int, 0)
	// RagistrationMuster allow in once
	RagistrationMuster *Ragistration
)

// InitialiseParkingMuster initialize parking muster
func init() {
	RagistrationMuster = &Ragistration{}
}

// CreateParkingLot allot Parking slot
func (r *Ragistration) CreateParkingLot(slot int) (string, error) {
	if len(ragistrationMuster) > 0 {
		return "", errors.New("parking slot already define")
	}
	if slot <= 0 {
		return "", errors.New("plase enter valid slot number")
	}
	for i := 1; i <= slot; i++ {
		ragistrationMuster = append(ragistrationMuster, Ragistration{Slot: i})
	}
	return fmt.Sprintf("Created a parking lot with %v slots", slot), nil
}

// ParkCar park the car on empty slot
func (r *Ragistration) ParkCar(ragistrationNumber, colour string) (string, error) {
	if len(ragistrationNumber) == 0 || len(colour) == 0 {
		return "", errors.New("please provide valid car details")
	}
	flag := false
	for index, val := range ragistrationMuster {
		if !val.isNonEmpty {
			flag = true
			addCarInRagister(r, ragistrationNumber, colour, index)
			return fmt.Sprintf("Allocated slot number: %v", r.Slot), nil
		}
	}
	if !flag {
		return "", errors.New("Sorry, parking lot is full")
	}
	return "", nil
}

// GetParkingLotStatus get parking slot status
func (r *Ragistration) GetParkingLotStatus() ([]Ragistration, error) {
	if len(ragistrationMuster) == 0 {
		return []Ragistration{}, errors.New("parking slot empty")
	}
	return ragistrationMuster, nil
}

// LeaveCar leave the car on given slot
func (r *Ragistration) LeaveCar(slot int) (string, error) {
	if slot <= 0 {
		return "", errors.New("plase enter valid slot number")
	}
	r.RagistrationNo = ""
	r.Colour = ""
	r.Slot = slot
	r.isNonEmpty = false
	// remove element in ragister
	carNumbers := carNumbersByColour[ragistrationMuster[slot-1].Colour]
	i := removeElement(ragistrationMuster[slot-1].RagistrationNo, carNumbers)
	if i == -1 {
		return "", errors.New("slot already available")
	}
	carNumbers = append(carNumbers[:i], carNumbers[i+1:]...)
	carNumbersByColour[ragistrationMuster[slot-1].Colour] = carNumbers

	// remove from slot information
	delete(slotsByRagistrationNumber, ragistrationMuster[slot-1].RagistrationNo)

	// remove slot
	ragistrationMuster[slot-1] = *r
	return fmt.Sprintf("Slot number %v is free", slot), nil
}

func addCarInRagister(r *Ragistration, ragistrationNo, colour string, index int) {
	r.RagistrationNo = ragistrationNo
	r.Colour = colour
	r.Slot = index + 1
	r.isNonEmpty = true
	carNumbersByColour[r.Colour] = append(carNumbersByColour[r.Colour], *r)
	slotsByRagistrationNumber[r.RagistrationNo] = r.Slot
	ragistrationMuster[index] = *r
}

func removeElement(number string, numbers []Ragistration) int {
	for i, v := range numbers {
		if v.RagistrationNo == number {
			return i
		}
	}
	return -1
}
