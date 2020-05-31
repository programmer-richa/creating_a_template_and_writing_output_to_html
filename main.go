package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// Creating and assigning values to variables
	title := "Home"
	author := "Richa"
	// Creating a template and passing values to it.
	tpl := `<!DOCTYPE html>
	<html>
	<head>
	<title>` + title + `</title>
	</head>
	<body>
	
	<h1>` + author + `</h1>
	<p>This is a paragraph.</p>
	
	</body>
	</html>`
	fmt.Print(tpl)
	// Creating a file to save html to
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error occured while creating file")
	}
	defer nf.Close()
	// Copy data from tpl string to new file.
	io.Copy(nf, strings.NewReader(tpl))

}
