package features_definition_steps

import (
	"fmt"
	"github.com/cucumber/godog"
	"os/exec"
)

func theInitCommandIsExecuted() error {
	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd ../e2e/tests; ../bin/adrgen init \"%s\"", directory),
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing the init command: %s %s", err, output)
	}

	return nil
}

func theSpecifiedDirectoryIsCreated() error {
	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd ../e2e/tests; ls \"%s\"", directory),
	).CombinedOutput()

	if err != nil {
		return fmt.Errorf("directory was not created: %s %s", err, output)
	}

	return nil
}

func theTemplateFileIsCreatedInTheLocation(location string) error {
	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd ../e2e/tests; ls \"%s\"", location),
	).CombinedOutput()

	if err != nil {
		return fmt.Errorf("file was not created: %s %s", err, output)
	}

	return nil
}

func aConfigFileIsCreated(file string) error {
	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd ../e2e/tests; ls \"%s\"", file),
	).CombinedOutput()

	if err != nil {
		return fmt.Errorf("file was not created: %s %s", err, output)
	}

	return nil
}

func theUserIsInAnInitialDirectory() error {
	return nil
}

func theUserSpecifyTheDirectory(_directory string) error {
	directory = _directory
	return nil
}

func InitFeatureContext(s *godog.ScenarioContext) {
	s.Step(`^the (.+) config file is created$`, aConfigFileIsCreated)
	s.Step(`^the init command is executed$`, theInitCommandIsExecuted)
	s.Step(`^the specified directory is created$`, theSpecifiedDirectoryIsCreated)
	s.Step(`^the template file is created in the (.+) location$`, theTemplateFileIsCreatedInTheLocation)
	s.Step(`^the user is in an initial directory$`, theUserIsInAnInitialDirectory)
	s.Step(`^the user specify the (.+) directory$`, theUserSpecifyTheDirectory)
}
