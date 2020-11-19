Feature: create an ADR File with meta

  In order to have better integration with other tools
  As an ADR writer
  I want to add meta information to the ADR files


  Background:

    ADRGen allows creating meta parameters in ADR Files with the command

      adrgen create "My ADR file" -m "param1, param2"

    This command creates the ADR file with a Markdown meta section on the top of the file and adds the
    specified params to it, for example:

      ---
      param1: ""
      param2: ""
      ---

      # My ADR file


  Scenario: create an ADR file with meta params

    Given there is not any config file
      And there is no ADR files
    When the user specify the My architecture decision title
      And the meta parameters param1, param2 are specified
      And the create command is executed
    Then the 0001-my-architecture-decision.md ADR file is created
      And has the following content:
      """
      ---
      param1: ""
      param2: ""
      ---

      # 1. My architecture decision

      Date: {date}

      ## Status

      Status: proposed

      ## Context

      What is the issue that we're seeing that is motivating this decision or change?

      ## Decision

      What is the change that we're proposing and/or doing?

      ## Consequences

      What becomes easier or more difficult to do because of this change?
      """


