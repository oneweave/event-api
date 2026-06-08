# AI Agent Instructions for async-client

## Repository overview
- Go-based event model library for OneWeave async event contract handling.
- `asyncapi.yaml` is the leading source of truth for event schemas and validation.
- `lib/*.go` implements the Go model layer and must mirror validations from `asyncapi.yaml`.
- Models in `lib` are CloudEvent payload types and data schemas for artifact release/build lifecycle events.

## Important files
- `asyncapi.yaml` — AsyncAPI contract definitions for messages and schemas.
- `lib/release.go` — core schema structs and builder/release payload models.
- `lib/event.go` — shared `BaseEvent` CloudEvent envelope definition.
- `lib/build_events.go` — build-related CloudEvent types.
- `lib/release_events.go` — release-related CloudEvent types.
- `LICENSE` — repository license text.

## Language and tooling
- Go 1.25
- Use `gofmt -w` on Go files before committing.
- Validation command: `go test ./...`

## Coding conventions
- Maintain `json` tags on every struct field.
- Maintain `bson` tags in snake_case for all struct fields.
- Maintain `validate` tags using `github.com/go-playground/validator/v10`.
- Go validation tags must reflect the API document, including `required`, `pattern`, and `format` checks where specified.
- All optional fields mapped from AsyncAPI to Go structs must use pointer fields.
- Constructor functions must set default struct values according to `asyncapi.yaml` defaults.
- Event structs should embed `BaseEvent` and still define their own `Type` field for exact equality validation.
- Keep AsyncAPI schema and Go model validation aligned; `asyncapi.yaml` is always the leading source of truth.

## Suggested behavior for AI coding tasks
- Prefer schema-driven changes: update `asyncapi.yaml` and corresponding Go structs together.
- Do not add new package directories unless the task specifically requires expansion.
- Preserve existing naming and event semantics for CloudEvent types and payloads.
- Keep changes small and well-scoped in this small repository.
