package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"indock-upload-service/app/handlers"
	"indock-upload-service/app/utils"
	"indock-upload-service/config"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	loadConfig := config.LoadConfig()

	uploader := utils.NewS3Uploader(loadConfig.S3)

	mux := http.NewServeMux()
    mux.HandleFunc("/upload", handlers.HandleUpload(uploader))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

    handler := cors.Default().Handler(mux)

	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
