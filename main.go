package main

import (
	"fasttech/database"
	"fasttech/pkg/mysql"
	"fasttech/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// initial Database
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("Server running localhost:5000");
	http.ListenAndServe("localhost:5000", r);

};