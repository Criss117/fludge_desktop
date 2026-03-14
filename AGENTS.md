# AGENTS.md - Developer Guide

This is a Wails desktop application with a Go backend and React/TypeScript frontend.

## Project Structure

```
/                    # Go backend (Wails app)
  internal/          # Go packages (clean architecture)
    auth/
      application/  # Use cases
      domain/       # Entities, repository interfaces
      infrastructure/ # Repository implementations
    shared/db/      # Database layer
  main.go           # Entry point
  app.go            # App struct definition
  wails.json        # Wails configuration

/frontend            # React/TypeScript frontend
  src/
    modules/        # Feature modules (auth, inventory, teams, employees, organizations)
      [module]/
        presentation/
          screen/  # Screen components
          component/ # Reusable components
    routes/        # TanStack Router file-based routing
  package.json     # Dependencies (bun)
```

## Build Commands

### Frontend (React/TypeScript)

```bash
cd frontend

# Install dependencies
bun install

# Development (hot reload)
bun run dev

# Production build
bun run build

# Preview production build
bun run preview
```

### Go Backend (Wails)

```bash
# Development (hot reload frontend + Go)
wails dev

# Production build (generates executable)
wails build
```

### Linting

```bash
cd frontend
bun run lint         # Not configured, use: npx eslint
bunx eslint src/      # Lint TypeScript/React files
```

### Single Test Execution

No tests are currently configured. To add tests:

- Frontend: Add Vitest/Jest to package.json
- Go: Add testing with `go test ./...`

## Code Style Guidelines

### General

- Keep files focused and small (under 200 lines when possible)
- Use meaningful variable/function names
- Avoid magic numbers - use constants
- Handle errors explicitly, never ignore with `_`

### Frontend (React/TypeScript)

**Imports:**

- Use path aliases defined in `vite.config.ts`:
  - `@/` → `./src/`
  - `@modules/[module]/` → `./src/modules/[module]/`
  - `@wails/` → `./wailsjs/`
- Order: React imports → external libs → internal modules → relative imports

**Types:**

- Use TypeScript strict mode (`tsconfig.app.json`)
- Prefer interfaces over types for object shapes
- Use `function` declarations over arrow functions for components
- Avoid `any`, use `unknown` when type is truly uncertain

**Naming:**

- Components: PascalCase (e.g., `SignInScreen.tsx`)
- Hooks: camelCase starting with `use` (e.g., `useAuth`)
- Files: kebab-case (e.g., `sign-in.screen.tsx`)
- CSS classes: Tailwind utility classes (prefer concise)

**React Patterns:**

- Use TanStack Router file-based routing in `src/routes/`
- Export components as named exports
- Use React 19 features (no explicit children prop typing needed)

**Formatting:**

- ESLint handles formatting (follow `eslint.config.js`)
- Use 2-space indentation
- Trailing commas in objects/arrays

### Go Backend

**Imports:**

- Group: standard library → external packages → internal packages
- Use explicit package imports (no unnamed imports)

**Naming:**

- PascalCase for exported identifiers
- camelCase for unexported
- Acronyms: `userID` not `userId` or `UserID`
- Avoid abbreviations unless well-known (URL, API, ID)

**Error Handling:**

- Always handle errors, never ignore
- Return errors early, wrap with context using `fmt.Errorf("context: %w", err)`
- Use custom error types for domain-specific errors

**Architecture (Clean Architecture):**

- `domain/` - Entities, repository interfaces (no external dependencies)
- `application/` - Use cases, DTOs
- `infrastructure/` - Repository implementations, external services
- Dependencies point inward: domain → application → infrastructure

**Patterns:**

- Constructor functions return pointers: `func NewX(...) *X`
- Use interfaces for dependencies (enables mocking)
- Embed `context.Context` as first parameter in method signatures

## Database

- SQLite with `modernc.org/sqlite` driver
- Schema in `internal/shared/db/schema.sql` (embedded via `//go:embed`)
- Generated code via sqlc in `internal/shared/db/generated/`

## Key Dependencies

### Frontend

- React 19, React DOM 19
- TanStack Router (file-based routing)
- Tailwind CSS v4, @tailwindcss/vite
- ShadCN (@shadcn/react)
- Vite 7

### Backend

- Go 1.24
- Wails v2
- modernc.org/sqlite

## Environment

- Node.js: bun (package manager)
- Go: 1.24+
- OS: Cross-platform (Windows, macOS, Linux)
