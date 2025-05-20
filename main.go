package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	file := filepath.Base(os.Args[0])

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get home directory: %v", err)
	}

	workingDir := filepath.Join(homeDir)
	if err := os.MkdirAll(workingDir, 0755); err != nil {
		log.Fatalf("Failed to create working directory: %v", err)
	}

	logPath := filepath.Join(workingDir, "workingon.log")

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags)

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Usage: %s log [message]\n", file)
		fmt.Printf("       %s ls | list\n", file)
		fmt.Printf("       %s last\n", file)
		fmt.Printf("       %s path\n", file)
		os.Exit(1)
	}

	if args[0] == "ls" || args[0] == "list" {
		data, err := os.ReadFile(logPath)
		if err != nil {
			log.Fatalf("Failed to read log file: %v", err)
		}
		fmt.Print(string(data))
		return
	}

	if args[0] == "last" {
		data, err := os.ReadFile(logPath)
		if err != nil {
			log.Fatalf("Failed to read log file: %v", err)
		}
		lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")
		if len(lines) > 0 && lines[0] != "" {
			fmt.Println(lines[len(lines)-1])
		}
		return
	}

	if args[0] == "path" {
		fmt.Println(logPath)
		return
	}

	if args[0] != "log" {
		fmt.Printf("Usage: %s log [message]\n       %s ls\n       %s last\n       %s path\n", file, file, file, file)
		os.Exit(1)
	}

	// Read from stdin if available and append to args
	text := strings.Join(args[1:], " ")

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		stdinBytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Failed to read from stdin: %v", err)
		}
		stdinText := strings.TrimSpace(string(stdinBytes))
		if stdinText != "" {
			// args = append(args, stdinText)
			text = fmt.Sprintf("%s %s", stdinText, text)
		}
	}

	shortHash := ""
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	cmd.Dir, _ = os.Getwd()
	output, err := cmd.Output()
	if err == nil {
		shortHash = fmt.Sprintf("[git:%s]", strings.TrimSpace(string(output)))
	}

	cwd, _ := os.Getwd()
	logMsg := fmt.Sprintf("\"%s\" %s (%s)", text, shortHash, cwd)
	log.Println(logMsg)

	red := "\033[31m"
	green := "\033[32m"
	reset := "\033[0m"
	coloredMsg := fmt.Sprintf("\"%s\" %s%s%s (%s%s%s)", text, red, shortHash, reset, green, cwd, reset)
	fmt.Println(coloredMsg)
}
