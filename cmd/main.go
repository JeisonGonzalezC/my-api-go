package main

import (
	"log"
	"net/http"
	"os"

	"myapi/infrastructure/database"
	"myapi/internal/bootstrap"
	"myapi/router"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Panic("Error loading .env file")
	}

	// Conectar a la base de datos
	database.ConnectDB()

	// Migrar la base de datos
	// database.Migrate(database.DB)

	mux := router.SetupRoutes(bootstrap.InitApp())

	// Definir puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor
	log.Println("Servidor corriendo en http://localhost:" + port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Println("Error al iniciar el servidor:", err)
	}
}
