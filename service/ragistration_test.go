package service

import (
	"fmt"
	"testing"
)

var ragistrationService = Ragistration{}

func TestRagistration_CreateParkingLot(t *testing.T) {
	slot := 6
	got, err := ragistrationService.CreateParkingLot(slot)
	if err != nil {
		t.Errorf("TestRagistration_CreateParkingLot() %v", err.Error())
	} else {
		expected := fmt.Sprintf("Created a parking lot with %v slots", slot)
		if expected != got {
			t.Errorf("MaintainBeneficiary() got = %v, expected = %v", got, expected)
		} else {
			t.Log(6)
		}
	}
}

func TestRagistration_ParkCar(t *testing.T) {
	ragistrationService.CreateParkingLot(2)
	tests := []struct {
		no, colour, expected string
		isErrror             bool
	}{
		{"KA-01-HH-1234", "White", "Allocated slot number: 1", false},
		{"KA-01-HH-9999", "White", "Allocated slot number: 2", false},
		{"", "White", "Allocated slot number: 2", true},
		{"", "", "Allocated slot number: 2", true},
		{"KA-01-HH-9999", "Black", "Sorry, parking lot is full", true},
	}
	for _, test := range tests {
		got, err := ragistrationService.ParkCar(test.no, test.colour)
		if (err != nil) != test.isErrror {
			t.Errorf("Ragistration.ParkCar() error = %v", err)
		} else {
			if (!test.isErrror) && test.expected != got {
				t.Errorf("Ragistration.ParkCar() got = %v, expected = %v", got, test.expected)
			} else {
				t.Log(got)
			}
		}

	}
}

func TestRagistration_GetParkingLotStatus(t *testing.T) {
	ragistrationService.CreateParkingLot(3)
	ragistrationService.ParkCar("KA-01-HH-1234", "White")
	ragistrationService.ParkCar("KA-01-HH-9999", "White")
	slots, err := ragistrationService.GetParkingLotStatus()
	if err != nil {
		t.Errorf("TestRagistration_GetParkingLotStatus() error = %v", err)
	}
	for _, ragi := range slots {
		if len(ragi.RagistrationNo) > 0 {
			t.Logf("%v \t\t %v \t\t %v\n", ragi.Slot, ragi.RagistrationNo, ragi.Colour)
		}
	}
}

func TestRagistration_LeaveCar(t *testing.T) {
	ragistrationService.CreateParkingLot(3)
	ragistrationService.ParkCar("KA-01-HH-1234", "White")
	ragistrationService.ParkCar("KA-01-HH-9999", "White")
	ragistrationService.ParkCar("KA-01-PP-1000", "Black")
	tests := []struct {
		slot     int
		expected string
		wantErr  bool
	}{
		{1, "Slot number 1 is free", false},
		{1, "Slot number 2 is free", true},
		{2, "Slot number 2 is free", false},
	}
	for _, tt := range tests {
		got, err := ragistrationService.LeaveCar(tt.slot)
		if (err != nil) != tt.wantErr {
			t.Errorf("Ragistration.LeaveCar() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if (!tt.wantErr) && tt.expected != got {
			t.Errorf("Ragistration.LeaveCar() got= %v, want %v", got, tt.expected)
		} else {
			t.Log(got)
		}
	}
}
