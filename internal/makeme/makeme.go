package makeme

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	fzf "github.com/junegunn/fzf/src"
)

// execCommand allows mocking exec.Command in tests
var execCommand = exec.Command

func FindMakefile(file string) (string, error) {
	if file != "" {
		info, err := os.Stat(file)
		if err != nil {
			return "", fmt.Errorf("path not found at %s", file)
		}

		if info.IsDir() {
			makefileNames := []string{"GNUmakefile", "Makefile", "makefile"}
			for _, name := range makefileNames {
				path := file + "/" + name
				if _, err := os.Stat(path); err == nil {
					return path, nil
				}
			}
			return "", fmt.Errorf("No Makefile found in %s", file)
		} else {
			return file, nil
		}
	}

	makefileNames := []string{"GNUmakefile", "Makefile", "makefile"}
	for _, name := range makefileNames {
		if _, err := os.Stat(name); err == nil {
			return name, nil
		}
	}

	return "", fmt.Errorf("No Makefile found")
}

func GetTargets(makefile string) ([]string, error) {
	contentBytes, err := os.ReadFile(makefile)
	if err != nil {
		return nil, err
	}
	content := string(contentBytes)

	parsedMakefile, err := parseMakefile(makefile)
	if err != nil {
		return nil, err
	}

	var staticTargets, fileTargets, generatedTargets []string

	sort.Strings(parsedMakefile)

	for _, row := range parsedMakefile {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}

		if strings.Contains(row, ".") || strings.Contains(row, "/") {
			fileTargets = append(fileTargets, row)
		} else {
			if strings.Contains(content, row+":") {
				staticTargets = append(staticTargets, row)
			} else {
				generatedTargets = append(generatedTargets, row)
			}
		}
	}

	return append(staticTargets, append(fileTargets, generatedTargets...)...), nil
}

func parseMakefile(filename string) ([]string, error) {
	cmd := execCommand("make", "--version")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var command string
	if strings.Contains(string(out), "GNU") {
		gnu_awk1 := `'{if (sub(/\\$/,"")) printf "%%s", $0; else print $0}'`
		gnu_awk2 := `'/^# Files/,/^# Finished Make data base/ {if ($1 == "# Not a target") skip = 1; if ($1 !~ "^[#.\t]") { if (!skip) print $1; skip=0 }}'`
		command = fmt.Sprintf("make -f %s -pRrq : 2>/dev/null | awk %s | awk -F: %s 2>/dev/null", filename, gnu_awk1, gnu_awk2)
	} else {
		bsd_awk := `'/^#\\*\\*\\* Input graph:/,/^$/ {if ($1 !~ "^#... ") {gsub(/# /, "", $1); print $1}}'`
		command = fmt.Sprintf("make -f %s -d g1 -rn >/dev/null 2>| awk -F, %s 2>/dev/null", filename, bsd_awk)
	}

	cmd = execCommand("bash", "-c", command)
	cmd.Env = append(os.Environ(), "LC_ALL=C")
	parsedOutput, err := cmd.Output()
	if err != nil {
		return []string{}, nil
	}

	return strings.Split(string(parsedOutput), "\n"), nil
}

func RunFzf(targets []string, query string, interactive bool, makefile string) (string, error) {
	inputChan := make(chan string)
	go func() {
		for _, target := range targets {
			inputChan <- target
		}
		close(inputChan)
	}()

	outputChan := make(chan string)
	var selectedTarget string
	go func() {
		for s := range outputChan {
			selectedTarget = s // Assuming single selection, capture the last one.
		}
	}()

	var fzfOptions []string
	fzfOptions = append(fzfOptions, "--height", "60%")
	fzfOptions = append(fzfOptions, "--layout=reverse")
	fzfOptions = append(fzfOptions, "--border")
	fzfOptions = append(fzfOptions, "--preview-window=right:60%")
	previewArg := fmt.Sprintf("grep --color=always -A 10 -B 1 ^{}: %s; or echo -GENERATED TARGET-", makefile)
	fzfOptions = append(fzfOptions, "--preview", previewArg)

	if query != "" {
		fzfOptions = append(fzfOptions, fmt.Sprintf("--query=%s", query))
	}

	if interactive {
		makeCommand := fmt.Sprintf("make -f %s", makefile)
		bindArg := fmt.Sprintf("enter:execute:%s {}; echo; echo Done; sleep 1", makeCommand)
		fzfOptions = append(fzfOptions, "--bind", bindArg)
	}

	options, err := fzf.ParseOptions(
		true,
		fzfOptions,
	)
	if err != nil {
		return "", err
	}

	options.Input = inputChan
	options.Output = outputChan

	_, err = fzf.Run(options)
	if err != nil {
		return "", err
	}

	return selectedTarget, nil
}

func RunMake(target string, makefile string) error {
	cmd := execCommand("make", "-f", makefile, target)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
