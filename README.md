# ADRgen

> work in progress


[![<ORG_NAME>](https://circleci.com/gh/asiermarques/adrgen.svg?style=svg)](https://circleci.com/gh/asiermarques/adrgen)

Another little tool for generating Architecture Decision Records


## Getting started

### Build the binary
```
go build -o adrgen
```

### Initializing the project and configuration

```
adrgen init "docs/adrs"
```

This creates a structure like this

```
your_dir
├── adrgen.config.yml
└── docs
    └── adrs
        └── adr_template.md
```

The adrgen.config.yml will be used by other commands.

It will include the following configuration:

| key                | value type | description                                                                   |
|--------------------|------------|-------------------------------------------------------------------------------|
| directory          | string     | the directory where the ADR files will be managed by adrgen                   |
| default_meta       | array      | the keys for meta that will include in all ADR files                          |
| supported_statuses | array      | the status that we will support for the ADRs                                  |
| default_status     | string     | the status that the ADR status will be set by default in the creation process |
| template_file      | string     | the template file that will be used to generate the ADR files                 |



### Create a new ADR

**Simple Usage**

```
adrgen create "My new adr"
```

**Specify meta parameters**

```
adrgen create "My new adr" -m "components, technologies"
```

The meta parameters will add a meta section on the top of the generated file.  
```
---
components: ""
technologies: ""
---

# My new adr
```

They could be useful for an automated process that uses the ADR files to generate a living documentation website.

