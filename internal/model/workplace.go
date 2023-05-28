package model

import (
	"IS_project/helper"
	"IS_project/internal"
	"fmt"
	"log"
	"time"
)

type Workplace struct {
	Id    int32
	Price float32
}

func (w *Workplace) GetAllWorkspaces() *[]Workplace {
	db := internal.DB()
	defer db.Close()
	rows, err := db.Query(`select * from COWORKINGPROSTOR;`)
	if err != nil {
		log.Println("Get all workspaces query error. \n", err)
	}
	defer rows.Close()

	var result []Workplace

	for rows.Next() {
		row := &Workplace{}
		if err := rows.Scan(&row.Id, &row.Price); err != nil {
			log.Println(err)
		}
		result = append(result, *row)
	}

	return &result
}

func (w *Workplace) GetWorkPlaceById() error {
	db := internal.DB()
	defer db.Close()

	queryString := fmt.Sprintf("select CENA from COWORKINGPROSTOR where ID_MESTA = \"%d\";", w.Id)
	err := db.QueryRow(queryString).Scan(&w.Price)
	if err != nil {
		return err
	}

	return nil
}

func GetPriceById(id int32) (float32, error) {
	db := internal.DB()
	defer db.Close()

	var price float32

	queryString := fmt.Sprintf("select CENA from COWORKINGPROSTOR where ID_MESTA = \"%d\"", id)
	err := db.QueryRow(queryString).Scan(&price)
	if err != nil {
		return float32(0), err
	}

	return price, nil
}

func GetAvailableWorkplaceByDate(date time.Time, category string) []Workplace {
	available := []Workplace{}
	minCena := 0
	maxCena := 0

	db := internal.DB()
	defer db.Close()

	switch category {
	case "budget":
		minCena = 0
		maxCena = 1500
	case "quatro":
		minCena = 1499
		maxCena = 2001
	case "duo":
		minCena = 2000
		maxCena = 99999
	}

	queryString := fmt.Sprintf("select ID_MESTA, CENA from COWORKINGPROSTOR where CENA > %d and CENA < %d and ID_MESTA not in (select COWORKINGPROSTOR.ID_MESTA FROM COWORKINGPROSTOR, REZERVACIJE where COWORKINGPROSTOR.ID_MESTA = REZERVACIJE.ID_MESTA and REZERVACIJE.DATUM = \"%s\");\n",
		minCena, maxCena, helper.StringDate(date))
	rows, err := db.Query(queryString)
	if err != nil {
		log.Fatalf("Error on execution query (get available workplace by date\n%v", err)
	}

	for rows.Next() {
		row := Workplace{}
		rows.Scan(&row.Id, &row.Price)

		available = append(available, row)
	}

	return available
}
