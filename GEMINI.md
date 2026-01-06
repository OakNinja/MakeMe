# Project Overview

`MakeMe` is a command-line tool written in Go that helps you interactively select and run `make` targets from a Makefile. It is a cross-platform replacement for the original `MakeMeFish` script.

# Building and Running

- **Build:** To build the executable, run the following command:
  ```
  go build ./cmd/makeme
  ```
- **Run:** To run the application, execute the following command in a directory containing a Makefile:
  ```
  ./makeme
  ```

# Installation

## Via Script

To install `MakeMe` and the optional shell integration, run the following command:

```bash
./install.sh
```

## Via Homebrew (macOS and Linux)

To install `MakeMe` via Homebrew, use the custom tap:

```bash
brew tap OakNinja/tap
brew install mm
```

To install the latest development version:
```bash
brew install OakNinja/tap/mm --HEAD
```


# Shell Integration

The `install.sh` script will offer to install shell integration for you. This will add a shell function `mm` that, when used, will place the selected `make` command directly into your command line, ready for you to execute.

After running the script, you will need to restart your shell or source the configuration file to apply the changes, for example:

```bash
source ~/.bashrc
```

## Usage

Once installed, you can use the `mm` command:

```bash
mm
```

After you select a target, the `make` command will appear in your prompt.

# Development Conventions

- **Project Structure:** The project follows the standard Go project structure.
- **Testing:** Tests are located in the `internal/makeme` directory and can be run with the following command:
  ```
  go test ./...
  ```