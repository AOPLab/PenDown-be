package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AOPLab/PenDown-be/src/config"
	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	var message = "Hello world!"
	c.IndentedJSON(http.StatusOK, message)
}

func main() {
	// envErr := godotenv.Load()
	// if envErr != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	host := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	username := os.Getenv("PG_USERNAME")
	password := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_DBNAME")

	dsn := "host=" + host + " port=" + pgPort + " user=" + username + " password=" + password + " dbname=" + dbName
	db, db_err := persistence.InitDB(dsn)
	if db_err != nil {
		log.Fatal("Error loading db")
	}

	db.AutoMigrate(&model.User{}, &model.Follow{}, &model.School{}, &model.Course{}, &model.Tag{}, &model.Note{}, &model.Download{}, &model.NoteTag{}, &model.Saved{})

	// port := os.Getenv("PORT")

	persistence.InitFirebase()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           time.Minute,
	}))

	router.GET("/test", test)
	config.Routes(router)
	// port1 := ":" + port
	router.Run(":8080")
}
