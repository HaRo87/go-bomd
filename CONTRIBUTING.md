# Contributing

Contributions are welcome, and they are greatly appreciated!
Every little bit helps, and credit will always be given.

## Environment setup

Nothing easier!

Fork and clone the repository, then:

```bash
cd go-bomd
task setup
```

!!! note
    If it fails for some reason,
    you can check your setup via `task check-setup`.

You now have the dependencies installed.

Getting a list of all available tasks can be achieved by running

```bash
task --list
```

which should provide an output similar to:

```bash
* build:                       Build the project
* check-setup:                 check system dependencies
* clean:                       Clean the project
* download-dependencies:       downloading all required Go dependencies
* format:                      Format the project
* install-tools:               installing required tools
* lint:                        Lint the code
* security:                    Run gosec for project
* setup:                       installing all required project dependencies
* test:                        Test the project
```

## Development

As usual:

1. create a new branch: `git checkout -b feature/\#<issue-number>---feature-name`
2. edit the code and/or the documentation

**Before committing:**

1. run `task format` to auto-format the code
2. run `task lint` to lint everything (fix any warning)
3. run `task test` to run the tests (fix any issue)
4. run `task security` to run gosec (fix any issue)

If you are unsure about how to fix or ignore a warning,
just let the continuous integration fail,
and we will help you during review.
