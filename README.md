# Bazel Go Hello World

A simple Go project demonstrating Bazel 7.4.1 compatibility with Bzlmod.

## Prerequisites

- Bazel 7.4.1 (specified in `.bazelversion`)
- No Go installation required (Bazel downloads Go SDK automatically)

## Project Structure

```
bazel_go_hello_world/
├── MODULE.bazel          # Bzlmod configuration
├── .bazelversion         # Pins Bazel version to 7.4.1
├── .bazelrc             # Bazel configuration
├── go.mod               # Go module definition
├── BUILD.bazel          # Root build file
├── main.go              # Main application
├── lib/
│   ├── BUILD.bazel      # Library build file
│   └── lib.go           # Library code
└── internal_deps/
    ├── BUILD.bazel      # Internal deps build file
    └── internal_deps.go # Internal dependencies
```

## Dependencies

- **External Go packages:**
  - `github.com/bwmarrin/snowflake` - Snowflake ID generation
  - `github.com/google/uuid` - UUID generation

## Building and Running

### Build the application:
```bash
bazel build //:my_app
```

### Run the application:
```bash
bazel run //:my_app
```

### Query external dependencies:
```bash
bazel query '@com_github_bwmarrin_snowflake//:snowflake union @com_github_google_uuid//:uuid'
```

## Bazel 7.4.1 Compatibility Features

- **Bzlmod**: Modern dependency management (no WORKSPACE file)
- **Explicit toolchain registration**: Better cross-platform compatibility
- **Platform-based resolution**: Improved toolchain selection
- **Lockfile support**: Reproducible builds with `MODULE.bazel.lock`

## Configuration

The project uses several Bazel 7.4.1 compatible settings in `.bazelrc`:
- Bzlmod enabled by default
- Platform-based toolchain resolution
- Improved caching and performance settings
- Better error reporting

## Module Information

- **Module name**: `github.com/sujay-biradar-25/bazel_go_hello_world`
- **Go version**: 1.21.6
- **rules_go version**: 0.50.1
- **gazelle version**: 0.40.0 