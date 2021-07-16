package routes

import (
	"net/http"
	"ocg-be/controller"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {

	//api-products
	r.HandleFunc("/products", controller.GetProducts).Methods(http.MethodGet)               //GET_PRODUCTS
	r.HandleFunc("/products", controller.CreateProduct).Methods(http.MethodPost)            //POST_PRODUCT
	r.HandleFunc("/products/{id}", controller.GetProductById).Methods(http.MethodGet)       //GET_PRODUCT_By_Id
	r.HandleFunc("/products/{id}", controller.DeleteProductById).Methods(http.MethodDelete) //DELETE_PRODUCT_By_Id
	r.HandleFunc("/products", controller.UpdateProduct).Methods(http.MethodPatch)           //GET_PRODUCT_By_Id

	//api-image
	r.HandleFunc("/uploads", controller.CreateProductImage).Methods(http.MethodPost)
	image := http.StripPrefix("/images/", http.FileServer(http.Dir("./")))
	r.PathPrefix("/images/").Handler(image)

	//Category
	r.HandleFunc("/categories", controller.GetAllCategory).Methods(http.MethodGet)

	r.HandleFunc("/login", controller.Login).Methods(http.MethodPost) //LOGIN
	r.HandleFunc("/register", controller.Register).Methods(http.MethodPost) //REGISTER
	r.HandleFunc("/logout", controller.Logout).Methods(http.MethodPost) //LOGOUT

	r.HandleFunc("/orders", controller.GetAllOrders).Methods(http.MethodGet) //GET_ORDERs
	r.HandleFunc("/orders", controller.CreateOrder).Methods(http.MethodPost) //POST_ORDERs
	
}
