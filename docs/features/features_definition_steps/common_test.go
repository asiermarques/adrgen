package features_definition_steps

import (
	"fmt"
	"os/exec"

	"github.com/cucumber/godog"
)

var directory string

func SuiteContext(t *godog.TestSuiteContext) {
	t.BeforeSuite(func() {
		output, err := exec.Command(
			"/bin/sh",
			"-c",
			"cd ../../../; go build -o docs/features/e2e/bin/adrgen",
		).CombinedOutput()
		if err != nil {
			fmt.Printf("error generating the adrgen binary: %s %s", err, output)
		}
		output, err = exec.Command("/bin/sh", "-c", "cd ../e2e; mkdir tests").CombinedOutput()
		if err != nil {
			fmt.Printf("error creating the tests directory: %s %s", err, output)
		}
	})

	t.AfterSuite(func() {
		output, err := exec.Command("/bin/sh", "-c", "rm -rf ../e2e").CombinedOutput()
		if err != nil {
			fmt.Printf("error removing test directory: %s %s", err, output)
		}
	})
}
