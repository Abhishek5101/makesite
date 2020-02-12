package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type entry struct {
	Paragraph string
}

func readFile(name string) string {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(fileContents)

}

func writeFile(name string, data string) {
	bytesToWrite := []byte(data)
	err := ioutil.WriteFile(name, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func writeTranslate(filename string, lang string) {
	FileText := readFile(filename)

	contents, error := TranslateText(lang, FileText)
	if error != nil {
		panic(error)
	}
	bytesToWrite := []byte(contents)

	err := ioutil.WriteFile(filename, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func templateRenderer(filename string, data interface{}) {
	t := template.Must(template.New(filename).ParseFiles(filename))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func txtToHTML(input string) string {
	fileName := strings.TrimSuffix(input, filepath.Ext(input)) + ".html"
	return fileName
}

func templateWriter(lang string, templateName string, fileName string) {

	textParagraph := entry{Paragraph: readFile(fileName)}
	t := template.Must(template.New("template.tmpl").ParseFiles(templateName))

	file, err := os.Create(txtToHTML(fileName))
	if err != nil {
		panic(err)
	}

	err = t.Execute(file, textParagraph)
	if err != nil {
		panic(err)
	}
}

func directoryParser() {
	var directory string
	flag.StringVar(&directory, "dir", ".", "This is the directory.")

	var lang string
	flag.StringVar(&lang, "lang", "mr", "This is the language you want to translate, inputting google's language abbreviations.")
	flag.Parse()

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			// fmt.Printf("[%t] %s\n", file.IsDir(), file.Name())
		} else {
			textFile := ".txt"
			if filepath.Ext(strings.TrimSpace(file.Name())) == textFile {
				fmt.Printf(" %s\n", file.Name())
				writeTranslate(file.Name(), lang)
				templateWriter(lang, "template.tmpl", file.Name())
			}
		}
	}
}

func main() {
	directoryParser()
}
