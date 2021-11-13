package main

import (
	"booklogger/models"
	"fmt"
)

func main() {
	book := models.NewBook("Hello World")
	fmt.Println(book) //nolint:forbidigo
}
