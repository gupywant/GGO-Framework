// main.go
package main

import (
	"GGO/controllers"
	"fmt"
	"GGO/utils"
	"net/http"
    "GGO/config"
)

func main() {

	config.ConnectDB()
	// Declare controller
	authController := &controllers.AuthController{}

	// Simplify method checking for the login route (only POST allowed)
	http.HandleFunc("/login", utils.MethodCheckJson(authController.Login, http.MethodPost))

	// Simplify method checking for the logout route (only GET allowed)
	http.HandleFunc("/logout", utils.MethodCheckJson(authController.Logout, http.MethodGet))

	// http.Handle("/dashboard", middlewares.AuthMiddleware(http.HandlerFunc(controllers.ShowDashboard)))

	// Start the server

	fmt.Println("Listen On Port 8080")
	http.ListenAndServe(":8080", nil)
}
