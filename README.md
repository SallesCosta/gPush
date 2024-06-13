```markdown
# gPush

gPush is a CLI tool written in Go, designed to manage and push to two different Git repositories simultaneously. It is useful when you have a project that is mirrored in two different repositories.

## Installation

To install gPush, you need to have Go installed on your machine. Once Go is installed, you can download and install gPush using the `go get` command:

```bash
go get github.com/sallescosta/gPush
```

## Basic Usage

The basic usage of gPush is as follows:

```bash
gPush [flags]
```

### Flags

- `-s, --settings`: Open the settings menu. In the settings menu, you can set the original and secondary repositories.

### Commands

- `settings`: This command is used to push or set the original and secondary repositories.

## Settings

In the settings menu, you can set the original and secondary repositories. If the repositories are not set, the program will prompt you to set them.

## Behavior

When you run gPush without any flags, it will check if there are any uncommitted changes in your repositories. If there are, it will prompt you to commit them before running the program.

If the original repository is not found, the program will notify you and exit.

If the secondary repository is not found or is undefined, the program will notify you and open the settings menu for you to set it.

Once both repositories are set and there are no uncommitted changes, the program will display the original and secondary repositories and proceed to push to both of them.

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

gPush is licensed under the MIT license. See the LICENSE file for more details.
```
