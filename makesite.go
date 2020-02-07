package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"path/filepath"
)

type entry struct {
	paragraph string
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

func templateRenderer(filename string, data interface{}) {
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))
	err := t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}

func txtToHTML(input string) string {
	fileName := strings.TrimSuffix(input, filepath.Ext(input)) + ".html"
	return fileName
}

func templateWriter(templateName string, fileName string) {

	textParagraph := entry{readFile(fileName)}
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

func main() {
	cmdFileName := os.Args[1]                  
	templateRenderer("template.tmpl", readFile(cmdFileName)) 
	templateWriter("template.tmpl", cmdFileName)

}
