package persistence

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"context"

	"cloud.google.com/go/firestore"
	cloud "cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"

	"google.golang.org/api/option"
)

var APP *firebase.App
var Firebase_client *firestore.Client
var Firebase_storage *cloud.Client

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

	Firebase_client, client_err := APP.Firestore(context.Background())
	if client_err != nil {
		log.Fatalf("error initializing Firestore: %v\n", err)
	} else {
		fmt.Println(Firebase_client)
	}

	Firebase_storage, storage_err := cloud.NewClient(context.Background(), sa)
	if storage_err != nil {
		log.Fatalf("error initializing cloud.NewClient: %v\n", err)
	} else {
		fmt.Println(Firebase_storage)
	}

	// Test: upload a file
	bucket := os.Getenv("BUCKET_NAME")
	filePath := "test.txt"
	src := strings.NewReader("Hello World!\n")

	wc := Firebase_storage.Bucket(bucket).Object(filePath).NewWriter(context.Background())
	_, err = io.Copy(wc, src)
	if err != nil {
		log.Fatalf("error: %v\n", err)
		return
	}
	if err := wc.Close(); err != nil {
		log.Fatalf("error: %v\n", err)
		return
	}
}
