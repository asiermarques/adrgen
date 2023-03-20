package main

import (
	"fmt"
	"github.com/asiermarques/adrgen/internal/_infrastructure/e2e"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"
	"os"
	"os/exec"
	"testing"
)

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress", // can define default values
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name: "adrgen e2e",
		TestSuiteInitializer: func(ts *godog.TestSuiteContext) {
			ts.BeforeSuite(func() {
				output, err := exec.Command(
					"/bin/sh",
					"-c",
					"go build -o features/e2e/bin/adrgen",
				).CombinedOutput()
				if err != nil {
					fmt.Printf("error generating the adrgen binary: %s %s", err, output)
				}
				output, err = exec.Command("/bin/sh", "-c", "cd features/e2e; mkdir tests").CombinedOutput()
				if err != nil {
					fmt.Printf("error creating the tests directory: %s %s", err, output)
				}
			})
			ts.AfterSuite(func() {
				output, err := exec.Command("/bin/sh", "-c", "rm -rf features/e2e").CombinedOutput()
				if err != nil {
					fmt.Printf("error removing test directory: %s %s", err, output)
				}
			})
		},
		ScenarioInitializer: e2e.DefinitionSteps,
		Options:             &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
