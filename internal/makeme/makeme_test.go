package makeme

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// TestHelperProcess isn't a real test. It's used as a helper process
// for mocking exec.Command.
func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No command\n")
		os.Exit(2)
	}

	cmd, args := args[0], args[1:]
	switch cmd {
	case "make":
		if len(args) > 0 && args[0] == "--version" {
			fmt.Println("GNU Make 4.3")
			os.Exit(0)
		}
		// Mock RunMake
		if len(args) >= 2 && args[0] == "-f" {
			// args: -f makefile target
			target := args[len(args)-1]
			fmt.Printf("Running target: %s\n", target)
			os.Exit(0)
		}
	case "bash":
		// Mock parseMakefile
		// The command passed to bash -c is complex, we just want to return some targets
		fmt.Println("target1")
		fmt.Println("target2")
		os.Exit(0)
	}
	fmt.Fprintf(os.Stderr, "Unknown command %q\n", cmd)
	os.Exit(2)
}

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestRunMake(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	// Capture stdout to verify output
	// Note: RunMake writes to os.Stdout directly, which is hard to capture in parallel tests
	// without pipe redirection. For simplicity in this unit test, we trust the mock runs.
	// But wait, RunMake sets cmd.Stdout = os.Stdout.
	// We can't easily capture os.Stdout in a pure unit test without replacing os.Stdout.
	// However, we can verify it doesn't error.

	err := RunMake("build", "Makefile")
	if err != nil {
		t.Errorf("RunMake() error = %v, want nil", err)
	}
}

func TestParseMakefile(t *testing.T) {
	execCommand = fakeExecCommand
	defer func() { execCommand = exec.Command }()

	// We need to set the env var in the current process so it propagates
	// to the subprocess via os.Environ() in parseMakefile.
	os.Setenv("GO_WANT_HELPER_PROCESS", "1")
	defer os.Unsetenv("GO_WANT_HELPER_PROCESS")

	targets, err := parseMakefile("Makefile")
	if err != nil {
		t.Fatalf("parseMakefile() error = %v, want nil", err)
	}

	// Filter out "PASS" and other go test noise
	var cleanTargets []string
	for _, tgt := range targets {
		if tgt != "" && !strings.HasPrefix(tgt, "PASS") && !strings.HasPrefix(tgt, "ok") {
			cleanTargets = append(cleanTargets, tgt)
		}
	}
	targets = cleanTargets

	if len(targets) != 2 {
		t.Errorf("parseMakefile() returned %d targets: %v, want 2", len(targets), targets)
	}
	if len(targets) > 0 && targets[0] != "target1" {
		t.Errorf("expected target1, got %s", targets[0])
	}
}

func TestBuildFzfCommand(t *testing.T) {
	// This function doesn't use exec, so we can test it directly.
	cmd := buildFzfCommand("Makefile", true, "make", "build")
	if !strings.Contains(cmd, "fzf") {
		t.Errorf("Expected fzf command, got %s", cmd)
	}
	if !strings.Contains(cmd, "--query=build") {
		t.Errorf("Expected query arg, got %s", cmd)
	}
}
