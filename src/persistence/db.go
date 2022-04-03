package persistence

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"context"

	firebase "firebase.google.com/go"

	cloud "cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var DB *gorm.DB
var APP *firebase.App

func Initialize(dsn string) (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)           // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(100)          // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.

	return DB, err
}

func InitFirebase() {
	var err error
	sa_path := os.Getenv("SA_PATH")
	sa := option.WithCredentialsFile(sa_path)
	APP, err = firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	} else {
		fmt.Println(APP)
	}

	client, client_err := APP.Firestore(context.Background())
	if client_err != nil {
		log.Fatalf("error initializing Firestore: %v\n", err)
	} else {
		fmt.Println(client)
	}

	storage, storage_err := cloud.NewClient(context.Background(), sa)
	if storage_err != nil {
		log.Fatalf("error initializing cloud.NewClient: %v\n", err)
	} else {
		fmt.Println(storage)
	}
}
