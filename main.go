package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var list = `
{{$files := array "default.yaml" (concat "default-" .Location ".yaml")}}
{{$basepath := path "/etc/puppet/environment" .Context "enc" "puppet"}}
{{range .Puppetclasses}}{{$workpath := path $basepath (convert .)}}{{range $files }}{{path $workpath .}}
{{end}}{{end}}`

var fmap = template.FuncMap{
	"array":   array,
	"concat":  concat,
	"path":    filepath.Join,
	"convert": convert,
}

func main() {
	t, err := template.New("filelist").Funcs(fmap).Parse(list)
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		Context       string
		Location      string
		Puppetclasses []string
	}{
		"test",
		"eh",
		[]string{"os::autofs", "db::oracle", "tools::patrol"},
	}

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

func array(s ...string) []string {
	return append([]string{}, s...)
}

func concat(s ...string) string {
	return strings.Join(append([]string{}, s...), "")
}

func convert(s string) string {
	return strings.Replace(s, "::", "/", -1)
}
