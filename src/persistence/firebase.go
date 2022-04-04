package persistence

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	"cloud.google.com/go/firestore"
	cloud "cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"

	"google.golang.org/api/option"
)

var APP *firebase.App
var Firebase_client *firestore.Client
var Firebase_storage *cloud.Client
var AccessId string
var Pkey string
var Bucket_name string

func InitFirebase() {
	var err error
	sa_path := os.Getenv("SA_PATH")
	sa := option.WithCredentialsFile(sa_path)
	// APP, err = firebase.NewApp(context.Background(), nil, sa)
	// if err != nil {
	// 	log.Fatalf("error initializing app: %v\n", err)
	// }

	// Firebase_client, client_err := APP.Firestore(context.Background())
	// if client_err != nil {
	// 	log.Fatalf("error initializing Firestore: %v\n", err)
	// } else {
	// 	fmt.Println(Firebase_client)
	// }

	var storage_err error
	Firebase_storage, storage_err = cloud.NewClient(context.Background(), sa)
	if storage_err != nil {
		log.Fatalf("error initializing cloud.NewClient: %v\n", err)
	}

	Bucket_name = os.Getenv("BUCKET_NAME")

	// Open our jsonFile
	jsonFile, err := os.Open(sa_path)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	AccessId = result["client_email"].(string)
	Pkey = result["private_key"].(string)

	jsonFile.Close()
}
