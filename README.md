# ADRgen

[![<ORG_NAME>](https://circleci.com/gh/asiermarques/adrgen.svg?style=svg)](https://circleci.com/gh/asiermarques/adrgen) [![Go Report Card](https://goreportcard.com/badge/github.com/asiermarques/adrgen)](https://goreportcard.com/report/github.com/asiermarques/adrgen)

A command-line tool for generating and managing Architecture Decision Records

## Why to use this tool for managing your ADRs?

At a difference with other ADR file creation tools, this one has been designed for allowing the customization to
create new ADR files while providing the best user experience for maintaining the created ADR files.

Features that this tool currently offers:

* Create ADR files
* Configure custom statuses: this allows you to design your ADR workflow
* Configure a custom Markdown template file
* Change the status of an ADR file, validating its status if we have configured custom statuses
* Amend and Supersede relationships between ADR files
* Adding meta attributes in the ADR files

Next features:

* History log
* History trace exportable to a json file
* Reports


## Getting started

### Download the binary
You can download the last release binary from the [releases section](https://github.com/asiermarques/adrgen/releases/latest)


### Show the version of the ADRgen

You can show the version of the ADRgen with the **version** command

```
adrgen version
```

### Initializing the project and configuration

We will use the **init** command specifying where the ADR files will be written.

```
adrgen init "docs/adrs"
```

This command creates the following structure:

```
your_dir
├── adrgen.config.yaml
└── docs
    └── adrs
        └── adr_template.md
```

As the result, we can see
* A config file is created
* A directory structure is created if it doesn't exist
* A markdown template is created in the desired directory


The adrgen.config.yaml config file will be used by other commands in order to know how to operate with the ADR files.

It will include the following configuration keys:

| key                | type       | description                                                                   |
|--------------------|------------|-------------------------------------------------------------------------------|
| directory          | string     | the directory where the ADR files will be managed by adrgen                   |
| default_meta       | array      | the keys for meta that will include in all ADR files                          |
| supported_statuses | array      | the statuses that will be supported for the ADRs                              |
| default_status     | string     | the status that the ADR status will be set by default in the creation process |
| template_file      | string     | the template file that will be used to generate the ADR files                 |
| id_digit_number    | int        | the number of digits for the ADR identifier, for example: "0001-"             |

Example:

````
default_meta: []
default_status: proposed
directory: docs/adrs
supported_statuses:
- proposed
- accepted
- rejected
- superseeded
- amended
- deprecated
template_file: docs/adrs/adr_template.md
id_digit_number: 4
````

### Create a new ADR

**Simple Usage**

```
adrgen create "My new a"
```

**Specify meta parameters**

Sometimes, adding meta variables are useful in order to make it simpler to integrate the ADR Files with other proceses.  
For example, we can add the meta parameters to generate a static generated documention website based on the Gatsby CMS system.

```
adrgen create "My new a" -m "components, technologies"
```

The meta parameters will add a meta section on the top of the generated file.
```
---
components: ""
technologies: ""
---

# My new ADR
```

**Supersede an ADR with another new ADR**

If you make a decision that supersedes another previous one, a good practice is specify the relation in both ADR files.

You can do it automatically adding the **-s** flag to the **create** command

```
adrgen create "Another a that supersedes the previous one" -s 1
```

This adds the relation in both files

```
# My new a

Date: 11-11-2020

## Status

Status: superseded

Superseded by [Another a that supersedes the previous one](0002-another-a-that-supersedes-the-previous-one.md)
```

```
# Another a that supersedes the previous one

Date: 13-11-2020

## Status

Status: proposed

Supersedes [My new a](0001-my-new-a.md)
```

**Amend an ADR with another new ADR**

If you make a decision that improves another previous one, a good practice is specify the relation in both ADR files.

You can do it automatically adding the **-a** flag to the **create** command

```
adrgen create "Another a that improves the previous one" -a 1
```

This adds the relation in both files

```
# My new a

Date: 11-11-2020

## Status

Status: amended

Amended by [Another a that improves the previous one](0002-another-a-that-improves-the-previous-one.md)
```

```
# Another a that improves the previous one

Date: 13-11-2020

## Status

Status: proposed

Amends [My new a](0001-my-new-a.md)
```

### Update the status of an ADR File

We can change the status for an ADR file specifying its ID and the new status with the **status** command

```
adrgen status 9 "accepted"
```

### List the ADR files

You can list the ADR files with the **list** command

```
adrgen list
```

**Filter the results by status**

The results from the **list** command can be filtered by status

```
adrgen list -f status=accepted
```

You can also filter results that match one of a list of specified statuses

```
adrgen list -f status=accepted&status=proposed
```



### To-Do

You can see the next planned features in the [ADRgen project's board](https://github.com/asiermarques/adrgen/projects/1)