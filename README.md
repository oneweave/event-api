# OneWeave Event API client (`async-client`)

A Go-based event model library for handling OneWeave asynchronous event contracts. 

The API schema definitions are managed in [asyncapi.yaml](asyncapi.yaml), and the Go code implements the model layer matching these schemas and validations.

## Features

- **Contract-Driven**: Schema-driven event payloads mirroring the definitions in the AsyncAPI contract.
- **Prefixed UUIDv7 Event Identifiers**: Uses [event-id](https://github.com/oneweave/event-id) for human-friendly, prefixed UUIDv7 IDs (e.g. `evt_06bgm7733st2576nx5jht4ecjw`).
- **Domain Prefixes**:
  - `evt` for event envelope / system
  - `rel` for releases
  - `bld` for builds
  - `brk` for broker updates
  - `ctl` for controller updates
- **Built-in Validator**: Custom struct validation using the `eventid` tag under the hood, with generic parsing and helper functions out of the box.

## Repository Structure

- `asyncapi.yaml` — Leading source of truth AsyncAPI contract.
- `lib/`
  - `envelope.go` — CloudEvent envelope definition (`BaseEvent`).
  - `validator.go` — Core validation registration and helpers.
  - `models.go` — Common schemas (e.g. release source, platform specs).
  - `build/` — Build lifecycle event payloads and types.
  - `release/` — Release lifecycle event payloads and types.
  - `broker/` — Broker update event payloads and types.
  - `controller/` — Controller update event payloads and types.

## Installation

```go
import "github.com/oneweave/event-api/lib"
```

## Usage

### 1. Parsing and Validating an Event
Consumers can easily deserialize and run structural validation against incoming CloudEvents using `ParseAndValidate`:

```go
package main

import (
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/oneweave/event-api/lib"
	"github.com/oneweave/event-api/lib/build"
)

func handleEvent(event cloudevents.Event) {
	// Automatically decodes event payload and runs validation against structural tags
	data, err := lib.ParseAndValidate[build.ArtifactBuildRequestedData](event)
	if err != nil {
		fmt.Printf("invalid build requested payload: %v\n", err)
		return
	}

	fmt.Printf("Valid build ID received: %s\n", data.BuildID)
}
```

### 2. Manual Struct Validation
If you already have a populated struct, you can validate it directly:

```go
package main

import (
	"fmt"

	"github.com/oneweave/event-api/lib"
	"github.com/oneweave/event-api/lib/release"
)

func main() {
	payload := release.ArtifactReleaseBaseData{
		ReleaseID:  "rel_06bgm7733st2576nx5jht4ecjw",
		ArtifactID: "art_invalid_format", // will fail validation
	}

	err := lib.ValidateStruct(payload)
	if err != nil {
		fmt.Printf("Validation failed: %v\n", err)
	}
}
```

## Running Tests

Verify models and validation tags are fully correct:

```bash
go test ./...
```
