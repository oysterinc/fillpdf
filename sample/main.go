package main

import (
	"log"

	"github.com/oysterinc/fillpdf"
)

func main() {
	// Create the form values.
	form := fillpdf.Form{
		"field_1": fillpdf.TextOption("HELLO"),
		"field_2": fillpdf.TextOption("WORLD"),
		// If the form had a button field it would look like this:
		// "fake_button": fillpdf.ButtonOption("buttonON")
	}

	// Fill the form PDF with our values.
	err := fillpdf.Fill(form, "form.pdf", "filled.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
