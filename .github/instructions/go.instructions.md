---
applyTo: "**/*.go"
---

# Go Instructions

This file now contains only always-on guidance. Detailed, task-specific rules are enforced via repository skills in `.github/skills/`.

# Always-On Rules

- Handle all returned errors explicitly.
- Wrap propagated errors with `fmt.Errorf` and `%w`.
- Use the request parent context for production paths; only use `context.Background()` in exceptional cases or tests.
- Keep code idiomatic and readable (explicit over clever).

# Notes

- Keep this file short and stable.
- Put future detailed implementation guidance in dedicated skills, not here.
