package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	employees := map[string]string{
		"Subhransu": "Maharana",
		"AAAAAAA":   "aaaaaa",
		"BBBBBBB":   "bbbbbb"}
	err := tpl.Execute(os.Stdout, employees)

	if err != nil {
		log.Fatalln(err)
	}

}
