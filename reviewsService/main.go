package reviewsService

import (
	"log"
	"net/http"
	"reviewsService/routes"
)

func main() {
	router := routes.InitRoutes()

	log.Println("Catalog Service running on port 8083")
	err := http.ListenAndServe(":8083", router)
	if err != nil {
		log.Fatal(err)
	}
}
