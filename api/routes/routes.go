package routes

import (
	"../controller"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {

	router.HandleFunc("/team", controller.CreateTeam).Methods("POST")

	//controller.CreateTeam

}
