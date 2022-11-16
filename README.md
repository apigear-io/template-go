# GO Blueprint

A Go blueprint.

## Features

* api: create interface files and simulation support
* scaffold: create a fully featured project with reference implementations and tests
* http: A http client to be used with the Apigear HTTP protocol.

## Development

You need to have go installed and configured.

* https://go.dev/

### Build

```bash
go run main.go
```

will print the targets available.

```bash
Targets:
  clean      removes all generated files.
  diff       runs the generator and compares the output with the golden master.
  install    installs the apigear cli.
  master     generates the golden master.
```

