package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"
)

func main() {

	imageFile, err := os.Open("image.jpg")
	if err != nil {
		log.Fatal(err)
	}
	// scanner := bufio.NewScanner(imageFile)
	// defer imageFile.Close()
	// for scanner.Scan() {
	// 	fmt.Println("\nval : ", scanner.Text())
	// }
	// fmt.Println(imageFile)
	imageData, _, err := image.Decode(imageFile)

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(imageData, str)
	fmt.Println(imageData.Bounds())
	out, err := os.Create("output2.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	chars := []string{"@", "#", "S", "%", "?", "*", "+", ";", ":", ",", "."}
	// fmt.Println(len(chars))
	// return
	width := 200
	aspectRatio := imageData.Bounds().Dy() / imageData.Bounds().Dx()
	height := width * aspectRatio
	resizedImage := resize(imageData, width, height)
	// basically at this point we can have nested loops and try to print the image to console
	for i := 0; i < imageData.Bounds().Dy(); i++ {

		for j := 0; j < imageData.Bounds().Dx(); j++ {

			// c := imageData.At(j, i)
			c := resizedImage.At(j, i)
			r, g, b, _ := c.RGBA()
			gray := uint8((r + g + b) / 3 >> 8)
			// fmt.Println(gray)
			// gray := color.GrayModel.Convert(c).(color.Gray)
			// fmt.Println(gray.Y, len(chars))
			index := (int(gray) * len(chars)) / 255
			out.WriteString(chars[index])
			// model := imageData.ColorModel().Convert()
			// line = append(line, r, g, b, a)
			// line = append(line, model)

			// break
			// trying grey
			// grey := image.NewGray(imageData.Bounds())
			// grey.
		}
		out.WriteString("\n")
		// fmt.Println(line)
	}
}

func resize(img image.Image, width int, height int) image.Image {

	newImage := image.NewRGBA(image.Rect(0, 0, width, height))
	scaleX := float64(img.Bounds().Dx()) / float64(width)
	scaleY := float64(img.Bounds().Dy()) / float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// avg color at this scale down or up
			srcX := int(x * int(scaleX))
			srcY := int(y * int(scaleY))
			newImage.Set(x, y, img.At(srcX, srcY)) // this picks color from the src and adds to the dest image
		}
	}
	return newImage
}
