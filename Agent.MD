# Agent

## Mission
This project delivers a Go command-line agent and Hiqh Quality Go Library which embraces Go idiomatics that wraps the OpenAI Responses API. It exposes the `oai` binary with a `responses create` workflow that assembles structured prompts, optional tool settings, and streams them to `/v1/responses`, printing model output back to the terminal.

## Capabilities
- CLI entrypoint via Cobra (`main.go`, `cmd/root.go`) with configuration sourced from flags, YAML, or environment variables using Viper.
- Request construction through `pkg/builder`, producing the typed `request.RequestDef` used by the OpenAI Responses endpoint.
- HTTP transport with configurable proxy, base URL, and bearer authentication in `pkg/net`.
- Response decoding into rich Go types in `pkg/openai/responses/response` with selector/serializer helpers for extracting the assistant message stream.

## Core Flow
1. `cmd/` initializes Cobra, wires persistent config, and registers `responses` subcommands.
2. `pkg/cli.ResponsesCreateCommand` gathers CLI options, converts them into `builder.ResponsesOption` functions, and calls the OpenAI client.
3. `pkg/client` marshals the request payload, invokes the HTTP client, and unmarshals the JSON response into `response.ResponseDef`.
4. Selectors (`response.SelectMessage`) and serializers (`response.SerializeText`) are used to print assistant text to stdout.

## Configuration
- Default configuration lives in `cmd/root.go` and can be overridden via `--config`, environment variables (`OAI_OPENAI_API_KEY`, etc.), or command flags.
- Example configuration: see `oai.yaml` for API key, proxy URL, and console output format.
- The agent requires `OPENAI_API_KEY` (or equivalent config) before making requests.

## Building Requests
- Text prompts use `registry.TextInput` to append a `user` role message.
- Files may be attached with `registry.FileInput` which base64-encodes the payload with MIME detection.
- Reasoning effort and summaries map to OpenAI enums in `builder/responses_reasoning.go`.
- Web search toggles activate a `web_search` tool entry, with optional context size parsing in `builder/responses_web_search.go`.
- All options mutate a shared `request.RequestDef`, following Go idioms of pure option functions that close over their operands.

## Output Handling
- JSON responses are parsed with registries that map `type` discriminators to concrete structs.
- `ResponseDef.SelectOutput` filters outputs, while serializers translate selected items into printable strings.
- Tests in `pkg/openai/responses/response/output_message_test.go` validate the polymorphic unmarshalling path against fixture data.

## HTTP Layer
- `pkg/net` builds `http.Request` objects, injects authorization headers, and optionally configures proxy transport.
- Errors propagate explicitly—non-200 responses return formatted errors that surface to the CLI.

## Testing & Quality Gates
- Run `go test ./...` for unit coverage across builders, CLI, and response parsing.
- `scripts/check_coverage.sh` enforces per-package coverage minima (60% for builder/cli, 50% for openai types) using an atomic coverage profile.
- Generated coverage artifacts are stored in `coverage.txt` and `coverage.html`.

## Go Idioms & Conventions
- Keep constructors (`NewClient`, `NewRegistry`) returning interfaces or pointers while hiding concrete types.
- Preserve context propagation; every outbound call accepts a `context.Context`.
- Prefer option functions over mutable global state; each option only mutates the provided request.
- Maintain explicit error checks—wrap with `fmt.Errorf` while preserving the original error value.
- Run `go fmt ./...` after changes and follow module-based dependency management in `go.mod`.

## Known Gaps / TODOs
- Streaming (`CreateStream`), retrieval, and cancellation methods are currently stubs returning `not implemented`.
- `builder.responsesWithTextInput` sets `Type` to `input_image`; correct to `input_text` before production use.
- Response models for reasoning/web-search calls have placeholder fields (`status`, `summary` formatting) that need verification against the latest API schema.
- No retry or rate-limit handling in the HTTP layer; add if interacting with production APIs.

## Extending the Agent
- Add new request options by extending `ResponseRegistry` and implementing the corresponding builder helper.
- When supporting additional output types, register the new struct in `response/output_unmarshaller.go` and add targeted tests with fixture JSON.
- Ensure CLI flags default to configuration values but require explicit overrides for sensitive inputs like prompts or file paths.
