# Bazel Go Hello World

A simple Go project demonstrating Bazel 7.4.1 compatibility with WORKSPACE.

## Prerequisites

- Bazel 7.4.1 (specified in `.bazelversion`)
- No Go installation required (Bazel downloads Go SDK automatically)

## Project Structure

```
bazel_go_hello_world/
├── WORKSPACE            # WORKSPACE configuration
├── .bazelversion        # Pins Bazel version to 7.4.1
├── .bazelrc            # Bazel configuration
├── go.mod              # Go module definition
├── BUILD.bazel         # Build file
├── main.go             # Main application (single file)
└── README.md           # This file
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

### Query external dependencies (WORKSPACE style):
```bash
# Query both dependencies using WORKSPACE syntax
bazel query //external:com_github_bwmarrin_snowflake union //external:com_github_google_uuid

# Query individual dependencies
bazel query //external:com_github_bwmarrin_snowflake
bazel query //external:com_github_google_uuid

# Query all external dependencies
bazel query //external:*
```

## What the Application Does

The application demonstrates:
1. Generating a Snowflake ID using the `github.com/bwmarrin/snowflake` package
2. Generating a UUID using the `github.com/google/uuid` package
3. Simple console output showing both dependencies work correctly

## Bazel 7.4.1 Compatibility Features

- **WORKSPACE**: Traditional dependency management (no MODULE.bazel)
- **External queries**: Support for `//external:` syntax
- **Platform-based resolution**: Improved toolchain selection
- **Explicit toolchain registration**: Better cross-platform compatibility

## Configuration

The project uses several Bazel 7.4.1 compatible settings in `.bazelrc`:
- Bzlmod disabled (using WORKSPACE instead)
- Platform-based toolchain resolution
- Improved caching and performance settings
- Better error reporting

## WORKSPACE vs Bzlmod

This project uses WORKSPACE for dependency management, which allows:
- Traditional `//external:` query syntax
- Explicit dependency declarations with checksums
- Compatibility with older Bazel tooling

## Module Information

- **Module name**: `github.com/sujay-biradar-25/bazel_go_hello_world`
- **Go version**: 1.21.6
- **rules_go version**: 0.48.1
- **gazelle version**: 0.35.0 