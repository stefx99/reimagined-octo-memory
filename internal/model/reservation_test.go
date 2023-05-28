package model

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestSaveAndGetReservation(t *testing.T) {
	r := Reservation{
		Id: 1,
		User: User{
			Id:      "stefan",
			Name:    "Milojko",
			Surname: "Milojkovic",
			Phone:   "+361123",
			Mail:    "adsad@sada.com",
		},
		Workplace: Workplace{
			Id: 10,
		},
		Date: time.Now(),
	}

	err := r.SaveReservation()
	if err != nil {
		fmt.Println(err)
	}
	res, _ := GetReservationById(1)
	if res.Id != r.Id || res.User.Id != r.User.Id {
		t.Fatalf("Reservation mismatch, \n %v", res)
	}
}

func TestGetReservationById(t *testing.T) {
	r, err := GetReservationById(1)
	log.Printf("%v\n%v", *r, err)
	if len(r.User.Id) == 0 {
		t.Fail()
	}
}
