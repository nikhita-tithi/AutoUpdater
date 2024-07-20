# AutoUpdater

# Self-Updating Program in Go

This Go application demonstrates a self-updating mechanism that continuously checks for updates while running and applies them seamlessly. The program performs two functions:
 1. Print the current date and time every 3 seconds.
 2. Perform a check asynchronously to see if updates are available and if they are valid updates, install and restart the application


## Features

- **Self-Update Capability**: The program can update itself by downloading a new version and replacing the current executable.
- **Periodic Update Check**: Checks for updates every 12 seconds
- **Logging**: Logs update checks and application activities.
- **Multi-OS compatible**: The program is compatible to run on Mac, Linux and Windows OS.

## Prerequisites

- Go programming language installed on your machine.

## Installation

1. **Clone the Repository**:
    ```sh
    git clone https://github.com/nikhita-tithi/AutoUpdater
    cd AutoUpdater
    ```

2. **Build the Application**:
    ```sh
    go build -o myapp main.go check_update_helper.go
    ```

## Code Structure

- `main.go`: The main entry point of the application.
- `check_update_helper.go`: Contains the logic for checking and applying updates.
- `update_version.txt`: Contains the latest version. This is used to update the client of new version availability.
- `mock_update.sh`: produces a new binary with an updated version number and updates the `update_version.txt`

## Usage

1. **Run the Application**:
    ```sh
    ./myapp
    on a separate terminal window, run this to create a new binary to mock an update : . ./mock_update.sh
    ```

2. **Logs**:
    - The application logs activities to `app.log` in the current directory. Tail using `tail -f app.log`
    - It prints the current date and time every 3 seconds.
    - It checks for updates every 12 seconds and logs the result.

3. **Graceful Shutdown**:
    - Press `Ctrl+C` to stop the application gracefully.


## Troubleshooting

- For exceptions like `Failed to apply update: failed to restart application: fork/exec /Users/<username>/Documents/AutoUpdaterWithHTTP/myapp: permission denied` try `chmod +x /Users/<username>/Documents/AutoUpdaterWithHTTP/myapp`

- After an update is installed, the main terminal window will observe that the application has terminated. Now the updated version of the application will run in the background. This is expected behavior. 
 - Just like any other background application, it can either by stopped by going to activity monitor(on mac) or task manager(Windows)
 - Or use `ps aux | grep 'myapp'` to get the processId and `kill -9 <processID>`
