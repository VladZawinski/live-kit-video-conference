package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/live-kit-video-conference/api"
	"github.com/live-kit-video-conference/repository"
	"github.com/live-kit-video-conference/sdk"
	"github.com/live-kit-video-conference/service"
)

func main() {
	err := godotenv.Load("development.env")
	if err != nil {
		log.Fatal("Error loading env file")
	}
	connectionString := os.Getenv("DB_CONNECTION")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	sdkServices := sdk.InjectSdkServices()
	repositories := repository.InjectRepository(db)
	appServices := service.InjectAppServices(*sdkServices, *repositories)
	api.BuildHandlers(*appServices)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	defer db.Close()

}
