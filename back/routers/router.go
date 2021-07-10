package routers

import (
	"net/http"
	"back/controllers"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {

	//Product
	r.HandleFunc("/products", controllers.GetAllProducts).Methods(http.MethodGet)    //GET_PRODUCTS
	r.HandleFunc("/products/{id}", controllers.GetProductByID).Methods(http.MethodGet)    //GET_PRODUCT_BY_ID
	r.HandleFunc("/products", controllers.CreateProduct).Methods(http.MethodPost) //POST_PRODUCT
	r.HandleFunc("/products/{id}", controllers.UpdateProductByID).Methods(http.MethodPut) //PUT_PRODUCT
	r.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods(http.MethodDelete) //DELETE_PRODUCT
	
	//Category
	r.HandleFunc("/categories", controllers.GetAllCategory).Methods(http.MethodGet)    //GET_CATEGORIES
	r.HandleFunc("/categories", controllers.CreateCategory).Methods(http.MethodPost) //POST_CATEGORY
	r.HandleFunc("/categories/{id}", controllers.UpdateCategoryByID).Methods(http.MethodPut) //PUT_CATEGORY

	//middiewares
	r.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
	r.HandleFunc("/logout", controllers.Logout).Methods(http.MethodPost)
	r.HandleFunc("/register", controllers.Register).Methods(http.MethodPost)

	//user
	r.HandleFunc("/users", controllers.GetAllProducts).Methods(http.MethodGet)    //GET_USERS
	r.HandleFunc("/users/{id}", controllers.GetProductByID).Methods(http.MethodGet)    //GET_USER_BY_ID
	r.HandleFunc("/users", controllers.CreateProduct).Methods(http.MethodPost) //POST_USER
	r.HandleFunc("/users/{id}", controllers.UpdateProductByID).Methods(http.MethodPut) //PUT_USER
	r.HandleFunc("/users/{id}", controllers.DeleteProduct).Methods(http.MethodDelete) 
}