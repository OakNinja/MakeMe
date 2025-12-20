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
	content, err := os.ReadFile(makefile)
	if err != nil {
		return nil, err
	}
	return GetTargetsFromContent(string(content))
}

func GetTargetsFromContent(content string) ([]string, error) {
	tmpfile, err := os.CreateTemp("", "Makefile.test")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(content); err != nil {
		return nil, err
	}
	if err := tmpfile.Close(); err != nil {
		return nil, err
	}

	parsedMakefile, err := parseMakefile(tmpfile.Name())
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

func buildFzfCommand(filename string, interactive bool, make_command string, query string) string {
	var fzf_opts []string
	fzf_opts = append(fzf_opts, "--read0")

	if query != "" {
		fzf_opts = append(fzf_opts, fmt.Sprintf("--query=%s", query))
	}

	if interactive {
		bind_arg := fmt.Sprintf("--bind=enter:execute:%s {}; echo; echo Done; sleep 1", make_command)
		fzf_opts = append(fzf_opts, bind_arg)
	}

	fzf_opts = append(fzf_opts, "--height", "60%")
	fzf_opts = append(fzf_opts, "--layout=reverse")
	fzf_opts = append(fzf_opts, "--border")
	fzf_opts = append(fzf_opts, "--preview-window=right:60%")
	preview_arg := fmt.Sprintf("--preview='grep --color=always -A 10 -B 1 ^^{}: %s; or echo -GENERATED TARGET-'", filename)
	fzf_opts = append(fzf_opts, preview_arg)

	fzf_tmux := os.Getenv("FZF_TMUX")
	if fzf_tmux == "" {
		fzf_tmux = "0"
	}
	fzf_tmux_height := os.Getenv("FZF_TMUX_HEIGHT")
	if fzf_tmux_height == "" {
		fzf_tmux_height = "60%"
	}

	var command string
	opts_str := strings.Join(fzf_opts, " ")
	if fzf_tmux == "1" {
		command = fmt.Sprintf("fzf-tmux -d%s %s", fzf_tmux_height, opts_str)
	} else {
		command = fmt.Sprintf("fzf %s", opts_str)
	}
	return command
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

	// The library doesn't return the selection directly in the same way.
	// We need to capture the output from the output channel.
	// Wait, looking at the example again:
	// outputChan := make(chan string)
	// go func() { for s := range outputChan { fmt.Println("Got: " + s) } }()
	// The output channel receives the selected items.

	// So I need to capture it.
	return selectedTarget, nil
}

func RunMake(target string, makefile string) error {
	cmd := execCommand("make", "-f", makefile, target)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
