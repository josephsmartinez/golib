package main

import (
	"fmt"
	"golib/datahandler"
)

func main() {

	// Call all functions form
	dh := datahandler.NewDataHandler()
	text := "Some text to a file."
	dh.PrintToFile(&text, "test.txt")

	json := `{"foo":"bar"}`
	dh.PrintToFile(&json, "test.json")

	// Read from the test file to pring to STDOUT
	text_text := dh.ReadJSONFile("test.txt")
	fmt.Printf("%s\n", text_text)
	json_text := dh.ReadJSONFile("test.json")
	fmt.Printf("%s\n", json_text)
}
