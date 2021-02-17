package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"strings"

)

type content struct {
	Description string
}

func readFile(name string) string {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(fileContents)

}
func writeTemplateToFile(lang string, templateName string, fileName string) {
	/*
			Creates new template with the filename given
	*/

	c := content{Description: readFile(fileName)}
	t := template.Must(template.New("template.tmpl").ParseFiles(templateName))

	filter := filterInput(fileName) //option 1
	// f, err := os.Create(arg[:len(arg)-4] + ".html") //option 2
	f, err := os.Create(filter)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, c)
	if err != nil {
		panic(err)
	}
}



func filterInput(input string) string {
	/*
		Makesite v1.1
			Traverse through input until you reach '.', then add '.html' to the end.
			return s
	*/
	ext := ".html"
	s := strings.Split(input, ".")[0] + ext
	return s
}

func main() {
	// arg := os.Args[1] // Makesite MVP
}
