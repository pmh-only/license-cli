package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/logrusorgru/aurora"
)

const endpoint = "https://raw.github.com/pmh-only/license-cli/master/licenses"

type LicenseData [2]string

func getLicenseData() []LicenseData {
	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)

	spin.Suffix = " - loading list of licenses..."
	spin.Start()

	res, err := http.Get(endpoint + "/list.json")

	if err != nil {
		log.Println(aurora.Red("loading failed - is github down?"), err)
		os.Exit(-1)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(aurora.Red("loading failed - is github down?"), err)
		os.Exit(-1)
	}

	var metadata []LicenseData

	err = json.Unmarshal(body, &metadata)
	if err != nil {
		log.Println(aurora.Red("loading failed - see: https://github.com/pmh-only/license-cli/issues"), err)
		os.Exit(-1)
	}

	spin.Stop()
	return metadata
}

func getLicenseFile(filename string) string {
	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	spin.Suffix = " - loading license contents..."
	spin.Start()

	res, err := http.Get("https://raw.githubusercontent.com/pmh-only/license-cli/master/licenses/" + filename)

	if err != nil {
		log.Println(aurora.Red("loading failed - is github down?"), err)
		os.Exit(-1)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(aurora.Red("loading failed - is github down?"), err)
		os.Exit(-1)
	}

	content := string(body)
	spin.Stop()

	if strings.Compare(content, "404: Not Found") == 0 {
		log.Println(aurora.Red("loading failed - is github down?"), content)
		os.Exit(-1)
	}

	return content
}
