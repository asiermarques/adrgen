package e2e

import (
	"fmt"
	"github.com/cucumber/messages-go/v10"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var directory string

func theInitCommandIsExecuted() error {
	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd features/e2e/tests; ../bin/adrgen init \"%s\"", directory),
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing the init command: %s %s", err, output)
	}

	return nil
}

func theInitCommandIsExecutedWithOption(option string) error {
	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd features/e2e/tests; ../bin/adrgen init \"%s\" %s", directory, option),
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
		fmt.Sprintf("cd features/e2e/tests; ls \"%s\"", directory),
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
		fmt.Sprintf("cd features/e2e/tests; ls \"%s\"", location),
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
		fmt.Sprintf("cd features/e2e/tests; ls \"%s\"", file),
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

var userTitle string
var userMetaParams string
var createdFilename string
var createdFilenameWithPath string
var userStatus string
var relation string
var targetADRId int
var ADRs map[int]string = map[int]string{}
var commandOutput string

var templateContent = `# {title}

Date: {date}

## Status

{status}

## Context

What is the issue that we're seeing that is motivating this decision or change?

## Decision

What is the change that we're proposing and/or doing?

## Consequences

What becomes easier or more difficult to do because of this change?
`

func getTitleInFile(filename string) (string, error) {
	searchCommand := fmt.Sprintf(`grep -E "^# (.+)$" %s`, filename)

	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd features/e2e/tests; %s", searchCommand),
	).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error searching string in file: %s %s", err, output)
	}

	return strings.TrimSpace(string(output)), nil
}

func aNewFileIsCreated(filename string) error {
	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd features/e2e/tests; ls \"%s\"", filepath.Join(directory, filename)),
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
	returned, err := getTitleInFile(createdFilenameWithPath)
	if err != nil {
		return err
	}

	if returned != titleInContent {
		return fmt.Errorf("expected title: \"%s\"  found: \"%s\"", titleInContent, returned)
	}

	return nil
}

func getStatusInFile(status string, file string) error {
	content, err := ioutil.ReadFile(filepath.Join("features/e2e/tests", file))
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`(?mi)^## Status\n\n?(.+)$`)
	matches := re.FindStringSubmatch(string(content))
	if len(matches) < 1 {
		return fmt.Errorf("target ADR content have not a status field")
	}

	returned := strings.TrimSpace(matches[0])
	expected := "## Status\n\n" + status
	if returned != expected {
		return fmt.Errorf("expected: \"%s\"  found: \"%s\"", expected, returned)
	}

	return nil
}

func theAdrHasTheStatus(status string) error {
	return getStatusInFile(status, createdFilenameWithPath)
}

func theAdrHasAnId(adrId int) error {
	re := regexp.MustCompile(`(?mi)^(\d+)-.+\.md`)
	matches := re.FindStringSubmatch(createdFilename)
	if len(matches) < 2 {
		return fmt.Errorf("filename not valid %s", createdFilename)
	}

	matchId, _ := strconv.Atoi(matches[1])
	if matchId != adrId {
		return fmt.Errorf("expected: \"%d\"  found: \"%d\"", adrId, matchId)
	}

	return nil
}

func theCreateCommandIsExecuted() error {
	metaCommandFlag := ""
	if userMetaParams != "" {
		metaCommandFlag = fmt.Sprintf("-m \"%s\"", userMetaParams)
	}

	var relationCommand string
	if relation == "supersede" {
		relationCommand = fmt.Sprintf("-s \"%d\"", targetADRId)
	}
	if relation == "amend" {
		relationCommand = fmt.Sprintf("-a \"%d\"", targetADRId)
	}

	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf(
			"cd features/e2e/tests; ../bin/adrgen create \"%s\" %s %s",
			userTitle,
			metaCommandFlag,
			relationCommand,
		),
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
		"cd features/e2e/tests; touch %s; echo \"%s\" > %s",
		configFile,
		content,
		configFile,
	)).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error generating the config file: %s %s", err, output)
	}

	output, err = exec.Command("/bin/sh", "-c", fmt.Sprintf(
		"cd features/e2e/tests; mkdir %s; touch %s;echo \"%s\" > %s",
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
	exec.Command("/bin/sh", "-c", "rm features/e2e/tests/adrgen.config.yml").CombinedOutput()
	directory = ""
	return nil
}

func hasTheFollowingContent(content *messages.PickleStepArgument_PickleDocString) error {
	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd features/e2e/tests; cat \"%s\"", createdFilenameWithPath),
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s file not found: %s %s", createdFilenameWithPath, err, output)
	}

	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")

	expected := strings.TrimSpace(content.Content)
	expected = strings.Replace(expected, "{date}", date, 1)

	returned := strings.TrimSpace(string(output))

	if returned != expected {
		return fmt.Errorf("expected:\n%s\n\nreturned:\n%s", expected, returned)
	}

	return nil
}

func thereIsNoADRFiles() error {
	exec.Command("/bin/sh", "-c", "rm features/e2e/tests/*.md").CombinedOutput()
	return nil
}

func theMetaParametersAreSpecified(params string) error {
	userMetaParams = params
	return nil
}

