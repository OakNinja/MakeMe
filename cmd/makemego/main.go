package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"makemego/internal/makemego"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	args = reorderArgs(args)
	flags := flag.NewFlagSet("mm", flag.ContinueOnError)
	helpFlag := flags.Bool("h", false, "Show help")
	fileFlag := flags.String("f", "", "Makefile path")
	interactiveFlag := flags.Bool("i", false, "Interactive mode")
	printCommandFlag := flags.Bool("print-command", false, "Print command instead of executing")

	if err := flags.Parse(args); err != nil {
		return err
	}

	if *helpFlag {
		fmt.Fprintln(out, "Usage: mm [-h] [-f Makefile] [-i] [query]")
		return nil
	}

	var path, query string
	nonFlagArgs := flags.Args()

	if len(nonFlagArgs) > 0 {
		if _, err := os.Stat(nonFlagArgs[0]); err == nil {
			path = nonFlagArgs[0]
			if len(nonFlagArgs) > 1 {
				query = strings.Join(nonFlagArgs[1:], " ")
			}
		} else {
			query = strings.Join(nonFlagArgs, " ")
		}
	}

	if *fileFlag != "" {
		path = *fileFlag
	}

	makefile, err := makemego.FindMakefile(path)
	if err != nil {
		return err
	}

	targets, err := makemego.GetTargets(makefile)
	if err != nil {
		return fmt.Errorf("error getting targets: %w", err)
	}

	if len(targets) == 0 {
		if *printCommandFlag {
			fmt.Fprintln(os.Stderr, "No make targets found.")
			return nil
		} else {
			fmt.Fprintln(out, "No make targets found.")
			return nil
		}
	}

	// In test environment, we don't want to run fzf
	if os.Getenv("GO_TESTING") == "true" {
		return nil
	}

	selectedTarget, err := makemego.RunFzf(targets, query, *interactiveFlag, makefile)
	if err != nil {
		if err.Error() != "exit status 1" {
			return fmt.Errorf("error running fzf: %w", err)
		}
		return nil
	}

	if selectedTarget != "" {
		if *printCommandFlag {
			var command string
			if *fileFlag != "" || path != "" {
				command = fmt.Sprintf("make -f %s %s", makefile, selectedTarget)
			} else {
				command = fmt.Sprintf("make %s", selectedTarget)
			}
			fmt.Fprint(out, command)
		} else if !*interactiveFlag {
			if err := makemego.RunMake(selectedTarget, makefile); err != nil {
				return fmt.Errorf("error running make: %w", err)
			}
		}
	}

	return nil
}

func reorderArgs(args []string) []string {
	var flags []string
	var other []string
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "-") {
			// Check for -f which takes an argument
			if arg == "-f" {
				flags = append(flags, arg)
				if i+1 < len(args) {
					flags = append(flags, args[i+1])
					i++
				}
			} else {
				flags = append(flags, arg)
			}
		} else {
			other = append(other, arg)
		}
	}
	return append(flags, other...)
}
