# skeval

[![CI](https://github.com/domehahn/skeval/actions/workflows/ci.yml/badge.svg)](https://github.com/domehahn/skeval/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/domehahn/skeval)](https://goreportcard.com/report/github.com/domehahn/skeval)
[![Go Reference](https://pkg.go.dev/badge/github.com/domehahn/skeval.svg)](https://pkg.go.dev/github.com/domehahn/skeval)
[![Go Version](https://img.shields.io/github/go-mod/go-version/domehahn/skeval)](https://github.com/domehahn/skeval/blob/main/go.mod)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/domehahn/skeval/blob/main/LICENSE)
[![Status](https://img.shields.io/badge/Status-Production--Ready-brightgreen.svg)](https://github.com/domehahn/skeval)

`skeval` (**Synthetic Behavioral & Adversarial Evaluation Generator**) automatically generates test suites, boundary probes, and grading metrics for AI agent skills and tool contracts.

## Features

- **Synthetic Test Generation**: Covers functional tool calls, parameter fuzzing, boundary refusals, and injection probes.
- **Contract-Based Assertions**: Enforces `expected_action`, tool match criteria, and forbidden response terms.
- **Quantitative Scoring**: Computes categorical and aggregate containment scores (0-100%).
- **Ecosystem Compatibility**: Outputs `skillspec`-compatible eval structures.

## Installation

```bash
go install github.com/domehahn/skeval/cmd/skeval@latest
```

## Quick Start

```bash
# Generate a synthetic evaluation suite for a skill
skeval gen --skill "git-reviewer" --tools "git_diff,git_commit"
```
