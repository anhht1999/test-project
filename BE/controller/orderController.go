package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"ocg-be/models"
	repo "ocg-be/repositories"

	"github.com/gorilla/mux"
)

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orders := repo.GetAllOrder()
	json.NewEncoder(w).Encode(orders)

}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	order := repo.GetOrderByID(int64(id))
	json.NewEncoder(w).Encode(order)
}

func UpdateStatusOrderByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var order models.Order
	json.Unmarshal(requestBody, &order)
	repo.UpdateStatusOrder(order)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "success",
	})
}

func CreateOrder(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = repo.CreateOrder(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "success"})
}

func GetRevenueByDay(w http.ResponseWriter, r *http.Request) {
	revs, err := repo.GetRevenueByDay(7)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(revs)
	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	"Revenue": revs,
	// })
}