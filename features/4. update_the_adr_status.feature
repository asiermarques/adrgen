Feature: update the ADR status

  In order easily maintain the ADR files
  As an ADR writer
  I want to change the status of an ADR file from the command line


  Background:

    ADRGen allows update the ADR status with the command

      adrgen status 1 "accepted"


  Scenario: update the ADR status without config to a default valid status

    Given there is a 0001-my-architecture-decision.md ADR file with the following content:
            """
            ---
            param1: ""
            param2: ""
            ---

            # 1. My architecture decision

            Date: {date}

            ## Status

            proposed

            ## Context

            What is the issue that we're seeing that is motivating this decision or change?

            ## Decision

            What is the change that we're proposing and/or doing?

            ## Consequences

            What becomes easier or more difficult to do because of this change?
            """
    When the user executes the status command specifying "accepted" for the ADR identified by the 1 id
    Then the adr has the accepted status

  Scenario: update the ADR status for an asciidoc decision

    Given there is a 0001-my-architecture-decision.adoc ADR file with the following content:
            """
            ---
            param1: ""
            param2: ""
            ---

            = 1. My architecture decision

            Date: {date}

            == Status

            proposed

            == Context

            What is the issue that we're seeing that is motivating this decision or change?

            == Decision

            What is the change that we're proposing and/or doing?

            == Consequences

            What becomes easier or more difficult to do because of this change?
            """
    When the user executes the status command specifying "accepted" for the ADR identified by the 1 id
    Then the adr has the accepted status