func theUserSpecifyTheStatusForTheADRIdentifiedByTheId(status string, adrId int) error {
	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd features/e2e/tests; ../bin/adrgen status %d \"%s\"", adrId, status),
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing the status command: %s %s", err, output)
	}

	return nil
}

func thereIsAnADRFileWithContent(
	filename string,
	content *messages.PickleStepArgument_PickleDocString,
) error {
	output, err := exec.Command("/bin/sh", "-c", fmt.Sprintf(
		"cd features/e2e/tests; touch %s; echo \"%s\" > %s",
		filename,
		content.Content,
		filename,
	)).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error generating the adr file: %s %s", err, output)
	}

	createdFilename = filename
	createdFilenameWithPath = filename

	return nil
}

func theAdrHasTheLinkOnIt(relation string) error {
	searchCommand := fmt.Sprintf(`grep -E "^(.+)\[(.+)\]((.+))$" %s`, createdFilenameWithPath)

	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd features/e2e/tests; %s", searchCommand),
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"error searching string in file %s: %s %s",
			createdFilenameWithPath,
			err,
			output,
		)
	}

	returned := strings.TrimSpace(string(output))

	targetADRTitle, err := getTitleInFile(ADRs[targetADRId])
	targetADRTitle = strings.Replace(targetADRTitle, "# ", "", 1)
	if err != nil {
		return err
	}

	var expected string
	switch relation {
	case "supersede":
		expected = fmt.Sprintf("Supersedes [%s](%s)", targetADRTitle, ADRs[targetADRId])
	case "amend":
		expected = fmt.Sprintf("Amends [%s](%s)", targetADRTitle, ADRs[targetADRId])
	}
	if returned != expected {
		return fmt.Errorf("expected: \"%s\"  found: \"%s\"", expected, returned)
	}

	return nil
}

func weHaveACleanedSystem() error {
	_, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd features/e2e/tests; rm -f *.md")).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error cleaning the test directory: %s", err)
	}
	return nil
}

func theFollowingAdrsInTheSystem(table *messages.PickleStepArgument_PickleTable) error {
	var content string
	for _, row := range table.GetRows() {
		content = templateContent
		content = strings.Replace(content, "{status}", row.GetCells()[1].Value, 1)
		content = strings.Replace(content, "{title}", row.GetCells()[3].Value, 1)
		output, err := exec.Command(
			"/bin/sh",
			"-c",
			fmt.Sprintf(
				"cd features/e2e/tests; touch \"%s\"; echo \"%s\"> %s",
				row.GetCells()[0].Value,
				content,
				row.GetCells()[0].Value,
			),
		).CombinedOutput()
		id, _ := strconv.Atoi(row.GetCells()[2].Value)
		ADRs[id] = row.GetCells()[0].Value
		if err != nil {
			return fmt.Errorf("error creating files in system: %s %s", err, output)
		}
	}

	return nil
}

func theTargetADRHasTheLinkOnItAndThStatus(relation string, status string) error {
	searchCommand := fmt.Sprintf(`grep -E "^(.+)\[(.+)\]((.+))$" %s`, ADRs[targetADRId])

	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd features/e2e/tests; %s", searchCommand),
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf(
			"error searching rel string in file %s: %s %s",
			ADRs[targetADRId],
			err,
			output,
		)
	}

	returned := strings.TrimSpace(string(output))

	ADRTitle, err := getTitleInFile(createdFilenameWithPath)
	ADRTitle = strings.Replace(ADRTitle, "# ", "", 1)
	if err != nil {
		return err
	}

	var expected string
	switch relation {
	case "supersede":
		expected = fmt.Sprintf("Superseded by [%s](%s)", ADRTitle, createdFilenameWithPath)
	case "amend":
		expected = fmt.Sprintf("Amended by [%s](%s)", ADRTitle, createdFilenameWithPath)
	}
	if returned != expected {
		return fmt.Errorf("expected: \"%s\"  found: \"%s\"", expected, returned)
	}

	return getStatusInFile(status, ADRs[targetADRId])
}

func theUserSpecifyTheRelationWithTheTargetADRWithTheId(_relation string, targetAdrId int) error {
	relation = _relation
	targetADRId = targetAdrId
	return nil
}

func theUserExecutesTheListCommand() error {
	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd features/e2e/tests; ../bin/adrgen list")).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing the list command: %s %s", err, output)
	}
	commandOutput = strings.TrimSpace(string(output))
	return nil
}

func theUserExecutesTheListCommandWithTheFilter(filter string) error {
	output, err := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf("cd features/e2e/tests; ../bin/adrgen list -f %s", filter)).CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing the list command: %s %s", err, output)
	}
	commandOutput = strings.TrimSpace(string(output))
	return nil
}

func theUserSeeTheResultOnTheScreen(contentRaw *messages.PickleStepArgument_PickleDocString) error {
	content := strings.TrimSpace(contentRaw.Content)
	content = strings.Replace(content, "Filename", "Filename         ", 1)
	content = strings.Replace(content, ".md", ".md  ", 4)
	content = strings.TrimSpace(content)
	if strings.Contains(commandOutput, content) == false {
		return fmt.Errorf("expected: \n%s\n\nreturned: \n%s", content, commandOutput)
	}
	return nil
}
