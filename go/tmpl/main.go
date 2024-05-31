package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

type Name struct {
	FirstName string
	LastName  string
}

type Nums struct {
	X int
	Y int
}

func main() {
	tmpl := `I'm {{ toUpper .FirstName }} {{ replace .LastName "D" "J" | toUpper }}`

	funcMaps := template.FuncMap{
		"toUpper": strings.ToUpper,
		"replace": strings.ReplaceAll,
	}

	t, err := template.New("tmpl").Funcs(funcMaps).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(os.Stdout, Name{FirstName: "John", LastName: "Doe"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()

	tmpl = "   {{- .X -}} > {{- .Y }}"
	t, err = t.New("newtmpl").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(os.Stdout, Nums{X: 10, Y: 12})
	fmt.Println()
}
