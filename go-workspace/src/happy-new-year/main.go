package main

import (
	"fmt"
	"happy-new-year/helpers"
	"validator"
)

func main() {
	fmt.Println("This main func")
	helpers.PutStringToIn()
	validator.CheckValidEmail()
}
