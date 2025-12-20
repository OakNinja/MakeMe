# Project Overview

`MakeMe` is a command-line tool written in Go that helps you interactively select and run `make` targets from a Makefile. It is a cross-platform replacement for the original `MakeMeFish` script.

# Building and Running

- **Build:** To build the executable, run the following command:
  ```
  go build ./cmd/makemego
  ```
- **Run:** To run the application, execute the following command in a directory containing a Makefile:
  ```
  ./makemego
  ```

# Installation

## Via Script

To install `MakeMe` and the optional shell integration, run the following command:

```bash
./install.sh
```

## Via Homebrew (macOS)

To install `MakeMe` with Homebrew, you will first need to publish the formula.

1.  **Push to GitHub:** Create a new repository on GitHub and push your project to it.
2.  **Create a Release:** On the GitHub repository page, create a new release (e.g., `v0.1.0`). This will generate a source code archive (`.tar.gz` file).
3.  **Update the Formula:**
    *   Open `Formula/mm.rb`.
    *   Update the `homepage` with the URL of your GitHub repository.
    *   Update the `url` with the URL of the `.tar.gz` archive from your release.
    *   Download the `.tar.gz` archive and calculate its SHA256 checksum. You can do this with the command `shasum -a 256 /path/to/your/archive.tar.gz`.
    *   Update the `sha256` in the formula with the checksum you calculated.
4.  **Install Locally:** You can then install your application on your machine with the command:
    ```
    brew install --build-from-source ./Formula/mm.rb
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
- **Testing:** Tests are located in the `internal/makemego` directory and can be run with the following command:
  ```
  go test ./...
  ```