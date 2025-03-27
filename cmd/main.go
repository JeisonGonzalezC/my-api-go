package main

import (
	"log"
	"net/http"
	"os"

	"myapi/infrastructure/database"
	"myapi/internal/bootstrap"
	"myapi/router"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// load .env file
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Panic("Error loading .env file")
	}

	// connect to database
	db := database.ConnectDB()

	// Migrate database
	// database.Migrate(database.DB)

	mux := router.SetupRoutes(bootstrap.InitApp(db))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("ORIGIN_FRONT_VUE")}, // My app frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // allow cookies
	})

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Println("Servidor corriendo en http://localhost:" + port)
	err := http.ListenAndServe(":"+port, c.Handler(mux))
	if err != nil {
		log.Println("Error al iniciar el servidor:", err)
	}
}
