
# SuiteKit-Go

A Go-based command-line tool for Suite operations, also known as `suitectl`.

## Description

`suitectl` is a command-line utility that provides functionalities related to the "Suite" operations. This tool is built using Go.

## Build

To build and install the project, follow these steps:

```bash
git clone [Your Repo URL]
cd suite-kit-go
BUILD_DATE=$(TZ=Asia/Shanghai date +"%Y-%m-%d_%H:%M:%S_%Z")
CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-extldflags '-static' -X suite-kit-go/cmd.buildDate=$BUILD_DATE" -o suitectl main.go
```

## Usage

After building, you can use the `suitectl` command:

```bash
./suitectl [command] [options]
```

### Available Commands:

- `silent-install`: Deploy itsma suite service.

For more details on commands and options, refer to the documentation or use:

```bash
./suitectl [command] --help
```

## Update go version 
- Update `go x.xx` in the `go.mod`
- Update the `GO_VERSION="x.xx.x"` in the `Build.sh`

## Contributing

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -am 'Add some YourFeature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Update the `%changelog` in `suitekit.rpm.spec`.
6. Open a new Pull Request.
   
> Please kindly add comments during programming.
