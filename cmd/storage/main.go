package main

import (
	"fmt"
	"log"

	"github.com/NizamutdinovEmir/01-workshop-golang/storage/internal/storage"
)

func main() {

	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetById(file.ID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("It works!", restoredFile)

}
