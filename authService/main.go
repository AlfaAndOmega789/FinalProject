package authService

import (
	"log"
	"net/http"
	"reviews/routes"
)

func main() {
	//router := routes.InitRoutes()

	log.Println("Catalog Service running on port 8084")
	err := http.ListenAndServe(":8084", router)
	if err != nil {
		log.Fatal(err)
	}
}
