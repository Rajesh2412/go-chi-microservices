package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Created a order"))
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listing a order"))
}
func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("order by ID"))
}
func (o *Order) PostById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created a order"))
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updated a order"))
}
func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Order is Deleted")
	w.Write([]byte("updated a order"))
}
