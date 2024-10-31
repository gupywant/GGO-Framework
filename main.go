// main.go
package main

import (
	"fmt"
	"net/http"
    "GGO/config"
	"GGO/routes"
)

func main() {

	config.ConnectDB()

	r := routes.ApiRouter()

	fmt.Println("Listen On Port 8080")
	http.ListenAndServe(":8080", r)
}
