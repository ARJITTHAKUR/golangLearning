package main

import (
	"errors"
	"fmt"
)

func main() {

	files := []string{
		"file1",
		"file2",
		"file3",
		"file4",
		"file5",
		"file6",
		"file7",
		"file8",
		"file9",
	}

	data, err := workWithErrors(files)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v", data)

}

type thumbnail struct {
}

func workWithErrors(files []string) (thumbnails []string, err error) {
	var thumbnailsSlice []string
	type item struct {
		file string
		err  error
	}
	ch := make(chan item)

	for i := 0; i < len(files); i++ {
		go func(idx int) {
			var itemCurr item
			itemCurr.file, itemCurr.err = work(idx, files[idx])
			ch <- itemCurr
		}(i)
	}

	for range files {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbnailsSlice = append(thumbnailsSlice, it.file)
	}

	return thumbnailsSlice, nil
}

func work(i int, fileName string) (file string, err error) {

	if i == 5 {
		return "", errors.New("some error occured")
	}
	return fileName, nil
}
