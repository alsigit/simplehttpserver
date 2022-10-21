package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"webserver/libs"
)

var PORT = ":8080"
var db []*libs.User
var svc1 = libs.MyUsers(db)
var people = []string{"Sigit", "Frisky", "Andi", "Diko", "Meuthia", "Rizky", "Theo"}
var wg sync.WaitGroup

type Employee struct {
	Name    string
	Address string
}

func main() {
	// go addData(people)

	http.HandleFunc("/", greet)
	http.HandleFunc("/get", getPerson)
	http.HandleFunc("/add", addPerson)
	http.HandleFunc("/jsonadd", newPerson)
	http.ListenAndServe(PORT, nil)

}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "hello world"
	w.Write([]byte(msg))
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	listUser := svc1.GetUser()
	data, _ := json.Marshal(listUser)
	w.Write(data)
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		name := r.FormValue("name")
		address := r.FormValue("address")

		new_employee := Employee{
			Name:    name,
			Address: address,
		}

		svc1.Register(&libs.User{Nama: new_employee.Name})
		json.NewEncoder(w).Encode(new_employee)

		return
	}
	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func newPerson(w http.ResponseWriter, r *http.Request) {
	var p *libs.User

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(p.Nama))
	w.Write([]byte(p.Alamat))
	svc1.Register(&libs.User{
		Nama:   p.Nama,
		Alamat: p.Alamat,
	})

}
