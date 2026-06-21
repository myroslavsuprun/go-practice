package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func main() {
	printMoney(false)
	var reader = bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("something has gone wrong")
			os.Exit(1)
		}

		var inputs = strings.Split(strings.ReplaceAll(input, "\n", ""), " ")
		var command = inputs[0]
		var args = inputs[1:]
		processCommand(command, args)
		printMoney(true)
	}
}

func processCommand(command string, args []string) {
	switch command {
	case "exit":
		os.Exit(0)
	case "echo":
		echoing(args)
	case "type":
		typing(args)
	default:
		if cmdPath, ok := buildCmdExecPath(command); ok {
			var _, err = exec.LookPath(cmdPath)
			if err != nil {
				commandNotFound(command)
				return
			}

			var cmd = exec.Command(command, args...)

			var output, _ = cmd.CombinedOutput()
			var before, _ = strings.CutSuffix(string(output), "\n")
			fmt.Printf("%v", before)

			return
		}

		commandNotFound(command)
	}
}

func buildCmdExecPath(command string) (string, bool) {
	if hasPrefixMultiple(command, []string{"./", "../"}) {
		return command, true
	}

	if path, exists := lookupCommandInPaths(command); exists {
		return fmt.Sprintf("%v/%v", path, command), true
	}

	return "", false
}

func hasPrefixMultiple(s string, prefix []string) bool {
	for _, p := range prefix {
		if strings.HasPrefix(s, p) {
			return true
		}
	}

	return false
}

var builtinCommands = []string{"type", "echo", "exit"}

func typing(args []string) {
	var command = strings.Join(args, "\n")

	if slices.Contains(builtinCommands, command) {
		fmt.Printf("%v is a shell builtin", command)
		return
	}

	var path, exists = lookupCommandInPaths(command)
	if exists {
		fmt.Printf("%v is %v/%v", command, path, command)
		return
	}

	fmt.Printf("%v: not found", command)
}

func lookupCommandInPaths(command string) (string, bool) {
	var path, exists = os.LookupEnv("PATH")
	if !exists {
		fmt.Printf("%v: not found", command)
		return "", false
	}

	var paths = strings.Split(path, ":")
	return hasCommandInPaths(command, paths)
}

func hasCommandInPaths(command string, paths []string) (string, bool) {
	for _, p := range paths {
		var entries, err = os.ReadDir(p)
		if err != nil {
			continue
		}
		hasPerm, err := hasXPerm(command, entries)
		if err != nil {
			return "", false
		}

		if hasPerm {
			return p, true
		}
	}

	return "", false
}

func hasXPerm(command string, entries []os.DirEntry) (bool, error) {
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if entry.Name() == command {
			var info, err = entry.Info()
			if err != nil {
				return false, err
			}

			var mode = info.Mode().String()
			if mode[3] == 'x' {
				return true, nil
			}
		}
	}

	return false, nil
}

func echoing(args []string) {
	fmt.Printf("%v", strings.Join(args, " "))
}

func commandNotFound(command string) {
	fmt.Printf("%v: command not found", command)
}

func printMoney(newLine bool) {
	var newLineVal string = ""

	if newLine {
		newLineVal = "\n"
	}

	fmt.Printf("%v$ ", newLineVal)
}
