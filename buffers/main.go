package main

import (
	"bufio"
	"fmt"
)

type Writer int

func (w *Writer) Write(p []byte) (n int, err error) {
	fmt.Println(len(p))
	return len(p), nil
}

func main() {
	w := new(Writer)
	fmt.Println("unbuffer i/o")
	w.Write([]byte{'a'})
	w.Write([]byte{'b'})
	w.Write([]byte{'c'})
	w.Write([]byte{'d'})

	fmt.Println("buffered i/o")
	bw := bufio.NewWriterSize(w, 3)
	bw.Write([]byte{'a'})
	bw.Write([]byte{'b'})
	bw.Write([]byte{'c'})
	bw.Write([]byte{'d'})
	if err := bw.Flush(); err != nil {
		panic(err)
	}

	fmt.Println("larger writes")
	bwNew := bufio.NewWriterSize(w, 3)
	bwNew.Write([]byte("abcd"))
}
