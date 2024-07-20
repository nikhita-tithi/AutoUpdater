package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
)

var currentVersion = "1.0.0" // Current version of the application

const (
    versionFilePath = "update_version.txt" // Local file containing the latest version info
    updateFilePath  = "update"             // Local file path of the new executable
)

func main() {
    // Initialize log file
    logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    defer logFile.Close()
    log.SetOutput(logFile)

    // Set up signal handling
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        handleSignals(sigs)
    }()

    // Start update checker in a separate goroutine
    go func() {
        checkForUpdates()
    }()

    // Main application logic here
    log.Println("Application is running...")
    fmt.Println("Application is running...Check app.log...")

    // Print the current date and time every 30 seconds
    ticker := time.NewTicker(3 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            log.Println("Current date and time:", time.Now().Format(time.RFC1123))
        case sig := <-sigs:
            fmt.Printf("\nReceived signal: %v. Exiting...\n", sig)
            os.Exit(0)
        }
    }
}

func handleSignals(sigs chan os.Signal) {
    for sig := range sigs {
        fmt.Printf("\nReceived signal: %v. Exiting...\n", sig)
        os.Exit(0)
    }
}

func checkForUpdates() {
    for {
        time.Sleep(12 * time.Second) // Check every 12 sec
        log.Println("Checking for updates...")
        updateAvailable, err := IsUpdateAvailable(versionFilePath, currentVersion)
        if err != nil {
            log.Printf("Failed to check for updates: %v", err)
            continue
        }

        if updateAvailable {
            log.Println("Update available. Fetching and applying update...")
            if err := CopyUpdate(updateFilePath, "new_executable"); err != nil {
                log.Printf("Failed to copy update: %v", err)
                continue
            }
            if err := ApplyUpdate("new_executable"); err != nil {
                log.Printf("Failed to apply update: %v", err)
                continue
            }
        } else {
            log.Println("No updates available.")
        }
    }
}
