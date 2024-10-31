package routes

import (
	"net/http"
	"GGO/controllers"
	"GGO/utils"

	"github.com/gorilla/mux"
)

func ApiRouter() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	// Routes without the need to repeat the "/api" prefix
	authController := &controllers.AuthController{}
	api.HandleFunc("/login", utils.MethodCheckJson(authController.Login, http.MethodPost))
	api.HandleFunc("/logout", utils.MethodCheckJson(authController.Logout, http.MethodGet))

	// Example of another route group
	// api.HandleFunc("/dashboard", middlewares.AuthMiddleware(http.HandlerFunc(controllers.ShowDashboard)))

	return router
}