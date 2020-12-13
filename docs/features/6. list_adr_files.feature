Feature: list ADR files

  In order to see easily the ADRs in the working directory
  As an ADR writer
  I want to be able to list the adr files

  Background:

    ADRGen allows listing the ADR files with this command

        adrgen list

    Given we have a cleaned system
      And there are the following adrs in the system
      | 0001-new-adr.md  | accepted    | 1   | New adr |
      | 0002-new-adr.md  | proposed    | 2   | New adr |
      | 0003-new-adr.md  | accepted    | 3   | New adr |
      | 0004-new-adr.md  | accepted    | 4   | New adr |

      And there is a config file created with this configuration
        | default_status | directory   | template_file         | id_digit_number |
        | custom         | ./          | ./template.md         | 4               |

  Scenario: listing the adr files
    When the user executes the list command
    Then the user see the result on the screen:
    """
    Title    Status    Date    ID  Filename
    New adr  accepted  {date}  1   0001-new-adr.md
    New adr  proposed  {date}  2   0002-new-adr.md
    New adr  accepted  {date}  3   0003-new-adr.md
    New adr  accepted  {date}  4   0004-new-adr.md
    """

  Scenario: filter the list of the adr files
    When the user executes the list command with the filter "status=proposed"
    Then the user see the result on the screen:
    """
    Title    Status    Date    ID  Filename
    New adr  proposed  {date}  2   0002-new-adr.md
    """

