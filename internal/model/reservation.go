package model

import (
	"IS_project/helper"
	"IS_project/internal"
	"fmt"
	"log"
	"time"
)

type Reservation struct {
	User      User
	Workplace Workplace
	Id        int32
	Date      time.Time
}

func (r *Reservation) SaveReservation() error {
	db := internal.DB()
	defer db.Close()

	queryString := fmt.Sprintf("insert into REZERVACIJE (ID_REZERVACIJE, ID_KORISNIKA, ID_MESTA, DATUM) "+
		"values (%d, \"%s\", %d, \"%s\");", r.Id, r.User.Id, r.Workplace.Id, r.Date.Format("2006-01-02"))
	_, err := db.Query(queryString)
	if err != nil {
		return err
	}
	return nil
}

func GetReservationById(id int32) (*Reservation, error) {
	db := internal.DB()
	defer db.Close()

	stringDate := ""
	row := &Reservation{}
	queryString := fmt.Sprintf("select ID_REZERVACIJE, ID_KORISNIKA, ID_MESTA, DATUM from REZERVACIJE where ID_REZERVACIJE = %d;", id)
	err := db.QueryRow(queryString).Scan(&row.Id, &row.User.Id, &row.Workplace.Id, &stringDate)

	row.Date = helper.ParseDate(stringDate)

	if err != nil {
		log.Fatalf("Can't load Reservation row\n%v", err)
	}

	if err := row.User.GetUserById(); err != nil {
		log.Printf("Cannot get user model from it's ID\n%v", err)
	}
	if err := row.Workplace.GetWorkPlaceById(); err != nil {
		log.Printf("Cannot get workplace by it's ID%v", err)
	}
	if err != nil {
		return &Reservation{}, nil
	}
	log.Printf("%v", *row)

	return row, nil
}

func GetReservationsByUserId(id string) []Reservation {
	db := internal.DB()
	defer db.Close()

	queryString := fmt.Sprintf("select ID_KORISNIKA, ID_MESTA, ID_REZERVACIJE, DATUM from REZERVACIJE where ID_KORISNIKA = \"%s\";", id)
	rows, err := db.Query(queryString)
	if err != nil {
		log.Printf("Error while trying to fetch Reservations by UserID\n%v", err)
	}

	reservationList := []Reservation{}
	for rows.Next() {
		resv := Reservation{}
		helpDate := ""
		rows.Scan(&resv.User.Id, &resv.Workplace.Id, &resv.Id, &helpDate)

		resv.Date = helper.ParseDate(helpDate)
		resv.User.GetUserById()
		resv.Workplace.GetWorkPlaceById()
		reservationList = append(reservationList, resv)
	}
	return reservationList
}
