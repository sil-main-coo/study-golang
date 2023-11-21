package main

import (
	"fmt"
	stringHelper "happy-new-year/helpers"
	"validator"
)

func main() {
	fmt.Println("This main func")
	stringHelper.PutStringToIn()
	validator.CheckValidEmail()
}
