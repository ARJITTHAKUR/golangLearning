package main

import (
	"fmt"
	"image"
	"image/png"
	"io/fs"
	"log"
	"os"
	"strconv"
)

func main() {

	// implementation using FS ============

	// path := "./images"

	// fileSystem := os.DirFS(path)

	// fs.WalkDir(fileSystem, ".", func(p string, d fs.DirEntry, err error) error {
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(p)
	// 	return nil
	// })
	// =====================

	// implementation using os ================
	// files, err := os.ReadDir("./images")

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// newfile, err := os.OpenFile("fileNames.txt", os.O_RDWR|os.O_CREATE, 0755)
	// defer newfile.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// 	bts := []byte(fmt.Sprintf("%s\n", file.Name()))
	// 	newfile.Write(bts)
	// }
	//========================================

	files, err := os.ReadDir("./images")

	if err != nil {
		log.Fatal(err)
	}
	filePointer := make(chan *os.File)
	ch := make(chan struct{})
	for idx, file := range files {

		go func(file fs.DirEntry, idx int) {
			currFile, err := os.Open(fmt.Sprintf("./images/%s", file.Name()))
			if err != nil {
				log.Fatal(err)
			}
			ConvertImageToThumbnail(idx, currFile, filePointer)

			ch <- struct{}{}
		}(file, idx)

		// currFile, err := os.Open(fmt.Sprintf("./images/%s", file.Name()))
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// go ConvertImageToThumbnail(idx, currFile, filePointer)

	}
	for range files {
		<-ch
	}
	// for fd := range filePointer {
	// 	fmt.Println("recieved fd of ", fd.Name())
	// 	err := fd.Close()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
}

func ConvertImageToThumbnail(idx int, src *os.File, filePointer chan<- *os.File) {
	fmt.Println("starting manipulation")
	img, err := png.Decode(src)

	if err != nil {
		log.Fatal(err)
	}

	xs := img.Bounds().Size().X
	ys := img.Bounds().Size().Y

	width, height := 128, 128

	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 * aspect)
	} else {
		height = int(128 * aspect)
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	// scaling image
	workingImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			workingImg.Set(x, y, workingImg.At(srcx, srcy))
		}
	}

	newImageFile, err := os.Create(fmt.Sprintf("./newImages/%s", strconv.Itoa(idx)+".png"))

	defer newImageFile.Close()

	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(newImageFile, workingImg)
	fmt.Printf("successfully created thumbnail for : %s \n", src.Name())
	if err != nil {
		log.Fatal(err)
	}
	src.Close()
	// filePointer <- srcs

}
