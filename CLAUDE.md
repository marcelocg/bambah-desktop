# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Building and Running
```bash
# Run the desktop application
go run cmd/main.go

# Build binary for current platform
go build -o bambah-desktop cmd/main.go

# Cross-platform builds
GOOS=windows GOARCH=amd64 go build -o bambah-desktop.exe cmd/main.go
GOOS=darwin GOARCH=amd64 go build -o bambah-desktop-mac cmd/main.go
GOOS=linux GOARCH=amd64 go build -o bambah-desktop-linux cmd/main.go

# Install dependencies
go mod tidy

# Format code (run before committing)
go fmt ./...

# Run static analysis
go vet ./...
```

### Testing Commands
```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run specific package tests
go test ./internal/ui/forms
```

## Project Architecture

This is a **Fyne-based desktop GUI application** for the Bambah financial management ecosystem.

### Core Components
- **cmd/main.go:15** - Application entry point and metadata setup
- **internal/ui/app.go:15** - Main window and application lifecycle management
- **internal/ui/forms/entry_form.go:25** - Financial entry form with Fyne widgets

### Key Design Patterns
- **SDK Integration**: Uses `bambah.FinancialService` interface for all backend operations
- **Widget Composition**: Fyne widgets organized in containers for responsive layout
- **Validation Pipeline**: Real-time field validation with user feedback
- **Environment Configuration**: Automatic backend detection via SDK

### Data Flow
1. UI captures user input through Fyne form widgets
2. Form validation occurs in real-time for immediate feedback
3. Data transformation (value sign adjustment, date formatting)
4. SDK service creation via `bambah.NewFinancialServiceFromEnv()`
5. Entry persistence through `service.CreateFinancialEntry()`
6. Success/error feedback displayed to user

## Key Files and Structure

### Main Application
- `cmd/main.go:8` - App initialization and metadata configuration
- `internal/ui/app.go:28` - Main window setup and content layout
- `internal/ui/app.go:47` - Content building with Fyne container structure

### Form Implementation
- `internal/ui/forms/entry_form.go:25` - Main form structure and widgets
- `internal/ui/forms/entry_form.go:58` - Widget initialization and configuration
- `internal/ui/forms/entry_form.go:106` - Form layout and container building
- `internal/ui/forms/entry_form.go:119` - Dynamic data loading from backend
- `internal/ui/forms/entry_form.go:135` - Save operation with validation
- `internal/ui/forms/entry_form.go:197` - Form clearing and reset functionality

## Required Environment Variables
```bash
APPWRITE_ENDPOINT=https://cloud.appwrite.io/v1
APPWRITE_PROJECT_ID=your-project-id
APPWRITE_DATABASE_ID=your-database-id
APPWRITE_API_KEY=your-api-key
APPWRITE_COLLECTION_ID=entries  # Optional, defaults to "entries"
```

## Fyne-Specific Patterns

### Widget Management
- **Radio Groups**: Entry type selection (Income/Expense)
- **Select Widgets**: Dynamic dropdowns for accounts and categories
- **Entry Widgets**: Text input with real-time validation
- **Form Layout**: Organized form structure with labels
- **Button Actions**: Save and clear operations with feedback

### Container Hierarchy
```
Border Container
├── Header: Title + Separator
└── Center: VBox Container
    ├── Radio Group Card (Entry Type)
    ├── Form Widget (Fields)
    └── Button Grid (Actions)
```

### Validation Strategy
- **Field Validators**: Attached to individual widgets for real-time feedback
- **Form Validation**: Pre-save validation for required fields
- **Error Display**: Modal dialogs for user-friendly error messages
- **Success Feedback**: Confirmation dialogs with automatic form clearing

## Integration with Bambah SDK

### Service Creation
```go
service, err := bambah.NewFinancialServiceFromEnv()
// Automatic backend detection and configuration
```

### Dynamic Data Loading
```go
accounts, err := service.ListAccounts()    // Populate account dropdown
categories, err := service.ListCategories() // Populate category dropdown
```

### Entry Persistence
```go
entry := types.FinancialEntry{...}
err := service.CreateFinancialEntry(entry) // Backend-agnostic save
```

## Development Workflow
1. Make changes to UI or form components
2. Test with sample data: `go run cmd/main.go`
3. Verify form validation and error handling
4. Format code: `go fmt ./...`
5. Run static analysis: `go vet ./...`
6. Build for target platforms if needed

## Testing Strategy
- Widget behavior testing (form validation, data binding)
- Service integration testing (SDK connectivity)
- Cross-platform build verification
- User interaction flow testing

## Technical Decisions
- **Fyne over other GUI frameworks**: Native performance, Go-native API, cross-platform compatibility
- **Local SDK dependency**: Uses `replace` directive for development, will be updated when SDK is published
- **Container-based layout**: Responsive design that adapts to different screen sizes
- **Real-time validation**: Immediate user feedback instead of form-submit validation
- **Modal feedback**: Clear success/error communication through dialogs