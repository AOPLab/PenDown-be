package service

import (
	"context"
	"io"
	"log"
	"mime/multipart"
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

func UploadFile(filePath string, file multipart.File) error {
	wc := persistence.Firebase_storage.Bucket(persistence.Bucket_name).Object(filePath).NewWriter(context.Background())
	_, err := io.Copy(wc, file)
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
