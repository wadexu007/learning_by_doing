package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/services"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders, err := services.GetAllOrders("orders.csv")
	if err != nil {
		http.Error(w, "[ERROR] Not found orders", http.StatusNotFound)
		return
	}
	if len(orders) == 0 {
		http.Error(w, "[WARN] No orders found, please palce order", http.StatusNotFound)
		return
	}
	log.Println("[INFO] get all orders")
	json.NewEncoder(w).Encode(orders)
}

func GetOrderByPizzaID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	log.Println("[DEBUG] params ====", params)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "[ERROR] Pizza ID can't convert to int", http.StatusBadRequest)
		return
	}
	os, err := services.GetOrderByID(id)
	if err != nil || os == nil {
		http.Error(w, "[WARN] Can't find order by PizzaID: "+strconv.Itoa(id), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(os)

}

func PlaceOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var o services.Order

	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, "[ERROR] Can't decode body", http.StatusBadRequest)
		return
	}

	error := services.PlaceOrder(o)
	if error != nil {
		http.Error(w, "[ERROR] Can't Place order", http.StatusBadRequest)
		return
	}
}
