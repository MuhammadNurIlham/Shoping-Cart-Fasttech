package routes

import (
	"fasttech/handlers"
	"fasttech/pkg/mysql"
	"fasttech/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/carts", h.FindCarts).Methods("GET")
	r.HandleFunc("/cart/{id}", h.GetCart).Methods("GET")
	r.HandleFunc("/cart", h.CreateCart).Methods("POST")
	r.HandleFunc("/cart/{id}", h.UpdateCart).Methods("PATCH")
	r.HandleFunc("/cart/{id}", h.DeleteCart).Methods("DELETE")
}