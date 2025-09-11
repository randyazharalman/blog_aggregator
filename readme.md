# Gator - Blog Aggregator CLI

A command-line tool for managing and aggregating blog feeds.

## Installation

1. Clone this repository
2. Install dependencies: `go mod tidy`
3. Build the application: `go build -o gator .`
4. (Optional) Add to PATH for global access

## Usage

### Quick Start

```bash
# Install dependencies
go mod tidy

# Login with your username
go run . login alice

# Check your status
go run . status

# See all available commands
go run . --help
```

### Available Commands

#### `login <username>`
Set the current user in the configuration file.

```bash
go run . login alice
go run . login john_doe
```

#### `status`
Show current configuration status including user and database settings.

```bash
go run . status
```

#### `version`
Display version information.

```bash
go run . version
```

#### `reset`
Reset configuration to defaults (useful for testing).

```bash
go run . reset
```

#### `help`
Get help for any command.

```bash
go run . --help
go run . login --help
```

## Configuration

The application stores configuration in `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://example",
  "current_user_name": "alice"
}
```

### Configuration File Location

- **Windows**: `C:\Users\<username>\.gatorconfig.json`
- **macOS/Linux**: `~/.gatorconfig.json`

The file is automatically created with default values when you first run any command.

## Examples

### Basic Workflow

```bash
# 1. Login
$ go run . login alice
✅ User has been set to: alice

# 2. Check status
$ go run . status
📊 Current Status:
User: alice
  ✅ Logged in as alice
Database URL: postgres://example
Config file: /Users/alice/.gatorconfig.json

# 3. Reset if needed
$ go run . reset
🔄 Configuration has been reset to defaults
```

### Error Handling

```bash
# No username provided
$ go run . login
Error: accepts 1 arg(s), received 0

# Empty username
$ go run . login ""
Error: username cannot be empty

# Invalid command
$ go run . invalid
Error: unknown command "invalid" for "gator"
```

## Project Structure

```
gator/
├── main.go                 # Application entry point
├── internal/
│   ├── cli/                # CLI command implementations
│   │   ├── root.go         # Root command and state management
│   │   ├── login.go        # Login command
│   │   ├── status.go       # Status command
│   │   ├── version.go      # Version command
│   │   └── reset.go        # Reset command
│   └── config/
│       └── config.go       # Configuration management
├── go.mod                  # Go module definition
└── README.md               # This file
```

## Clean Code Principles

This project follows clean architecture principles:

### 1. **Separation of Concerns**
- `main.go`: Entry point only
- `internal/cli/`: All CLI logic separated by command
- `internal/config/`: Configuration handling only

### 2. **Single Responsibility**
- Each command has its own file
- Each function has a single, clear purpose
- Clear naming conventions

### 3. **Dependency Direction**
- `main` depends on `cli`
- `cli` depends on `config`
- No circular dependencies

### 4. **Error Handling**
- Consistent error wrapping with context
- Clear error messages for users
- Proper error propagation

## Development

### Adding New Commands

1. Create a new `cobra.Command` in `main.go`
2. Add it to the root command in the `init()` function
3. Implement the command logic in the `RunE` function

Example:
```go
var newCmd = &cobra.Command{
    Use:   "new-command",
    Short: "Description of new command",
    RunE: func(cmd *cobra.Command, args []string) error {
        // Command implementation
        return nil
    },
}

func init() {
    rootCmd.AddCommand(newCmd)
}
```

### Building for Production

```bash
# Build for current platform
go build -o gator .

# Build for different platforms
GOOS=windows GOARCH=amd64 go build -o gator.exe .
GOOS=linux GOARCH=amd64 go build -o gator-linux .
GOOS=darwin GOARCH=amd64 go build -o gator-mac .
```

## Next Steps

This is a foundation for a blog aggregator CLI. Future features could include:

- `add-feed <url>` - Add RSS/Atom feeds
- `list-feeds` - Show all subscribed feeds
- `browse` - Browse latest posts
- `read <post-id>` - Read a specific post
- `remove-feed <url>` - Remove a feed subscription
- Database integration for storing feeds and posts