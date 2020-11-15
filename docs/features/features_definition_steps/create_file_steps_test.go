package features_definition_steps

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
)

var userTitle string
var createdFilename string
var createdFilenameWithPath string

func aNewFileIsCreated(filename string) error {
	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd ../e2e/tests; ls \"%s\"", filepath.Join(directory, filename)),
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf("file was not created: %s %s", err, output)
	}
	createdFilename = filename
	createdFilenameWithPath = filepath.Join(directory, filename)
	return nil
}

func theAdrFileContentHasTheTitle(titleInContent string) error {
	titleInContent = "# " + titleInContent
	searchCommand := fmt.Sprintf(`grep -E "^# (.+)$" %s`, createdFilenameWithPath)

	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd ../e2e/tests; %s", searchCommand),
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error searching string in file: %s %s", err, output)
	}

	returned := strings.TrimSpace(string(output))
	if returned != titleInContent {
		return fmt.Errorf("expected title: \"%s\"  found: \"%s\"", titleInContent, returned)
	}

	return nil
}

func theAdrHasTheStatus(status string) error {
	searchCommand := fmt.Sprintf(`grep -E "^Status: (.+)$" %s`, createdFilenameWithPath)

	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd ../e2e/tests; %s", searchCommand),
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error searching string in file: %s %s", err, output)
	}

	returned := strings.TrimSpace(string(output))
	expected := "Status: " + status
	if returned != expected {
		return fmt.Errorf("expected status: \"%s\"  found: \"%s\"", expected, returned)
	}

	return nil
}

func theAdrHasAnId(adrId int) error {
	re := regexp.MustCompile(`(?mi)^(\d+)-.+\.md`)
	matches := re.FindStringSubmatch(createdFilename)
	if len(matches) < 2 {
		return fmt.Errorf("filename not valid %s", createdFilename)
	}

	matchId, _ := strconv.Atoi(matches[1])
	if matchId != adrId {
		return fmt.Errorf("expected status: \"%d\"  found: \"%d\"", adrId, matchId)
	}

	return nil
}

func theCreateCommandIsExecuted() error {
	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd ../e2e/tests; ../bin/adrgen create \"%s\"", userTitle),
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing the create command: %s %s", err, output)
	}

	return nil
}

func theUserSpecifyTheTitle(title string) error {
	userTitle = title
	return nil
}

func thereIsAConfigFileCreatedWithThisConfiguration(
	table *messages.PickleStepArgument_PickleTable,
) error {
	row := table.GetRows()[1]
	content := fmt.Sprintf(`
default_meta: []
default_status: %s
directory: %s
id_digit_number: %s
supported_statuses:
- proposed
- accepted
- rejected
- superseded
- amended
- deprecated
- custom
template_file: %s

`,
		row.GetCells()[0].Value,
		row.GetCells()[1].Value,
		row.GetCells()[3].Value,
		row.GetCells()[2].Value,
	)

	configFile := "adrgen.config.yml"

	output, err := exec.Command("/bin/sh", "-c", fmt.Sprintf(
		"cd ../e2e/tests; touch %s; echo \"%s\" > %s",
		configFile,
		content,
		configFile,
	)).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error generating the config file: %s %s", err, output)
	}

templateContent := `# {title}

Date: {date}

## Status

Status: {status}

## Context

What is the issue that we're seeing that is motivating this decision or change?

## Decision

What is the change that we're proposing and/or doing?

## Consequences

What becomes easier or more difficult to do because of this change?
`


	output, err = exec.Command("/bin/sh", "-c", fmt.Sprintf(
		"cd ../e2e/tests; mkdir %s; touch %s;echo \"%s\" > %s",
		row.GetCells()[1].Value,
		row.GetCells()[2].Value,
		templateContent,
		row.GetCells()[2].Value,
	)).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error generating the config file: %s %s", err, output)
	}

	directory = row.GetCells()[1].Value

	return nil
}

func thereIsNotAnyConfigFile() error {
	exec.Command("/bin/sh", "-c", "rm ../e2e/tests/adrgen.config.yml" ).CombinedOutput()
	directory = ""
	return nil
}

func CreateFeatureContext(s *godog.ScenarioContext) {
	s.Step(`^a (.+) is created$`, aNewFileIsCreated)
	s.Step(`^the adr file content has the (.+) title$`, theAdrFileContentHasTheTitle)
	s.Step(`^the adr has a (.+) status$`, theAdrHasTheStatus)
	s.Step(`^the adr has an id (\d+)$`, theAdrHasAnId)
	s.Step(`^the create command is executed$`, theCreateCommandIsExecuted)
	s.Step(`^the user specify the (.+) title$`, theUserSpecifyTheTitle)
	s.Step(
		`^there is a config file created with this configuration$`,
		thereIsAConfigFileCreatedWithThisConfiguration,
	)
	s.Step(`^there is not any config file$`, thereIsNotAnyConfigFile)
}
