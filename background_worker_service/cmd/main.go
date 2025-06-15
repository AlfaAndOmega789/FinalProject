package cmd

import (
	"log"
	"net/http"
)

func main() {
	//router := routes.InitRoutes()

	log.Println("Catalog Service running on port 8085")
	err := http.ListenAndServe(":8085", router)
	if err != nil {
		log.Fatal(err)
	}
}
