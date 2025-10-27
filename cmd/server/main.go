package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"lesson-management/entities"
	"lesson-management/pkg/common"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	common.InitDB()
	// Migrate all entities including new ones
	common.DB.AutoMigrate(
		&entities.Admin{},
		&entities.Teacher{},
		&entities.Student{},
		&entities.Lesson{},
	)

	api := InitRoutes()
	fmt.Println("✅ Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, api))
}
