package routes

import (
	"net/http"
	"ocg-be/controller"
	"ocg-be/middlewares"
	"ocg-be/util"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {
	routeCustomer := r.PathPrefix("/customers").Subrouter()
	routeAdmin := r.PathPrefix("/admin").Subrouter()
	routeAdmin.Use(middlewares.IsAuthorized)
	routeCustomer.Use(middlewares.IsAuthenticated)
	routePublic(r)
	routesAdmin(routeAdmin)
	routesCustomer(routeCustomer)
}

func routePublic(r *mux.Router) {

	//product
	r.HandleFunc("/products", controller.GetProducts).Methods(http.MethodGet) //GET_PRODUCTS
	r.HandleFunc("/products/{id}", controller.GetProductById).Methods(http.MethodGet)//GET_PRODUCT_By_Id
	
	//category
	r.HandleFunc("/categories", controller.GetAllCategory).Methods(http.MethodGet)
	
	//login-register-logout
	r.HandleFunc("/login", controller.Login).Methods(http.MethodPost) //LOGIN
	r.HandleFunc("/register", controller.Register).Methods(http.MethodPost) //REGISTER
	r.HandleFunc("/logout", controller.Logout).Methods(http.MethodPost) //LOGOUT

}

func routesAdmin(r *mux.Router) {

	//product 
	r.HandleFunc("/products", controller.GetProducts).Methods(http.MethodGet)               //GET_PRODUCTS
	r.HandleFunc("/products", controller.CreateProduct).Methods(http.MethodPost)            //POST_PRODUCT
	r.HandleFunc("/products/{id}", controller.DeleteProductById).Methods(http.MethodDelete) //DELETE_PRODUCT_By_Id
	r.HandleFunc("/products", controller.UpdateProduct).Methods(http.MethodPut)           //GET_PRODUCT_By_Id

	//image
	r.HandleFunc("/uploads", util.UploadFile).Methods(http.MethodPost)
	images := http.StripPrefix("/images/", http.FileServer(http.Dir("./uploads/")))
	r.PathPrefix("/images/").Handler(images)

	//order
	r.HandleFunc("/orders", controller.GetAllOrders).Methods(http.MethodGet) //GET_ORDERs
	r.HandleFunc("/orders/{id}", controller.GetOrderByID).Methods(http.MethodGet) //Get_ORDERs_BY_ID
	r.HandleFunc("/orders", controller.UpdateStatusOrderByID).Methods(http.MethodPut) //UPDATE_STATUS_ORDERs_BY_ID
	r.HandleFunc("/revenues", controller.GetRevenueByDay).Methods(http.MethodGet)
}

func routesCustomer(r *mux.Router) {
	r.HandleFunc("/orders", controller.CreateOrder).Methods(http.MethodPost) //POST_ORDERs
}