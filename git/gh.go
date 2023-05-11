package git

import (
	"os/exec"
	"strings"
)

func GetPrCommits() ([]string, error) {
	cmd := exec.Command(
		"gh",
		"pr",
		"view",
		"--json commits",
		"--jq",
		`'.commits[]|"- \(.messageHeadline)\(.messageBody)"'`,
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return []string{}, err
	}

	s := strings.Split(strings.TrimSpace(string(out)), "\n\n")

	return s, err
}
