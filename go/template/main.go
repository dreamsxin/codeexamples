package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	//"github.com/oxtoacart/bpool"
)

var templates map[string]*template.Template

//var bufpool *bpool.BufferPool

type UserData struct {
	Name        string
	City        string
	Nationality string
}

type SkillSet struct {
	Language string
	Level    string
}

type TemplateConfig struct {
	TemplateLayoutPath  string
	TemplateIncludePath string
}

type SkillSets []*SkillSet

var mainTmpl = `{{define "main" }} {{ template "base" . }} {{ end }}`

var templateConfig TemplateConfig

func loadConfiguration() {
	templateConfig.TemplateLayoutPath = "tmpl/layout/"
	templateConfig.TemplateIncludePath = "tmpl/"
}

func loadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layoutFiles, err := filepath.Glob(templateConfig.TemplateLayoutPath + "*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	includeFiles, err := filepath.Glob(templateConfig.TemplateIncludePath + "*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	mainTemplate := template.New("main")

	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			log.Fatal(err)
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}

	log.Println("templates loading successful")
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("The template %s does not exist.", name),
			http.StatusInternalServerError)
		return
	}

	//buf := bufpool.Get()
	//defer bufpool.Put(buf)
	//buf.WriteTo(w)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.tmpl", nil)
}

func main() {
	loadConfiguration()
	loadTemplates()
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)
	server.ListenAndServe()
}
