package model

import (
	"IS_project/helper"
	"IS_project/internal"
	"fmt"
	"log"
)

type User struct {
	Id      string
	Name    string
	Surname string
	Phone   string
	Mail    string
}

func NewUser(id, name, surname, phone, mail string) *User {
	return &User{
		Id:      id,
		Name:    name,
		Surname: surname,
		Phone:   phone,
		Mail:    mail,
	}
}

func (u *User) CheckIdExist() bool {
	db := internal.DB()
	defer db.Close()

	var row string = ""

	db.QueryRow(fmt.Sprintf("select ID_KORISNIKA from KORISNIK where ID_KORISNIKA = \"%s\"", u.Id)).Scan(&row)
	defer db.Close()
	fmt.Printf("%v", row)
	if len(row) > 0 {
		return true
	} else {
		return false
	}
}

func (u *User) SaveUser() error {
	db := internal.DB()
	defer db.Close()

	insertQuery := fmt.Sprintf("insert into KORISNIK (ID_KORISNIKA, IME, PREZIME, TELEFON, EMAIL) values (\"%s\", \"%s\", \"%s\", \"%s\", \"%s\");", u.Id, u.Name, u.Surname, u.Phone, u.Mail)
	_, err := db.Query(insertQuery)
	if err != nil {
		return err
	}

	return nil
}

func ValidateUserByPhone(phone string) User {
	db := internal.DB()
	defer db.Close()

	user := User{}
	queryString := fmt.Sprintf("select ID_KORISNIKA, IME, PREZIME, TELEFON, EMAIL from KORISNIK where TELEFON = \"%s\";", phone)
	err := db.QueryRow(queryString).Scan(&user.Id, &user.Name, &user.Surname, &user.Phone, &user.Mail)
	if err != nil {
		log.Printf("Error while fetching User by phone\n%v", err)
	}

	return user
}

func (u *User) GetUserById() error {
	db := internal.DB()
	defer db.Close()

	queryString := fmt.Sprintf("select IME, PREZIME, TELEFON, EMAIL from KORISNIK where ID_KORISNIKA = \"%s\";", u.Id)
	err := db.QueryRow(queryString).Scan(&u.Name, &u.Surname, &u.Phone, &u.Mail)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateUser() error {
	if u.CheckIdExist() {
		u.Id = helper.GenerateRandomString(20)

	} else {
		// Create User
		u.Id = helper.GenerateRandomString(20)
	}
	err := u.SaveUser()
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(id string) (*User, error) {
	db := internal.DB()
	defer db.Close()

	user := User{}
	queryString := fmt.Sprintf("select * from KORISNIK where ID_KORISNIKA = \"%s\";", id)
	err := db.QueryRow(queryString).Scan(&user.Id, &user.Name, &user.Surname, &user.Phone, &user.Mail)

	if err != nil {
		return &User{}, err
	}

	return &user, nil
}
