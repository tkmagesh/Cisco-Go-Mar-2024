package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var ErrDataProcessing error = errors.New("unable to process the data")
var ErrReadFile error = errors.New("error reading file")

func main() {
	var err error
	err = process()
	if err == ErrDataProcessing {
		fmt.Println("unable to process")
	} else if err == ErrReadFile {
		fmt.Println("unable to read the file")
	} else if err != nil {
		fmt.Println("unknown error :", err)
	} else {
		fmt.Println("All good!")
	}
}

func process() error {
	r := rand.Intn(int(time.Now().Unix())) % 3
	switch r {
	case 1:
		return ErrDataProcessing
	case 2:
		return ErrReadFile
	default:
		return errors.New("some random error")
	}
}
