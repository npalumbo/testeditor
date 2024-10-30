package main

import (
	"fmt"
	testeditorText "testeditor/internal/text"
)

func main() {
	text := testeditorText.CreateInternalTextExtension()

	text.Insert(0, "Hola Golang Sevilla")
	fmt.Println(text.Render())

	text.SearchAndReplace("Sevilla", "SEVILLA")
	fmt.Println(text.Render())
}
