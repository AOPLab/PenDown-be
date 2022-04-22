package service

import (
	"image/jpeg"
	"io"
	"mime/multipart"

	"github.com/gen2brain/go-fitz"
)

func Fitz(filepath string, file multipart.File) error {
	// doc, err := fitz.New("test.pdf")
	doc, err := fitz.NewFromReader(file)
	if err != nil {
		return err
	}

	defer doc.Close()

	img, err := doc.Image(0)
	if err != nil {
		return err
	}

	r, w := io.Pipe()

	go func() {
		o := jpeg.Options{Quality: 80}
		jpeg.Encode(w, img, &o)
		w.Close()
	}()

	upload_err := UploadImg(filepath, r)
	if upload_err != nil {
		return upload_err
	}

	return nil
}
