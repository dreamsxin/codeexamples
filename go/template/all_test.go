package main

import (
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	// template action
	// {{template "templ_name" }}
	// {{template "templ_name" pipeline}}
	// Here, the template with name “templ_name” will be executed and output will be rendered at the same place.
	var templateDemo = `
	{{ define "a" }} 
	Template A includes Template B
	{{ template "b" .}}
	Template A ends
	{{ end }}
	{{define "b"}} 
	Template B
	{{end}}
	`

	var err error

	tt := template.New("templateActionDemo")

	tt, err = tt.Parse(templateDemo)

	if err != nil {
		t.Fatalf("parsing failed: %s", err)
	}
	err = tt.ExecuteTemplate(os.Stdout, "b", nil)
	if err != nil {
		t.Fatalf("execution failed: %s", err)
	}
	err = tt.ExecuteTemplate(os.Stdout, "a", nil)
	if err != nil {
		t.Fatalf("execution failed: %s", err)
	}

}

func TestBlock(t *testing.T) {
	// block action
	// Block action is defining a template and executing in place.
	var templateDemo = `
	{{ define "a" }} 
	Template A includes Template B
	{{ block "b" .}} Template B {{ end }}
	Template A ends
	{{ end }}
	`
	var err error

	tmpl := template.New("templateActionDemo")

	tmpl, err = tmpl.Parse(templateDemo)
	if err != nil {
		t.Fatal("parse failed:", err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "b", nil)
	if err != nil {
		t.Fatalf("execution failed: %s", err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "a", nil)
	if err != nil {
		t.Fatalf("execution failed: %s", err)
	}

}
