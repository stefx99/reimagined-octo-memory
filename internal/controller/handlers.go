package controller

import (
	"IS_project/helper"
	"IS_project/internal/model"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Fatalf("Can't load .html files\n%v", err)
	}
	tmpl.Execute(w, "")
}

func reservationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/rezervacija.html")
	if err != nil {
		log.Fatalf("Can't load .html files\n%v", err)
	}

	tmpl.Execute(w, struct {
		Message string
	}{Message: ""})
}

func submitReservation(w http.ResponseWriter, r *http.Request) {
	handleUser := model.NewUser(r.FormValue("username"), r.FormValue("name"), r.FormValue("surname"), r.FormValue("phone"), r.FormValue("email"))

	if err := handleUser.ValidateUser(); err != nil {
		log.Printf("Error during user validation\n%v", err)
	}

	requestedDate := helper.ParseDate(r.FormValue("datumRezervacije"))
	wps := model.GetAvailableWorkplaceByDate(requestedDate, r.FormValue("category"))
	if len(wps) == 0 {
		tmpl, _ := template.ParseFiles("static/rezervacija.html")

		tmpl.Execute(w, struct {
			Message string
		}{Message: "Sva mesta za izabrani sto su zauzeta, molimo pokusajte sa drugim stolom."})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reserv := model.Reservation{
		Id:        int32(rand.Intn(999999999)),
		Workplace: wps[0],
		User:      *handleUser,
		Date:      requestedDate,
	}
	if err := reserv.SaveReservation(); err != nil {
		http.Redirect(w, r, "/reservation", http.StatusBadRequest)
		log.Printf("Fail while saving reservation\n%v", err)
	}
	http.Redirect(w, r, "/reservation/success", http.StatusMovedPermanently)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/login.html")
	if err != nil {
		log.Printf("Can't load .html template file\n%v", err)
	}

	tmpl.Execute(w, struct {
		Message string
	}{Message: ""})
}

func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/login.html")
	if err != nil {
		log.Printf("Can't load .html template file\n%v", err)
	}

	phone := r.FormValue("phone")
	handleUser := model.ValidateUserByPhone(phone)
	// Login failed
	if len(handleUser.Id) == 0 {
		tmpl.Execute(w, struct {
			Message string
		}{Message: "Zadati broj ne postoji u nasoj bazi podataka."})
	}

	// Login successful
	successTmpl, err := template.ParseFiles("static/profile.html")
	if err != nil {
		log.Printf("Can't load .html template file\n%v", err)
	}

	if rsvList := model.GetReservationsByUserId(handleUser.Id); len(rsvList) > 0 {
		err = successTmpl.Execute(w, rsvList)
		if err != nil {
			log.Println(err)
		}
	}

}

func reservationSuccessHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/placanje.html")
	if err != nil {
		log.Printf("Can't load .html template\n%v", err)
	}
	tmpl.Execute(w, "")
}
