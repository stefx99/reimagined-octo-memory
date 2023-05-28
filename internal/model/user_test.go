package model

import (
	"testing"
)

func TestUser_CreateAndCheckIfExist(t *testing.T) {
	u := NewUser("stefanaa", "Stefan", "Man", "062", "sass.com")

	if !u.CheckIdExist() {
		t.Fail()
	}

}

func TestGetUserById(t *testing.T) {
	user, _ := GetUserById("stefanaa")

	if user.Id != "stefanaa" {
		t.Fail()
	}
}

func TestGetUserByIdMethod(t *testing.T) {
	user := User{Id: "stefan"}
	user.GetUserById()

	if len(user.Name) == 0 {
		t.Fatalf("User Name not found")
	}
}

func TestValidateUserByPhone(t *testing.T) {
	u := ValidateUserByPhone("123213")
	if u.Id != "iprvabyuhezvmceuwqvb" {
		t.Fail()
	}
}
