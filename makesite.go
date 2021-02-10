package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
)

type FileCont struct {
	Title   string
	Content string
}

func save() {
	// 1. Read in the contents of the provided first-post.txt file.
	fileContents, err := ioutil.ReadFile("first-post.txt")

	// 3. Render the contents of first-post.txt using Go Templates and print it to stdout.
	fmt.Println(string(fileContents))
	if err != nil {
		panic(err)
	}

	content := FileCont{
		Title:   "first-post-w",
		Content: string(fileContents),
	}
	t := template.Must(template.ParseFiles("template.tmpl"))

	// 4. Write the HTML template to the filesystem to a file. Name it first-post.html
	t.Execute(file, content)

}

func main() {
	save()
}
