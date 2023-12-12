
# Node.js Project Generator

This tool is a command-line application to generate Node.js projects. It is built using Cobra in Go and provides functionalities for creating a new project directory, initializing a Node.js application, and installing required npm packages.

## Features

- Create a new project directory.
- Initialize a Node.js application with `npm init`.
- Install npm libraries.
- Install npm development libraries.

## Usage

```
nixjs node-gen [flags]
```

### Flags

- `-D`, `--directory`: Set the output directory for the project. Default is `./myApp`.
- `-y`, `--yes`: Generate default NodeJs package.json file without manual input.
- `-l`, `--libs`: Specify a list of Node.js libraries to install. Separate multiple libraries with spaces and enclose them in quotes.
- `-d`, `--dev-libs`: Specify a list of Node.js development libraries to install. Separate multiple libraries with spaces and enclose them in quotes.

### Example

```
nixjs node-gen -D ./projects/myapp -l "lib1 lib2" -d "lib3 lib4"
```

This command will create a project in `./projects/myapp`, initialize a Node.js project, install `lib1` and `lib2` as dependencies, and `lib3` and `lib4` as development dependencies.

## Installation and Setup

Ensure you have Go and npm installed on your system to use this tool.

1. Clone the repository.
2. Build the project using Go.
3. Run the built executable with the desired flags.

## Contribution

Feel free to contribute to this project by submitting issues and pull requests.

## License

This project is licensed under the [MIT License](LICENSE).
