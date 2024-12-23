# Time Tracker CLI

This is a command-line interface (CLI) tool for tracking your TODO projects. It allows you to create projects, list them, stop tracking, generate reports, and edit project details.

## Features

- **Create Projects**: Initialize new projects to track your tasks.
- **List Projects**: Display a list of all your projects.
- **Stop Tracking**: Stop tracking time for a specific project.
- **Generate Report**: Create a report of the time spent on each project.
- **Edit Projects**: Modify the details of your existing projects.

## Installation

To install the Time Tracker CLI, clone the repository and install the dependencies:

```sh
git clone https://github.com/yourusername/timetrack.git
cd timetrack
npm install
```

## Usage

Here are the basic commands to use the Time Tracker CLI:

- **Create a new project**:
    ```sh
    timetrack create <project-name>
    ```

- **List all projects**:
    ```sh
    timetrack list
    ```

- **Stop tracking a project**:
    ```sh
    timetrack stop <project-name>
    ```

- **Generate a report**:
    ```sh
    timetrack report
    ```

- **Edit a project**:
    ```sh
    timetrack edit <project-name>
    ```

Alternatively, you can use the executable directly:

- **Start tracking a task**:
    ```sh
    .\timetrack.exe start "My first task"
    ```

- **List all tasks**:
    ```sh
    .\timetrack.exe list
    ```

- **Stop tracking a task**:
    ```sh
    .\timetrack.exe stop 1
    ```

- **Generate a report**:
    ```sh
    .\timetrack.exe report
    ```

## Building

To build the project, run:

```sh
go build
```

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License.
