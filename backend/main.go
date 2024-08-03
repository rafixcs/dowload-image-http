package main

import (
	"log"
	"net/http"

	http_logger "github.com/azvaliev/http-logger"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

const UPLOAD_PATH = "./datasets/tmp"

func main() {

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir(UPLOAD_PATH))
	mux.Handle("/files/", http.StripPrefix("/files", fs))

	// Enalble CORS with a middleware setting witch sources are able to acces the resource
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Your React app's origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Custom-Header"},
		AllowCredentials: true,
	}).Handler(mux)

	/*http.HandleFunc("/dowload", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile()
	})*/

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	log.Print("Server started on localhost:3002")
	log.Fatal(http.ListenAndServe(":3002", http_logger.WithLogging(handler, logger)))
}
