# MakeMe

> Easing the usage of make and Makefiles.
> A cross-platform Go port of [MakeMeFish](https://github.com/OakNinja/MakeMeFish).

MakeMe simplifies the usage of Makefiles by providing quick navigation and searching through make targets.

## Features

- **Cross-platform** - Works on macOS, Linux, and Windows (Shell integration varies)
- **Type ahead searching** - just write a few characters to filter out the targets you are looking for
- **Preview** - When selecting a target, an excerpt of the target will be shown in the makefile, with match highlighting
- **Snappy** - Go fast!

## Installation

### Automatic Installation

Run the installation script to build the binary and set up shell integration:

```bash
# For Bash and Zsh
./install.sh

# For Fish
./install.fish
```

### Manual Shell Integration

If you prefer to set up shell integration manually, source the appropriate script in your shell configuration.

**Zsh (`~/.zshrc`):**
```zsh
source /path/to/MakeMe/misc/mm.zsh
```
*Note: In Zsh, the selected command is placed on the command line buffer for editing.*

**Bash (`~/.bashrc`):**
```bash
source /path/to/MakeMe/misc/mm.bash
```
*Note: In Bash, the selected command is executed immediately and added to history.*

**Fish (`~/.config/fish/config.fish`):**
```fish
source /path/to/MakeMe/misc/mm.fish
```

### Via Homebrew (macOS)

```bash
# Assuming you have the formula available locally
brew install --build-from-source ./Formula/mm.rb
```

## Usage

_Basic usage:_
type `mm`, if there is a Makefile in the current working directory, all targets will be listed. Start typing to filter targets.

```bash
mm
```

_Parameters:_

- `-h` or `--help` to print the help.
- `-f <filename>` to specify a makefile if you have several in the cwd, or if you have a non-standard name.
- `-i` to start MakeMe in interactive mode. In interactive mode, the selected target will be executed and you will then be returned to the selection prompt. Please note that executed commands won't be added to your command history.
- `<foo>` eg. add an arbitrary keyword to start MakeMe with a pre-populated query (editable at runtime)

## Examples

`mm build` will start `MakeMe` with an initial query which will filter for targets containing the substring `build`.
Similarly, `mm foo bar` will filter on targets containing both `foo` and `bar`

---

`mm -f MyFancyMakeFile` will start `MakeMe` and parse the file `MyFancyMakeFile` instead of trying to find a makefile with a GNU make standard name.

---

`mm -i` will run `MakeMe` in interactive mode

---

_All flags and parameters can be combined, and added in any order, eg._

`mm foo -i -f MyFancyMakeFile` is equivalent to `mm -f MyFancyMakeFile foo -i`

## Development

### Testing

```bash
# Run Go tests
go test ./...

# Run installation tests
./tests/test_install.sh
```
