package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"flag"
	"fmt"
	"log"
	"context"

	"cloud.google.com/go/translate"
	_ "cloud.google.com/go/translate/apiv3"
	"golang.org/x/text/language"

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


func translateText(targetLanguage, text string) (string, error) {
	// text := "The Go Gopher is cute"
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	// fmt.Println(resp[0].Text, nil)
	return resp[0].Text, nil
}


func writeTranslate(filename string, lang string) {
	/*
			Reads/translates the .txt files, writes them into a template file
	*/
	FileText := readFile(filename)

	contents, error := translateText(lang, FileText)
	if error != nil {
		panic(error)
	}
	bytesToWrite := []byte(contents)

	err := ioutil.WriteFile(filename, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func parser() {
	/*
			Collects files in given directory,
			checks if file includes '.txt',
			creates a '.html' file for the .txt files
	*/
	/* type, -dir=[name of directory you want to scan], to search for all .txt files you want to convert into a template.
	Default = current directory */
	var dir string
	flag.StringVar(&dir, "dir", ".", "This is the directory.")

	// type, -lang=[google language abbreviation], to choose translation. Default = espanol
	var lang string
	flag.StringVar(&lang, "lang", "es", "This is the language you want to translate, inputting google's language abbreviations.")
	flag.Parse()

	fmt.Println("Directory:", dir)
	fmt.Println("Language:", lang)

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filenameCheck(file.Name()) == true {
			fmt.Println(file.Name())
			writeTranslate(file.Name(), lang)                       // Google translate function
			writeTemplateToFile(lang, "template.tmpl", file.Name()) // writes file contents into newly-created template
		}
	}
}

func filenameCheck(filename string) bool {
	/*
			checks if filename includes .txt,
			if so, returns True
			else, returns false
	*/
	tail := "txt"
	for i := range filename {
		if filename[i] == '.' {
			s := strings.Split(filename, ".")[1]
			// fmt.Println(s)
			if s == tail {
				return true
			}
		}
	}
	return false
}


func filterInput(input string) string {
	/*
			Traverse through input until you reach '.', then add '.html' to the end.
			return s
	*/
	ext := ".html"
	s := strings.Split(input, ".")[0] + ext
	return s
}

func main() {
	// arg := os.Args[1] // Makesite MVP

	parser() //Makesite v1.1 + v1.2
	// renderTemplate("template.tmpl", readFile(arg))
	// writeTemplateToFile("template.tmpl", arg)
}
