package features_definition_steps

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var userTitle string
var createdFilename string

func aNewFileIsCreated(filename string) error {
	output, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd ../e2e/tests; ls \"%s\"",filename)).CombinedOutput()
	if err != nil {
		return fmt.Errorf("file was not created: %s %s", err, output)
	}

	createdFilename = filename
	return nil
}

func theAdrFileContentHasTheTitle(titleInContent string) error {
	titleInContent = "# " + titleInContent
	searchCommand := fmt.Sprintf(`grep -E "^# (.+)$" %s`, createdFilename)

	output, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd ../e2e/tests; %s", searchCommand)).CombinedOutput()
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
	searchCommand := fmt.Sprintf(`grep -E "^Status: (.+)$" %s`, createdFilename)

	output, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd ../e2e/tests; %s", searchCommand)).CombinedOutput()
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

func theCommandIsExecuted() error {
	output, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd ../e2e/tests; ../bin/adrgen create \"%s\"",userTitle)).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing the create command: %s %s", err, output)
	}

	return nil
}

func theUserSpecifyTheTitle(title string) error {
	userTitle = title
	return nil
}

func thereIsAConfigFileCreatedWithThisConfiguration(table *messages.PickleStepArgument_PickleTable) error {
	content := `default_meta: []
default_status: proposed
directory: docs
id_digit_number: 4
supported_statuses:
	- proposed
	- accepted
	- rejected
	- superseded
	- amended
	- deprecated
template_file: docs/adr_template.md
`
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

	return nil
}

func FeatureContext(s *godog.ScenarioContext) {
	s.Step(`^a (.+) is created$`, aNewFileIsCreated)
	s.Step(`^the adr file content has the (.+) title$`, theAdrFileContentHasTheTitle)
	s.Step(`^the adr has a (.+) status$`, theAdrHasTheStatus)
	s.Step(`^the adr has an id (\d+)$`, theAdrHasAnId)
	s.Step(`^the command is executed$`, theCommandIsExecuted)
	s.Step(`^the user specify the (.+) title$`, theUserSpecifyTheTitle)
	s.Step(`^there is a config file created with this configuration$`, thereIsAConfigFileCreatedWithThisConfiguration)
}
