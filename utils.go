package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
)

func getGitUsername() string {
	cmd := exec.Command("git", "config", "user.name")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Println(aurora.Yellow("get git username failed"))
		return ""
	}

	return strings.Trim(out.String(), "\n")
}

func getCurrentYear() int {
	return time.Now().Year()
}
