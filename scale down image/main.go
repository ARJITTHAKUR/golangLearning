package main

import (
	"fmt"
	"image"
	jpeg "image/jpeg"
	"log"
	"os"
)

func main() {

	imageFile, err := os.Open("image.jpg")

	if err != nil {
		log.Fatal(err)
	}
	defer imageFile.Close()
	imageData, _, err := image.Decode(imageFile)

	if err != nil {
		log.Fatal(err)
	}

	width := 100
	aspect := imageData.Bounds().Dx() / imageData.Bounds().Dy()
	height := width * aspect

	fmt.Println("aspect :", aspect)

	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// make scale and pick color from the scale down or up color

	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			scaleX := (imageData.Bounds().Dx() / width) * x
			scaleY := (imageData.Bounds().Dy() / height) * y
			newImage.Set(x, y, imageData.At(scaleX, scaleY))
		}
	}

	// open a file
	newImageFile, err := os.OpenFile("modified.jpg", os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		log.Fatal(err)
	}

	err = jpeg.Encode(newImageFile, newImage, nil)

	if err != nil {
		log.Fatal(err)
	}

}
