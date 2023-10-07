Feature: create ADR files

  In order to have registered context about architecture decisions
  As an ADR writer
  I want to execute a command that register the decision in an ADR file

  Background:

    ADRGen allows creating ADR Files with the command

        adrgen create "My ADR file"

    This creates an markdown file taking setting a new id automatically, by calculating the last id
    from all the ADR files in the specified directory in the configuration or if there is no configuration
    taking the current directory instead.


  Scenario Outline: create adr files

    Given there is a config file created with this configuration
      | default_status | directory       | template_file              | id_digit_number |
      | custom         | ./adrs          | ./adrs/template.md         | 2               |

    When the user specify the <title> title
      And the create command is executed
    Then the <filename> ADR file is created
      And the adr has an id <id>
      And the adr has the <status> status
      And the adr file content has the <title_in_file> title

    Examples:
    | title             | filename       | id  | status    | title_in_file |
    | New adr           | 01-new-adr.md  | 1   | custom    | 1. New adr    |
    | New adr           | 02-new-adr.md  | 2   | custom    | 2. New adr    |
    | New adr           | 03-new-adr.md  | 3   | custom    | 3. New adr    |
    | New adr           | 04-new-adr.md  | 4   | custom    | 4. New adr    |


  Scenario Outline: create adr files without configuration

    Given there is not any config file

    When the user specify the <title> title
      And the create command is executed
    Then the <filename> ADR file is created
      And the adr has the <status> status
      And the adr file content has the <title_in_file> title

    Examples:
      | title             | filename         | status    | title_in_file |
      | New adr           | 0001-new-adr.md  | proposed  | 1. New adr    |
      | New adr           | 0002-new-adr.md  | proposed  | 2. New adr    |

  Scenario Outline: create adr files with asciidoc template

    When the user specify the adr-asciidoc directory
      And the init command is executed with option --asciidoc
      And the user specify the <title> title
      And the create command is executed
    Then the <filename> ADR file is created
      And the adr has an id <id>
      And the adr has the <status> status
      And the adr file content has the <title_in_file> title

    Examples:
    | title             | filename           | id  | status     | title_in_file |
    | New adr           | 0001-new-adr.adoc  | 1   | proposed   | 1. New adr    |
    | New adr           | 0002-new-adr.adoc  | 2   | proposed   | 2. New adr    |
    | New adr           | 0003-new-adr.adoc  | 3   | proposed   | 3. New adr    |
    | New adr           | 0004-new-adr.adoc  | 4   | proposed   | 4. New adr    |