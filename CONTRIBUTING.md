# Contributing to OpenZam

First off, thank you for considering contributing to OpenZam! It's people like you that make OpenZam such a great tool.

We welcome contributions of all kinds, including bug fixes, new features, and documentation improvements. This document provides guidelines for contributing to this repository.

## Table of Contents
- [Code of Conduct](#code-of-conduct)
- [How Can I Contribute?](#how-can-i-contribute)
  - [Reporting Bugs](#reporting-bugs)
  - [Suggesting Enhancements](#suggesting-enhancements)
  - [Pull Requests](#pull-requests)
- [Development Setup](#development-setup)
- [Styleguides](#styleguides)

## Code of Conduct

This project and everyone participating in it is governed by the OpenZam [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

Bugs are tracked as GitHub issues. When creating an issue, please use the Bug Report template and provide as much context as possible:
* **Use a clear and descriptive title** for the issue to identify the problem.
* **Describe the exact steps which reproduce the problem** in as many details as possible.
* **Provide specific examples** to demonstrate the steps. Include links to files or copy/paste snippets, which you use in those examples.

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating a feature request, please use the Feature Request template:
* **Use a clear and descriptive title** for the issue.
* **Provide a step-by-step description** of the suggested enhancement.
* **Explain why this enhancement would be useful** to most OpenZam users.

### Pull Requests

The process described here has several goals:
- Maintain OpenZam's quality
- Fix problems that are important to users
- Engage the community in working toward the best possible open-source project

Please follow these steps to have your contribution considered by the maintainers:

1. **Fork the repository** and create your branch from `main`.
2. **Set up the development environment**.
3. **If you've added code that should be tested, add tests.**
4. **Ensure the test suite passes** by running `go test ./...`.
5. **Update the documentation** if your change impacts any existing behavior.
6. **Open a Pull Request** using the provided template and link any relevant issues.

## Development Setup

1. **Prerequisites**: Ensure you have [Golang](https://go.dev/dl/) >= 1.20 installed.
2. **Clone your fork**:
   ```bash
   git clone https://github.com/your-username/openzam.git
   cd openzam
   ```
3. **Install dependencies**:
   ```bash
   go mod tidy
   ```
4. **Run the server locally**:
   ```bash
   go run main.go
   ```

## Styleguides

### Git Commit Messages

* Use the present tense ("Add feature" not "Added feature")
* Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
* Limit the first line to 72 characters or less
* Reference issues and pull requests liberally after the first line

### Go Styleguide

* Follow standard Go conventions (use `gofmt`).
* Keep code modular and clean.
* Document public functions and types.
