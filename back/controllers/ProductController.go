package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"back/models"
	repo "back/repositories"
	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	products := repo.GetAllProducts()
	json.NewEncoder(w).Encode(products)
}

func GetProductByID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	product := repo.GetProductByID(int64(id))
	json.NewEncoder(w).Encode(product)
}

func UpdateProductByID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	product.Id = int64(id)
	err = repo.UpdateProductByID(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product = repo.GetProductByID(product.Id)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "success",
		"product updated": product,
	})
}

func CreateProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = repo.CreateProduct(product)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{ "message" : "success"})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id,_ := strconv.Atoi(vars["id"])
	repo.DeleteProduct(int64(id))
	json.NewEncoder(w).Encode(map[string]string{ "message" : "success"})
}