package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AOPLab/PenDown-be/src/config"
	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func test(c *gin.Context) {
	var message = "Hello world!"
	c.IndentedJSON(http.StatusOK, message)
}

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	username := os.Getenv("PG_USERNAME")
	password := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_DBNAME")

	dsn := "host=" + host + " port=" + pgPort + " user=" + username + " password=" + password + " dbname=" + dbName
	db, db_err := persistence.Initialize(dsn)
	if db_err != nil {
		log.Fatal("Error loading db")
	}

	db.AutoMigrate(&model.User{})

	// t, _ := time.Parse(time.RFC3339, "2030-01-02T15:04:05Z")
	// fmt.Println(t)
	// url := model.Url{
	// 	Original_url: "https://www.google.com.tw/",
	// 	Expired_date: t,
	// 	Url_id:       "ABCDEF",
	// }
	// Insert
	// db.Model(&model.Url{}).Create(&url)

	// port := os.Getenv("PORT")

	router := gin.Default()
	router.GET("/test", test)
	config.Routes(router)
	// port1 := ":" + port
	router.Run(":8080")
}
