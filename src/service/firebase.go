package service

import (
	"context"
	"io"
	"log"
	"strings"
	"time"

	"github.com/AOPLab/PenDown-be/src/persistence"

	cloud "cloud.google.com/go/storage"
)

func SignedFileUrl(filePath string) (string, error) {
	url, err := cloud.SignedURL(persistence.Bucket_name, filePath, &cloud.SignedURLOptions{
		GoogleAccessID: persistence.AccessId,
		PrivateKey:     []byte(persistence.Pkey),
		Method:         "GET",
		Expires:        time.Now().Add(time.Hour),
	})
	if err != nil {
		log.Fatalf("Signed file error: %v\n", err)
		return "", err
	}
	return url, nil
}

// TODO: Change to file format
func UploadFile(filePath string) error {
	src := strings.NewReader("Hello World!\n")

	wc := persistence.Firebase_storage.Bucket(persistence.Bucket_name).Object(filePath).NewWriter(context.Background())
	_, err := io.Copy(wc, src)
	if err != nil {
		log.Fatalf("error: %v\n", err)
		return err
	}
	if err := wc.Close(); err != nil {
		log.Fatalf("error: %v\n", err)
		return err
	}
	return nil
}
