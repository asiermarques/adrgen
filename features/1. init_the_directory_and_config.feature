Feature: init the directory and config

  In order to be able to customize the aspects of the ADR file creation
  As an ADR writer
  I want to execute a command that creates the initial config file and template

  Background:

    ADRGen allows to create an initial configuration with the command

        adrgen init "./target_dir"


  Scenario Outline: initialize the configuration and directories

    Given the user is in an initial directory
    When the user specify the <target_directory> directory
      And the init command is executed
    Then the specified directory is created
      And the template file is created in the <template_created> location
      And the <config_file> config file is created

    Examples:

    | target_directory       | template_created                | config_file          |
    | ./target               | ./target/adr_template.md        | ./adrgen.config.yaml |
    | ./target/level         | ./target/level/adr_template.md  | ./adrgen.config.yaml |

