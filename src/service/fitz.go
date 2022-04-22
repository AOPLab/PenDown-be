package service

import (
	"fmt"
	"image/jpeg"
	"io"

	"github.com/gen2brain/go-fitz"
)

func Fitz() {
	doc, err := fitz.New("test.pdf")
	if err != nil {
		panic(err)
	}

	defer doc.Close()

	img, err := doc.Image(0)
	if err != nil {
		panic(err)
	}

	// f, err := os.Create(filepath.Join("./tempFile", "temp.jpg"))
	// if err != nil {
	// 	panic(err)
	// }

	r, w := io.Pipe()

	go func() {
		jpeg.Encode(w, img, &jpeg.Options{jpeg.DefaultQuality})
		w.Close()
	}()

	// err = jpeg.Encode(w, img, &jpeg.Options{jpeg.DefaultQuality})
	// if err != nil {
	// 	panic(err)
	// }

	// f, err = os.Open("./tempFile/temp.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	UploadImg("temp.jpg", r)

	url, _ := SignedFileUrl("temp.jpg")
	fmt.Print(url)
}
