package utils

import (
	"bytes"
	"fmt"
	"github.com/gedex/inflector"
	"github.com/serenize/snaker"
	"go/format"
	"gopkg.in/godo.v2/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func Check(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}
func FirstCharLower(s string) string {
	ss := strings.Split(s, "")
	ss[0] = strings.ToLower(ss[0])

	return strings.Join(ss, "")
}

func GetFirstChar(s string) string {
	return s[0:1]
}

var FuncMap = template.FuncMap{
	"pluralize":   inflector.Pluralize,
	"singularize": inflector.Singularize,
	"title":       strings.Title,
	"firstLower":  FirstCharLower,
	"toLower":     strings.ToLower,
	"toSnakeCase": snaker.CamelToSnake,
	"firstChar":   GetFirstChar,
}

func GenerateFile(templateFile string, outputPath string, data interface{}) {
	path := filepath.Join("templates", templateFile)
	body, _ := ioutil.ReadFile(path)
	tmpl := template.Must(template.New("model").Option("missingkey=error").Funcs(FuncMap).Parse(string(body)))

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	Check(err)

	src, _ := format.Source(buf.Bytes())
	dstPath := filepath.Join(outputPath)

	if !util.FileExists(filepath.Dir(dstPath)) {
		if err := os.Mkdir(filepath.Dir(dstPath), 0644); err != nil {
			fmt.Println(err)
		}
	}
	if err := ioutil.WriteFile(dstPath, src, 0644); err != nil {
		fmt.Println(err)
	}
}
