# CLAUDE.md - Witchcraft Game Development Guidelines

## Build Commands
- Run game: `go run main.go`
- Run with specific scene: `START_SCENE=fitting_room go run main.go`
- Build: `go build -o witchcraft`
- Build for web: `GOOS=js GOARCH=wasm go build -o web/main.wasm`

## Code Style Guidelines

### Imports
- Group imports: standard library first, then third-party packages, then local packages
- Sort alphabetically within groups
- Use proper aliasing for name conflicts (e.g., `stdmath "math"`)

### Component Architecture
- Follow Entity-Component-System (ECS) pattern using donburi library
- Components define data in `/component/` directory
- Systems implement logic in `/system/` directory
- Archetypes combine components in `/archetype/` directory

### Naming Conventions
- Use CamelCase for exported names, camelCase for internal names
- Use descriptive, specific names for functions and variables
- Make component and system names match their purpose

### Error Handling
- Use `panic` only for unrecoverable errors during startup
- For runtime errors, return error values where appropriate
- Use logging for debugging information