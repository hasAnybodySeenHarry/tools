# Tools

A lightweight microservice designed to demonstrate the end-to-end experience of the development phases.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project is primarily developed in Go, a minimalistic language designed to achieve efficiency and reliability. The service is intentionally kept super simple to make it approachable for beginners.

## Features

The repository has a couple of directories to focus on, mainly the following:

- `app`: the service itself and the web API
- `scripts`: scripts to initialize the database
- `charts`: Helm charts to build Kubernetes manifests
- `.githubs`: CI and CD related tasks

There's also a dockerfile to build the container image.

## Getting Started

You can clone the repository and run the app locally. Since it's a standalone service with no external dependencies, you should have no difficulties running it. Make sure to create an .env file and populate it with your custom credentials.

### Prerequisites

I don't expect you to be an expert in any of the technologies used, but it would be helpful if you're familiar with:

- Go (Golang)
- Docker
- Kubernetes
- Helm
- Some CI/CD

If you need a little warm-up, there are a few resources I can point you to.

### Installation

## Contributing

If you're interested in contributing to this project, please follow the guidelines in [CONTRIBUTING.md](CONTRIBUTING.md).

## License

This project is licensed under the [Creative Commons Attribution-NonCommercial 4.0 International License](LICENSE.md) - see the [LICENSE.md](LICENSE.md) file for details.
