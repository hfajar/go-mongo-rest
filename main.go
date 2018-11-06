package main

import (
	"encoding/json"
	"fmt"
	controller "go-mongo-rest/controller"
	model "go-mongo-rest/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPo for get all data po
func GetPo(w http.ResponseWriter, r *http.Request) {
	var data = []model.Po{}
	data = controller.FindAll()
	json.NewEncoder(w).Encode(data)
}

// GetPoByID for get data po by Id
func GetPoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var data []model.Po
	fmt.Println(params["id"])
	data = controller.FindOne(params["id"])
	json.NewEncoder(w).Encode(data)
}

// SetPo for input po to database
func SetPo(w http.ResponseWriter, r *http.Request) {
	var data = model.Po{}
	_ = json.NewDecoder(r.Body).Decode(&data)
	var hasil = controller.Insert(data)
	fmt.Println(hasil)
	json.NewEncoder(w).Encode(data)
}

// DeletePo for remove po from database by ID
func DeletePo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	fmt.Println("id dengan " + id + " akan segera dihapus")
	var hasil = controller.Delete(id)
	fmt.Println(hasil)
	var data = []model.Po{}
	data = controller.FindAll()
	json.NewEncoder(w).Encode(data)
}

// UpdatePo for remove po from database by ID
func UpdatePo(w http.ResponseWriter, r *http.Request) {
	var data = model.Po{}
	_ = json.NewDecoder(r.Body).Decode(&data)
	var hasil = controller.Update(data)
	fmt.Println(hasil)
	json.NewEncoder(w).Encode(data)
}

func main() {
	router := mux.NewRouter()
	fmt.Println("starting web server at http://localhost:8585/")
	router.HandleFunc("/po", GetPo).Methods("GET")
	router.HandleFunc("/po/{id}", GetPoByID).Methods("GET")
	router.HandleFunc("/po", SetPo).Methods("POST")
	router.HandleFunc("/po/{id}", DeletePo).Methods("DELETE")
	router.HandleFunc("/po", UpdatePo).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8585", router))
}
