package main

import (
	"fmt"
	"log"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
)

func selectLicense(licenses []LicenseData) string {
	items := []string{}

	for i := range licenses {
		items = append(items, licenses[i][0])
	}

	prompt := promptui.Select{
		HideHelp: true,
		Items:    items,
		Label:    "Select license",
	}

	result, _, err := prompt.Run()

	if err != nil {
		log.Println(aurora.Red("not selected"))
		os.Exit(-1)
	}

	return licenses[result][1]
}

func promptYear() string {
	year := getCurrentYear()
	prompt := promptui.Prompt{
		Label: fmt.Sprintf("Year (%d)", year),
	}

	result2, err := prompt.Run()

	if err != nil {
		log.Println(aurora.Red("No data provided"))
		return fmt.Sprintf("%d", year)
	}

	if len(result2) < 1 {
		result2 = fmt.Sprintf("%d", year)
	}

	return result2
}

func promptFullname() string {
	fullname := getGitUsername()
	prompt := promptui.Prompt{
		Label: "Full name (" + fullname + ")",
	}

	result, err := prompt.Run()

	if err != nil {
		log.Println(aurora.Red("No data provided"))
		return fullname
	}

	if len(result) < 1 {
		result = fullname
	}

	return result
}

func promptFilename() string {
	prompt := promptui.Prompt{
		Label: "File name (LICENSE)",
	}

	result, err := prompt.Run()

	if err != nil {
		log.Println(aurora.Red("No data provided"))
		return "LICENSE"
	}

	if len(result) < 1 {
		result = "LICENSE"
	}

	return result
}
