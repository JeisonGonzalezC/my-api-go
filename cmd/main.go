package main

import (
	"log"
	"myapi/infrastructure/database"
	"myapi/router"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Panic("Error loading .env file")
	}

	database.ConnectDB()
	// database.Migrate(database.DB)

	mux := router.SetupRoutes()

	const port string = ":8080"
	log.Println("Servidor corriendo en http://localhost" + port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Println("Error al iniciar el servidor:", err)
	}
}
