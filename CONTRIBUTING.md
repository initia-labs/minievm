# Contributing

This repository uses Conventional Commits for commit messages and pull request
titles. Before opening a PR, make sure your change matches the repository's
existing Go, protobuf, contract, and generated-file workflows.

## Commit Convention

Use this format for commits:

```text
type(scope): subject
```

Examples:

```text
feat(evm): add tx simulation guard
fix(jsonrpc): normalize block tag parsing
docs(contributing): update validation commands
test(app): cover ante handler wiring
chore(swagger): regenerate API docs
ci(spellcheck): run typo fixes on main
```

Rules:

- use lowercase `type` and `scope`
- keep the subject short and imperative
- do not end the subject with a period
- if the change is breaking, add `!` in the Conventional Commit prefix

Common types used in this repository:

- `feat`
- `fix`
- `docs`
- `refactor`
- `test`
- `chore`
- `build`
- `ci`

## Branch Naming

Use a short branch name that describes the change.

Format:

```text
type/short-description
```

Examples:

```text
fix/swagger-initia-ibc
feat/evm-estimate-gas
docs/contributing-guide
test/jsonrpc-block-tags
```

Keep branches focused. If a branch mixes unrelated changes, split it before
opening a PR.

## Pull Requests

PR titles should follow the same Conventional Commit format:

```text
type(scope): subject
```

For breaking changes, add `!` in the PR title prefix:

```text
type(scope)!: subject
```

Follow the PR template in [.github/PULL_REQUEST_TEMPLATE.md](/Users/beer-1/Workspace/minievm/.github/PULL_REQUEST_TEMPLATE.md).
At minimum, every PR should clearly describe:

- what changed
- why it changed
- how it was validated
- whether the change is breaking

If the change is tied to an issue, proposal, or spec, link it in the PR body.

## Validation

Run the smallest relevant validation set for your change before pushing.

Common commands:

- focused package tests: `go test ./path/to/package -run <TestName> -count=1`
- full unit and integration suite: `make test`
- full repository suite: `make test-all`
- unit tests only: `make test-unit`
- integration tests only: `make test-integration`
- race-enabled suite: `make test-race`
- coverage run: `make test-cover`
- benchmark run: `make benchmark`
- benchmark e2e run: `make benchmark-e2e`
- fuzz target: `make fuzz`
- lint: `make lint`

Prefer focused tests while iterating, then run the broader repository command
that matches the risk of the change.

Examples:

- EVM keeper or state transition change:
  `go test ./x/evm/keeper -run <TestName> -count=1`
- app wiring change:
  `go test ./app/... -run TestNonExistent -count=0`
- JSON-RPC behavior change:
  `go test ./jsonrpc/... -run <TestName> -count=1`

If a change affects cross-module behavior, add or run at least one regression
test that exercises the full path.

## Formatting and Generated Files

When you change Go code, keep formatting and imports clean and make sure lint
still passes.

Common commands:

- `make lint`
- `make lint-fix`

When you change Solidity contracts under `x/evm/contracts`, regenerate the
derived outputs:

- `make contracts-gen`

When you change protobuf definitions, regenerate the related outputs:

- `make proto-gen`
- `make proto-swagger-gen`
- `make proto-pulsar-gen`

For protobuf changes, also run:

- `make proto-lint`

If swagger or generated protobuf outputs change, include those generated files
in the same PR.

## Scope Discipline

Keep each PR focused on one logical change. Avoid mixing unrelated refactors,
formatting-only edits, generated-file churn, and behavior changes unless they
must land together.
