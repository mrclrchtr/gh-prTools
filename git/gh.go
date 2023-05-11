package git

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func GetPrCommits() (commits []string, err error) {
	cmd := exec.Command(
		"gh",
		"pr",
		"view",
		"--json",
		"commits",
		"--jq",
		`.commits[]|"- \(.messageHeadline)\(.messageBody)"`,
	)

	out, err := executeWithLogging(cmd)
	if err != nil {
		return []string{}, err
	}

	s := strings.Split(strings.TrimSpace(out.String()), "\n\n")

	for _, commitMessage := range s {
		message := strings.Replace(commitMessage, "……", "", 1)
		commits = append(commits, message)
	}

	return commits, err
}

func executeWithLogging(cmd *exec.Cmd) (out bytes.Buffer, err error) {
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(stderr.String())
	}
	return out, err
}
