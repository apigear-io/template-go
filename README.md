# ApiGear template for Go

This is a template for the [ApiGear](https://apigear.io) code generator.

## Features

* api: create interface files and simulation support
* scaffold: create a fully featured project with reference implementations and tests
* http: A http client to be used with the Apigear HTTP protocol.

## Prerequisites

* [Task](https://taskfile.dev/#/installation)
* [Go](https://go.dev/)

See the current tasks:

```bash
task --list
```

## Test

```bash
task test
```

This will generate a test output in the `test` folder.

To diff the test output with the golden master:

```bash
task diff
```

## License

[MIT](LICENSE)
