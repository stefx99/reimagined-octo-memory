package model

import (
	"IS_project/helper"
	"fmt"
	"testing"
)

func TestWorkplace_GetAllWorkspaces(t *testing.T) {

}

func TestGetPriceById(t *testing.T) {
	price, err := GetPriceById(3)
	fmt.Println(price)
	if err != nil || price == float32(0) {
		t.Fail()
	}

}

func TestWorkplace_GetWorkPlaceByIdMethod(t *testing.T) {
	w := Workplace{Id: 5}

	w.GetWorkPlaceById()

	if w.Price == float32(0) {
		t.Fail()
	}
}

func TestGetAvailableWorkplaceByDate(t *testing.T) {
	w := GetAvailableWorkplaceByDate(helper.ParseDate("2023-05-26"), "budget")

	for _, v := range w {
		if v.Id == 5 {
			t.Fail()
		}
	}
}
