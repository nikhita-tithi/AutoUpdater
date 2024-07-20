package main

import (
    "fmt"
    "io"
    "log"
    "os"
    "os/exec"
    "runtime"
    "strings"
)

// IsUpdateAvailable checks if an update is available by reading the version from a local file.
func IsUpdateAvailable(versionFilePath, currentVersion string) (bool, error) {
    log.Printf("Opening version file: %s", versionFilePath)
    file, err := os.Open(versionFilePath)
    if err != nil {
        return false, fmt.Errorf("failed to open version file: %v", err)
    }
    defer file.Close()

    var latestVersion string
    log.Println("Reading version file...")
    _, err = fmt.Fscanf(file, "%s", &latestVersion)
    if err != nil {
        return false, fmt.Errorf("failed to read version file: %v", err)
    }

    log.Printf("Current version: %s, Latest version: %s", currentVersion, latestVersion)

    return compareVersions(latestVersion, currentVersion), nil
}

// CopyUpdate copies the new executable to the specified path.
func CopyUpdate(sourcePath, destinationPath string) error {
    log.Printf("Opening source file: %s", sourcePath)
    srcFile, err := os.Open(sourcePath)
    if err != nil {
        return fmt.Errorf("failed to open source file: %v", err)
    }
    defer srcFile.Close()

    log.Printf("Creating destination file: %s", destinationPath)
    dstFile, err := os.Create(destinationPath)
    if err != nil {
        return fmt.Errorf("failed to create destination file: %v", err)
    }
    defer dstFile.Close()

    log.Printf("Copying data from %s to %s", sourcePath, destinationPath)
    _, err = io.Copy(dstFile, srcFile)
    if err != nil {
        return fmt.Errorf("failed to copy file: %v", err)
    }

    // Set execute permissions on the new executable (Unix-based systems)
    if runtime.GOOS != "windows" {
        log.Printf("Setting execute permissions on %s", destinationPath)
        err = os.Chmod(destinationPath, 0755)
        if err != nil {
            return fmt.Errorf("failed to set execute permissions: %v", err)
        }
    }

    log.Println("Update copied successfully.")
    return nil
}

// ApplyUpdate replaces the current executable with the new one and restarts the application.
func ApplyUpdate(newExecutablePath string) error {
    currentExecutablePath, err := os.Executable()
    if err != nil {
        return fmt.Errorf("failed to get current executable path: %v", err)
    }

    backupPath := currentExecutablePath + ".bak"

    // Rename current executable to backup
    log.Println("Renaming current executable to backup...")
    err = os.Rename(currentExecutablePath, backupPath)
    if err != nil {
        return fmt.Errorf("failed to rename current executable: %v", err)
    }

    // Replace current executable with new one
    log.Println("Replacing current executable with new one...")
    err = os.Rename(newExecutablePath, currentExecutablePath)
    if err != nil {
        return fmt.Errorf("failed to rename new executable: %v", err)
    }

    // Restart the application
    log.Println("Restarting application...")
    cmd := exec.Command(currentExecutablePath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err = cmd.Start()
    if err != nil {
        return fmt.Errorf("failed to restart application: %v", err)
    }

    // Exit the current process
    log.Println("Exiting current process...")
    os.Exit(0)
    return nil
}

// Compare two version strings
func compareVersions(new_version, old_version string) bool {
    return versionToInt(new_version) > versionToInt(old_version)
}

// Convert version string to integer for comparison
func versionToInt(version string) int {
    parts := strings.Split(version, ".")
    var versionInt int
    for _, part := range parts {
        num := 0
        fmt.Sscanf(part, "%d", &num)
        versionInt = versionInt*100 + num // assuming no more than 99 for each part
        // You can modify this calculation based on your versioning scheme
    }
    return versionInt
}