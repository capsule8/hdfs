package main

import (
	"bufio"
	"bytes"
	"flag"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

type data struct {
	Package        string
	InitStatements string
}

var initRegex = regexp.MustCompile(`(?s)func init\(\) {(.*?)}`)
var packageRegex = regexp.MustCompile(`package (.+)`)

func extractInit(data []byte) ([]string, error) {
	initStatements := []string{}
	for _, match := range initRegex.FindAllSubmatch([]byte(data), -1) {
		scanner := bufio.NewScanner(strings.NewReader(strings.TrimSpace(string(match[1])))) // f is the *os.File
		for scanner.Scan() {
			initStatements = append(initStatements, strings.TrimSpace(scanner.Text()))
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}
	return initStatements, nil
}

func extractPackage(data []byte) string {
	matches := packageRegex.FindAllSubmatch([]byte(data), 1)
	if len(matches) != 1 {
		return ""
	}
	return string(matches[0][1])
}

func deleteInits(fn string, data []byte) error {
	data = initRegex.ReplaceAll([]byte(data), []byte{})
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_TRUNC, 0)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}
	formatted, err := format.Source(data)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(formatted)
	return err
}

func parseFiles(files string) (string, []string, error) {
	fns, err := filepath.Glob(files)
	if err != nil {
		return "", nil, err
	}
	packageName := ""
	initStatements := []string{}
	for _, fn := range fns {
		data, err := ioutil.ReadFile(fn)
		if err != nil {
			return packageName, nil, err
		}
		if packageName == "" {
			packageName = extractPackage(data)
		}
		lines, err := extractInit(data)
		if err != nil {
			return packageName, nil, err
		}
		initStatements = append(initStatements, lines...)
		err = deleteInits(fn, data)
		if err != nil {
			return packageName, nil, err
		}
	}
	return packageName, initStatements, nil
}

func main() {
	var files, output string
	flag.StringVar(&files, "files", "", "The files to remove init mentods from")
	flag.StringVar(&output, "output", "", "The file to add a LoadAPI method to")
	flag.Parse()

	out, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	packageName, inits, err := parseFiles(files)
	if err != nil {
		panic(err)
	}
	d := data{
		Package:        packageName,
		InitStatements: "    " + strings.Join(inits, "\n    "),
	}
	buf := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("loadAPI").Parse(loadAPITemplate))
	t.Execute(buf, d)
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	_, err = out.Write(formatted)
	if err != nil {
		panic(err)
	}
}

var loadAPITemplate = `
package {{.Package}}

import (
	"sync"

	proto "github.com/golang/protobuf/proto"
)

var loadOnce sync.Once

func LoadAPI() {
	loadOnce.Do(func(){
{{.InitStatements}}
    })
}
`
