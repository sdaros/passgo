// NEXT: finish up the template
package cmd

import (
	"bufio"
	"flag"
	"os"
	"path/filepath"
	"text/template"
)

type Options struct {
	CommandFlagName      string
	CommandFlagHyphenName string
	CommandFlagUsage string
	CommandFlagDefaultValue  string
	Output string
}

var commandFlagTemplate string = `
package cmd

import (
        "fmt"
        "strconv"
)

type {{.CommandFlagName}} struct {
	name      string `schema.org: "/name"`
	usage     string `schema.org: "/description"`
	value     int    `schema.org: "/value"`
	isCommand bool
}

func New{{.CommandFlagName}}() *{{.CommandFlagName}} {
	return &{{.CommandFlagName}}{
		name:      "{{.CommandFlagHyphenName}}",
		usage:     "{{.CommandFlagUsage}}",
		value:     {{.CommandFlagDefaultValue}},
		isCommand: false,
	}
}

func (self *{{.CommandFlagName}}) Name() string {
	return self.name
}

func (self *{{.CommandFlagName}}) Usage() string {
	return self.usage
}

func (self *{{.CommandFlagName}}) IsCommand() bool {
	return self.isCommand
}

func (self *{{.CommandFlagName}}) String() string {
	return fmt.Sprint(self.value)
}

func (self *{{.CommandFlagName}}) Set(value string) (err error) {
	r, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	if err := self.Validate(r); err != nil {
		return err
	}
	self.value = r
	return nil
}

func (self *{{.CommandFlagName}}) Validate(r int) (err error) {
	return nil
}
`

func main() {
	options := parseFlags()
	tmpl, err := template.New("method").Parse(CommandFlagTemplate)
	if err != nil {
		panic(err)
	}
	createDirectoryIfNotExist(options.Output)
	f, err := os.Create(options.Output)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(f)
	defer func() {
		writer.Flush()
		f.Close()
	}()

	err = tmpl.Execute(writer, options)
	if err != nil {
		panic(err)
	}
}

func createDirectoryIfNotExist(file string) {
	directory := filepath.Dir(file)
	_, err := os.Stat(directory)
	if os.IsNotExist(err) {
		os.MkdirAll(directory, 0777)
	}
}

func parseFlags() Options {
	output := flag.String("output", "", "Output file")
	methodName := flag.String("method", "", "Method name")
	apiEndpoint := flag.String("endpoint", "", "Endpoint of API")
	outputType := flag.String("type", "", "Output Type")
	flag.Parse()
	return Options{
		Output:      *output,
		MethodName:  *methodName,
		ApiEndpoint: *apiEndpoint,
		OutputType:  *outputType,
	}
}
