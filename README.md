# Introduction

`go-bomd` is a CLI which can be used to transform software dependency
information from a [CycloneDX](https://cyclonedx.org) Software Bill Of 
Materials - SBOM - file into other formats based on Golang's template
engine. It is a replacement for [mdbom](https://github.com/HaRo87/mdbom).

# Getting started

## Development

Make sure to have the following tools/languages installed:

- [Go](https://go.dev/dl/) -> 1.19
- [Task](https://taskfile.dev)

You can get a list of available tasks by running `task --list`
which should produce an output similar to:

```bash
task: Available tasks for this project:
* build:          Build the project
* clean:          Clean the project
* format:         Format the project
* security:       Run gosec for project
* test:           Test the project
```

After building the CLI via `task build` you should be able to
study the available commands via `./build/go-bomd --help` which
should produce an output similar to:

```bash
go-bomd can read in Software Bill Of Materials (SBOMs)
        based on the CycloneDX standard and convert relevant information
        into markdown based documents using custom templates.

Usage:
  bomd [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Generate a specified item
  help        Help about any command
  validate    Validate a specified item

Flags:
  -c, --config string   config file (default ./config.yml) (default "config.yml")
  -f, --file string     the file on which an operation should be performed
  -h, --help            help for bomd
      --ignore-errors   do not error out
  -v, --verbose count   logger verbosity

Use "bomd [command] --help" for more information about a command.
```
