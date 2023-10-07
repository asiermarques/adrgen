package e2e

import "github.com/cucumber/godog"

func DefinitionSteps(s *godog.ScenarioContext) {
	s.Step(`^the (.+) config file is created$`, aConfigFileIsCreated)
	s.Step(`^the init command is executed$`, theInitCommandIsExecuted)
	s.Step(`^the init command is executed with option (.+)$`, theInitCommandIsExecutedWithOption)
	s.Step(`^the specified directory is created$`, theSpecifiedDirectoryIsCreated)
	s.Step(
		`^the template file is created in the (.+) location$`,
		theTemplateFileIsCreatedInTheLocation,
	)
	s.Step(`^the user is in an initial directory$`, theUserIsInAnInitialDirectory)
	s.Step(`^the user specify the (.+) directory$`, theUserSpecifyTheDirectory)
	s.Step(`^the (.+) ADR file is created$`, aNewFileIsCreated)
	s.Step(`^the adr file content has the (.+) title$`, theAdrFileContentHasTheTitle)
	s.Step(`^the meta parameters (.+) are specified$`, theMetaParametersAreSpecified)
	s.Step(`^the adr has the (.+) status$`, theAdrHasTheStatus)
	s.Step(`^the adr has an id (\d+)$`, theAdrHasAnId)
	s.Step(`^the create command is executed$`, theCreateCommandIsExecuted)
	s.Step(`^the user specify the (.+) title$`, theUserSpecifyTheTitle)
	s.Step(`^has the following content:$`, hasTheFollowingContent)
	s.Step(
		`^there is a config file created with this configuration$`,
		thereIsAConfigFileCreatedWithThisConfiguration,
	)
	s.Step(`^there is not any config file$`, thereIsNotAnyConfigFile)
	s.Step(`^there is no ADR files$`, thereIsNoADRFiles)
	s.Step(
		`^the user executes the status command specifying (.+) for the ADR identified by the (\d+) id$`,
		theUserSpecifyTheStatusForTheADRIdentifiedByTheId,
	)
	s.Step(`^there is a (.+) ADR file with the following content:$`, thereIsAnADRFileWithContent)

	s.Step(`^the adr has the (.+) link on it$`, theAdrHasTheLinkOnIt)
	s.Step(`^there are the following adrs in the system$`, theFollowingAdrsInTheSystem)
	s.Step(
		`^the target ADR has the (.+) relation link on it and the (.+) status$`,
		theTargetADRHasTheLinkOnItAndThStatus,
	)
	s.Step(
		`^the user specify the (.+) relation with the target ADR with the (\d+) id$`,
		theUserSpecifyTheRelationWithTheTargetADRWithTheId,
	)
	s.Step(`^the user executes the list command$`, theUserExecutesTheListCommand)
	s.Step(`^the user see the result on the screen:$`, theUserSeeTheResultOnTheScreen)
	s.Step(`^the user executes the list command with the filter "([^"]*)"$`, theUserExecutesTheListCommandWithTheFilter)
	s.Step(`^we have a cleaned system$`, weHaveACleanedSystem)

}
