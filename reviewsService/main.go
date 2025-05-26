package reviewsService

import (
	"log"
	"net/http"
	"reviewsService/routes"
)

func main() {
	router := routes.InitRoutes()

	log.Println("Catalog Service running on port 8001")
	err := http.ListenAndServe(":8001", router)
	if err != nil {
		log.Fatal(err)
	}
}
