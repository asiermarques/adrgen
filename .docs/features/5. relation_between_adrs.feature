Feature: relation link between ADR files

  In order to have a clear trace and the evolution of my architecture
  As an ADR writer
  I want to be able to manager the relations between ADR files

  Background:

    ADRGen allows creating ADR Files that supersede another one with the command

        adrgen create "My ADR file" -s 1

    ADRGen allows creating ADR Files that amends another one with the command

        adrgen create "My ADR file" -a 1

    Given there are the following adrs in the system
      | 0001-new-adr.md  | accepted    | 1   | New adr |
      | 0002-new-adr.md  | accepted    | 2   | New adr |
      | 0003-new-adr.md  | accepted    | 3   | New adr |
      | 0004-new-adr.md  | accepted    | 4   | New adr |

      And there is a config file created with this configuration
        | default_status | directory   | template_file         | id_digit_number |
        | custom         | ./          | ./template.md         | 4               |

  Scenario Outline: relations between adr files
    When the user specify the <title> title
      And the user specify the <relation> relation with the target ADR with the <targetADR> id
      And the create command is executed
    Then the <filename> ADR file is created
      And the adr has the <expectedStatus> status
      And the adr has the <relation> link on it
      And the target ADR has the <relation> relation link on it and the <expectedTargetStatus> status

    Examples:
    | title    | filename         | relation     | targetADR | expectedStatus   | expectedTargetStatus |
    | New adr  | 0005-new-adr.md  | amend        | 2         | custom           | amended              |
    | New adr  | 0006-new-adr.md  | supersede    | 2         | custom           | superseded           |
    | New adr  | 0007-new-adr.md  | amend        | 3         | custom           | amended              |
    | New adr  | 0008-new-adr.md  | amend        | 4         | custom           | amended              |



