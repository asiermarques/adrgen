# ADRgen

> work in progress


[![<ORG_NAME>](https://circleci.com/gh/asiermarques/adrgen.svg?style=svg)](https://circleci.com/gh/asiermarques/adrgen)

Another little tool for generating Architecture Decision Records


## Getting started

### Build the binary
```
go build -o adrgen
```

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

They could be useful for an automated process that use the ADR files to generate living documentation.

