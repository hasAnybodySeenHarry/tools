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

`tools` is primarily developed in Go, a minimalistic language designed to achieve efficiency and reliability. At this moment, the service is intentionally kept super simple to have a greater focus on the pipeline.

## Features

The repository has a couple of directories to focus on, mainly the following:

- `app`: the service itself and the web API
- `scripts`: scripts to initialize the database
- `charts`: Helm charts to build Kubernetes manifests
- `.githubs`: CI and CD related tasks

There's also a `Dockerfile` to build the container image.

## Getting Started

You can clone the repository and run the app locally. Since it's a standalone service with no external dependencies, you should have no difficulties running it. Make sure to create an .env file and populate it with your custom credentials.

### Prerequisites

I don't expect you to be an expert in any of the technologies used, but it would be helpful if you're familiar with:

- Go (Golang)
- Docker
- Kubernetes
- Some CI/CD

If you need a little warm-up, there are a few resources I can point you to.

### Installation

To run the app service locally, follow these steps:

#### Environment

Make sure you have the following installed on your machine:

- [Go](https://golang.org/doc/install): The Go programming language.
- [Docker](https://www.docker.com/get-started): Docker tools to containerize.
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git): VCS.
- [MySQL](https://dev.mysql.com/downloads/mysql): MySQL or any relational db.

#### Clone the Repository

```bash
git clone https://github.com/hasAnybodySeenHarry/tools.git
cd tools
```

#### Set Up Environment Variables

Create a database.env file in the app folder and add the following:

```bash
# replace with your own database details
DatabaseURL=user:password@tcp(localhost:3306)/db_name
```
If you give your database's name other than `tools`, make sure to modify the ``scripts/init.sql`` as well.

Also, create a jwtSecret.json file in the secret folder and populate with a JSON object with an entry of "key": "secret".

```bash
{
  "key": "your secret jwt token"
}
```

#### Run the service

```bash
go run main.go
```
Now you have the `tools` up and running locally! Feel free to explore the API and make changes as needed.

## Contributing

If you're interested in contributing to this project, please follow the guidelines in [CONTRIBUTING.md](CONTRIBUTING.md).

## License

This project is licensed under the [Creative Commons Attribution-NonCommercial 4.0 International License](LICENSE.md) - see the [LICENSE.md](LICENSE.md) file for details.
