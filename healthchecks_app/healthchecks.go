package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "strconv"
    "strings"
)

func main() {
    // Log file setup
    logFile, err := os.OpenFile("/var/log/healthcheck.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    defer logFile.Close()
    log.SetOutput(logFile)

    // Perform health checks
    if err := performHealthChecks(); err != nil {
        log.Printf("Health check failed: %v", err)

        // Send an email notification (You can uncomment this part when needed)
        // subject := "Server Health Check Failed"
        // message := "One or more health checks have failed on the server. Please investigate."
        // recipient := "your.email@example.com"
        // smtpServer := "smtp.example.com"
        // smtpPort := "587"
        // smtpUsername := "your.smtp.username"
        // smtpPassword := "your.smtp.password"
        //
        // if err := sendEmail(subject, message, recipient, smtpServer, smtpPort, smtpUsername, smtpPassword); err != nil {
        //     log.Printf("Failed to send email notification: %v", err)
        // }

        // Exit with an error code
        os.Exit(1)
    }

    fmt.Println("Server health is OK")
}

func performHealthChecks() error {
    // Check disk space
    cmd := exec.Command("df", "-h", "/")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("Error checking disk space: %v", err)
    }

    // Split the output into lines and skip the header
    lines := strings.Split(string(output), "\n")[1:]
    if len(lines) == 0 {
        return fmt.Errorf("No disk space information found")
    }

    // Parse the disk usage percentage (use the root directory as an example)
    fields := strings.Fields(lines[0])
    if len(fields) < 5 {
        return fmt.Errorf("Unexpected disk space information format")
    }

    // Extract the percentage as a string (e.g., "75%")
    usageString := fields[4]

    // Remove the "%" character
    usageString = strings.TrimSuffix(usageString, "%")

    // Parse the percentage as an integer
    usagePercent, err := strconv.Atoi(usageString)
    if err != nil {
        return fmt.Errorf("Error parsing disk usage percentage: %v", err)
    }

    // Log the disk space usage
    log.Printf("Disk Space Usage: %d%%", usagePercent)

    // Set the threshold for disk usage (e.g., 90%)
    threshold := 90

    // Check if disk usage exceeds the threshold
    if usagePercent >= threshold {
        // Log a warning message
        log.Printf("Warning: Disk space usage exceeds threshold of %d%%", threshold)

        // You can add additional actions here if needed

        return fmt.Errorf("Disk space usage exceeds threshold")
    }

    // Disk space is OK
    log.Println("Disk space is OK")
    return nil
}
