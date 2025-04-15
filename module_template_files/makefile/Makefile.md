# Makefile

## Usage Notes for Makefile

- currently a work in progress
- goal is using go templating to facilitate code generation

## Targets

### `build`

- **Description**: Builds the project and places the binary in the specified directory.
- **Usage**: `make build`

### `test`

- **Description**: Runs tests for all packages except those in the `/test/` directory.
- **Usage**: `make test`

### `test-with-cover`

- **Description**: Runs tests with coverage for all packages except those in the `/test/` directory. Generates a coverage report.
- **Usage**: `make test-with-cover`

### `generate-mocks`

- **Description**: Generates mocks using `mockery` with specific options.
- **Usage**: `make generate-mocks`

### `clean`

- **Description**: Cleans the build and removes the binary and vendor directories.
- **Usage**: `make clean`

### `run`

- **Description**: Builds the project and runs the main Go application.
- **Usage**: `make run`

### `deps`

- **Description**: Fetches all dependencies.
- **Usage**: `make deps`

### `mod`

- **Description**: Downloads and tidies up modules and creates a vendor directory.
- **Usage**: `make mod`

### `prod`

- **Description**: Builds the project for production using vendored dependencies and without debug symbols.
- **Usage**: `make prod`

### `asm`

- **Description**: Generates assembly output for debugging and optimization.
- **Usage**: `make asm`

### `lint`

- **Description**: Runs linting on the code with all enabled checks.
- **Usage**: `make lint`

---

Here's an explanation of the most relevant Go environment variables and paths you listed, detailing what they do and whether they need user management based on your current working Go module or project:

### Key Go Environment Variables

1. **`GO111MODULE`**:
   - **Purpose**: Controls whether Go modules are enabled. Possible values are `on`, `off`, or `auto`. By default, `auto` enables modules when a `go.mod` file is present.
   - **User-managed**: No need for manual management in most cases. Leave it empty or set to `auto`, unless you're working with older Go projects that don’t use modules.

2. **`GOARCH='amd64'`**:
   - **Purpose**: Specifies the architecture for which the Go code is compiled (e.g., `amd64`, `arm`, etc.).
   - **User-managed**: No, this is usually auto-detected based on your system. Change only if cross-compiling for another architecture.

3. **`GOBIN='/home/stephen/go/bin'`**:
   - **Purpose**: Directory where compiled Go binaries are installed when using `go install`.
   - **User-managed**: Yes, this can be managed if you want to control where your Go binaries are placed. Typically, it's a subdirectory of `GOPATH`, but you can set this to any path.

4. **`GOCACHE='/home/stephen/.cache/go-build'`**:
   - **Purpose**: Cache directory for compiled Go files, speeding up subsequent builds.
   - **User-managed**: No, typically doesn’t need manual management.

5. **`GOENV='/home/stephen/.config/go/env'`**:
   - **Purpose**: Stores Go environment configuration.
   - **User-managed**: No, this file is automatically managed by Go.

6. **`GOMODCACHE='/home/stephen/go/pkg/mod'`**:
   - **Purpose**: Location where Go modules are cached after they are downloaded. This cache helps in reusing modules across projects.
   - **User-managed**: No, unless you want to relocate the module cache directory. It can usually remain under `GOPATH`.

7. **`GOOS='linux'`**:
   - **Purpose**: Specifies the target operating system for the build (e.g., `linux`, `windows`, `darwin`).
   - **User-managed**: No, typically auto-detected by Go. Only change this when cross-compiling for other operating systems.

8. **`GOPATH='/home/stephen/go'`**:
   - **Purpose**: This is a critical path. It specifies the root directory for your Go workspace, where Go looks for downloaded dependencies, binaries, and source files.
   - **User-managed**: Yes, this path should be managed based on where you want your Go workspace to reside. You can set it to any directory you want, but the default is usually `$HOME/go`.

9. **`GOPROXY='https://proxy.golang.org,direct'`**:
   - **Purpose**: Specifies the proxy for downloading Go modules. The default Go proxy is `https://proxy.golang.org`.
   - **User-managed**: Typically no, but you might change this if working in an environment with custom module proxies (e.g., private corporate proxies).

10. **`GOROOT='/home/stephen/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.22.4.linux-amd64'`**:
    - **Purpose**: Directory where the Go SDK is installed. It contains the standard library and Go compiler tools.
    - **User-managed**: No, unless you are using a custom Go installation. This path is managed by the Go installer.

11. **`GOMOD='/home/stephen/code/workspaces/go-workspace/go-init_workspace/go.mod'`**:
    - **Purpose**: Path to the `go.mod` file in your current project, which defines the module and its dependencies.
    - **User-managed**: Yes, this is automatically set when working inside a Go module, but you should manage the `go.mod` file itself for dependency management.

12. **`GOWORK=''`**:
    - **Purpose**: Path to a `go.work` file for multi-module workspaces. This is used if you work with multiple Go modules in a single workspace.
    - **User-managed**: Yes, if you use multi-module workspaces, but it's empty for single-module projects.

### Variables You Generally Don’t Need to Manage

- **`GOEXE=''`**: Suffix for executable files (only relevant on Windows).
- **`GOEXPERIMENT=''`**: Flags for experimental Go features.
- **`GOFLAGS=''`**: Flags passed to all `go` commands.
- **`GOINSECURE=''`**: Allows use of insecure Git repos or modules.
- **`GONOPROXY=''`, `GONOSUMDB=''`, `GOPRIVATE=''`, `GOSUMDB='sum.golang.org'`: These are related to module proxy behavior and security. Rarely need user modification.
- **`GCCGO='gccgo'`, `GOAMD64='v1'`, `AR='ar'`, `CC='gcc'`, `CXX='g++'`, `CGO_ENABLED='1'`, `CGO_CFLAGS='-O2 -g'`: These settings control CGo (C bindings in Go) and the toolchain. These are auto-managed for most Go users.
- **`PKG_CONFIG='pkg-config'`**, **`GOGCCFLAGS='-fPIC -m64 ...'`**: Settings related to C/C++ toolchain for CGo.

### Key Paths to Manage

- **`GOPATH`**: This is where your Go workspace lives. It contains binaries, package caches, and source files. You should manage this based on your project setup.
- **`GOBIN`**: Where Go installs binaries. This can be set if you want to manage where your compiled executables are placed.
- **`GOMOD`**: While this is automatically set, you manage the `go.mod` file within your module for dependency management.

### Summary of What to Manage

1. **`GOPATH`**: Set this path based on where you want your Go workspace to be.
2. **`GOBIN`**: Manage this if you want to specify a directory for compiled binaries.
3. **`GOMOD`**: The `go.mod` file is crucial for managing module dependencies.
4. **`GOWORK`**: Manage this if using a multi-module workspace setup.
