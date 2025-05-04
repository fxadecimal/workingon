package workon

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

// VERSION is set at build time using -ldflags "-X main.VERSION=..."
var VERSION string

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

    logPath := filepath.Join(workingDir, "workon.log")

    logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    defer logFile.Close()

    log.SetOutput(logFile)
    log.SetFlags(log.LstdFlags)

    args := os.Args[1:]
    if len(args) == 0 || args[0] != "log" {
        fmt.Printf("Usage: %s log [message]\n", file)
        fmt.Printf("Version: %s\n", VERSION)
        os.Exit(1)
    }

    text := strings.Join(args[1:], " ")

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
    fmt.Println(logMsg) // Also print to stdout, similar to StreamHandler
}