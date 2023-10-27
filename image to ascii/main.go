package main

import (
	"fmt"
	"image"
	"image/color"
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

	width := 80
	aspectRatio := float64(imageData.Bounds().Dy()) / float64(imageData.Bounds().Dx())
	height := int(float64(width) * aspectRatio)
	// fmt.Println(imageData, str)
	fmt.Println(imageData.Bounds())
	out, err := os.Create("output.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	chars := []string{"@", "#", "S", "%", "?", "*", "+", ";", ":", ",", "."}
	// basically at this point we can have nested loops and try to print the image to console
	for i := 0; i < height; i++ {
		// line :=
		// var line []uint32
		// var line []interface{}

		for j := 0; j < width; j++ {
			// imageData.At()

			c := imageData.Bounds().At(i, j)
			// r, g, b, a := color.RGBA()

			gray := color.GrayModel.Convert(c).(color.Gray)

			index := int(gray.Y) * len(chars) / 255
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
