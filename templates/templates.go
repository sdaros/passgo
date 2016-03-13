package main

import (
	"bufio"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	templates = make(map[string]string)
)

type Options struct {
	Name         string
	LcName       string
	HyphenName   string
	Usage        string
	DefaultValue string
	ShortName    string
	Output       string
	Template     string
}

func main() {
	options := parseFlags()
	tmpl, err := template.New("method").Parse(options.Template)
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
	template := flag.String("template", "", "template type to generate")
	output := flag.String("output", "", "Output file")
	name := flag.String("name", "", "Command Flag Name")
	hyphenName := flag.String("hyphen-name", "", "Command Flag Name with hyphens for registrar")
	usage := flag.String("usage", "", "Short description of usage for the command flag")
	defaultValue := flag.String("default-value", "", "Default value for the command flag")
	shortName := flag.String("short-name", "", "short name for flag")
	flag.Parse()
	return Options{
		Name:         toPublicName(name),
		LcName:       toPrivateName(name),
		HyphenName:   *hyphenName,
		Usage:        *usage,
		DefaultValue: *defaultValue,
		ShortName:    *shortName,
		Output:       *output,
		Template:     templates[*template],
	}
}

// toPrivateName changes the first letter of name to lowercase.
func toPrivateName(value *string) string {
	var privateName []string
	name := strings.Split(*value, "")
	privateName = append(privateName, strings.ToLower(name[0]))
	privateName = append(privateName, name[1:]...)
	return strings.Join(privateName, "")
}

// toPublicName changes the first letter of name to uppercase.
func toPublicName(value *string) string {
	var publicName []string
	name := strings.Split(*value, "")
	publicName = append(publicName, strings.ToUpper(name[0]))
	publicName = append(publicName, name[1:]...)
	return strings.Join(publicName, "")
}

func init() {
	templates["newCommandTemplate"] = newCommandTemplate
	templates["newCommandAsFlagTemplate"] = newCommandAsFlagTemplate
	templates["newFlagTemplate"] = newFlagTemplate
}
