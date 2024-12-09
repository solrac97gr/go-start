# Go-Start (Beta Version) 🚀

## Overview 📖
Go-Start is a powerful CLI tool designed to bootstrap Go projects with a production-ready structure. It automatically generates projects following microservices architecture patterns, clean architecture principles, and includes pre-configured Docker setup.

## Features ✨
- Clean Architecture structure
- Microservices-oriented setup
- Pre-configured Docker environment
- Automated code generation for entities
- Built-in Makefile for common operations

## Prerequisites 🗝️
Before installing Go-Start, ensure you have the following installed:
- [Go](https://golang.org/doc/install) (1.16 or later)
- [Git](https://git-scm.com/downloads)
- [Docker](https://docs.docker.com/get-docker/)

## Installation 🛠️

1. Clone the repository:
```bash
git clone git@github.com:solrac97gr/go-start.git
```

2. Navigate to the repository and run the installation script:
```bash
cd go-start
bash scripts/install.sh
```

## Usage Guide 📚

1. Navigate to your desired project location and run:
```bash
go-start
```

2. Follow the interactive prompts:
   - Enter project name (e.g., `ecommerce`)
   - Provide GitHub username (e.g., `solrac97gr`)
   - List your project entities, separated by commas (e.g., `product,user,payment`)

3. Start your application:
```bash
# Using Make
make run

# Using Docker
docker-compose up --build
```

## Project Structure 🏗️
The generated project follows clean architecture principles with the following structure:
```
your-project/
├── cmd/
│   └── http/
│       └── main.go
├── infrastructure/
│   └── docker/
│       └── Dockerfile
├── internal/
│   ├── {entity-name}/
│   │   ├── application/
│   │   │   └── {entity}_application.go
│   │   ├── domain/
│   │   │   ├── models/
│   │   │   │   └── {entity}.go
│   │   │   └── ports/
│   │   │       └── {entity}_repository.go
│   │   └── infrastructure/
│   │       ├── handler/
│   │       │   └── {entity}_handler.go
│   │       ├── repository/
│   │       │   └── {entity}_repository.go
│   │       └── server/
│   │           └── {entity}_server.go
├── pkg/
│   ├── factory/
│   │   └── {entity}_factory.go
│   └── server/
│       └── super_server.go
├── docker-compose.yaml
├── go.mod
├── go.sum
└── Makefile
```

## Current Limitations ⚠️
This beta version includes:
- MongoDB as the default database
- Fiber as the default web framework

Future releases will support:
- Multiple database options
- Various web frameworks
- Additional architectural patterns

## Contributing 🤝
Contributions are welcome! Feel free to:
- Report bugs
- Suggest new features
- Submit pull requests

## License 📄
This project is licensed under the MIT License

## Support 💬
For support, please [open an issue](https://github.com/solrac97gr/go-start/issues) on GitHub.

---
**Note**: This is a beta version. Use in production environments at your own discretion.
