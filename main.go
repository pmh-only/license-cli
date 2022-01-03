package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
)

func init() {
	log.SetOutput(colorable.NewColorableStdout())
	log.SetFlags(0)
}

func main() {
	licenses := getLicenseData()
	chosenLicense := selectLicense(licenses)
	content := getLicenseFile(chosenLicense)

	if strings.Contains(content, "{{year}}") {
		year := promptYear()
		content = strings.Replace(content, "{{year}}", year, -1)
	}

	if strings.Contains(content, "{{fullname}}") {
		fullname := promptFullname()
		content = strings.Replace(content, "{{fullname}}", fullname, -1)
	}

	filename := promptFilename()

	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(aurora.Green("License file created"))
}
