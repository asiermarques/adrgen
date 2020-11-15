package features_definition_steps

import (
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
)

func aNewadrmdIsCreated(arg1 int) error {
	return godog.ErrPending
}

func theAdrFileContentHasTheNewAdrTitle(arg1 int) error {
	return godog.ErrPending
}

func theAdrHasAProposedStatus() error {
	return godog.ErrPending
}

func theAdrHasAnId(arg1 int) error {
	return godog.ErrPending
}

func theCommandIsExecuted() error {
	return godog.ErrPending
}

func theUserSpecifyTheNewAdrTitle() error {
	return godog.ErrPending
}

func thereIsAConfigFileCreatedWithThisConfiguration(arg1 *messages.PickleStepArgument_PickleTable) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a (\d+)-new-adr\.md is created$`, aNewadrmdIsCreated)
	s.Step(`^the adr file content has the (\d+)\. New adr title$`, theAdrFileContentHasTheNewAdrTitle)
	s.Step(`^the adr has a proposed status$`, theAdrHasAProposedStatus)
	s.Step(`^the adr has an id (\d+)$`, theAdrHasAnId)
	s.Step(`^the command is executed$`, theCommandIsExecuted)
	s.Step(`^the user specify the New adr title$`, theUserSpecifyTheNewAdrTitle)
	s.Step(`^there is a config file created with this configuration$`, thereIsAConfigFileCreatedWithThisConfiguration)
}
