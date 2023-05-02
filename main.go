package main

import (
	"fmt"
	"log"
	"test-case/config"
	"test-case/docs"
	"test-case/routes"
	"test-case/utils"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {

	// for load godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Test Case"
	docs.SwaggerInfo.Description = "This is a Test Case."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.Getenv("SWAGGER_HOST", "localhost:8080")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// database connection
	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// router
	r := routes.SetupRouter(db)

	fmt.Println("http://localhost:8080/swagger/index.html")
	// just remove port 8080
	r.Run()
}
