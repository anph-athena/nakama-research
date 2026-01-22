# Development Container Configuration

This directory contains the configuration for GitHub Codespaces and VS Code Dev Containers.

## Features

- **Go 1.21** development environment
- **Docker-in-Docker** for running Nakama and CockroachDB
- **GitHub CLI** for repository management
- Pre-configured VS Code extensions
- Automatic port forwarding for all services

## Included Extensions

- **golang.go** - Go language support
- **GitHub.copilot** - AI-powered code completion
- **GitHub.copilot-chat** - AI chat assistant
- **ms-azuretools.vscode-docker** - Docker support
- **redhat.vscode-yaml** - YAML language support
- **DavidAnson.vscode-markdownlint** - Markdown linting
- **eamodio.gitlens** - Git supercharged

## Forwarded Ports

| Port  | Service              | Description                    |
|-------|----------------------|--------------------------------|
| 7349  | Nakama Client API    | HTTP/REST API                  |
| 7350  | Nakama gRPC          | gRPC API                       |
| 7351  | Nakama Console       | Admin web interface            |
| 8080  | CockroachDB UI       | Database admin interface       |
| 26257 | CockroachDB SQL      | Database connection            |

## Quick Start

### Using GitHub Codespaces

1. Click "Code" button on GitHub
2. Select "Codespaces" tab
3. Click "Create codespace on main"
4. Wait for environment to build
5. Run `./setup.sh` to start services

### Using VS Code Dev Containers

1. Install "Dev Containers" extension in VS Code
2. Open command palette (Ctrl+Shift+P)
3. Select "Dev Containers: Reopen in Container"
4. Wait for container to build
5. Run `./setup.sh` to start services

## Development Workflow

### Starting Nakama Server

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f nakama

# Stop services
docker-compose down
```

### Running Go Commands

```bash
# Format code
go fmt ./...

# Run tests
go test ./...

# Build plugin
go build -buildmode=plugin -trimpath -o /tmp/main.so main.go

# Install dependencies
go mod download
go mod tidy
```

### Accessing Services

- **Nakama Console**: Click the forwarded port notification or go to Ports tab
- **CockroachDB UI**: Access via port 8080
- **Client API**: Use port 7349 for testing with client SDKs

## Customization

Edit `.devcontainer/devcontainer.json` to:
- Add more VS Code extensions
- Change Go version
- Modify port forwarding
- Add custom scripts
- Configure environment variables

## Troubleshooting

### Docker Socket Issues

If Docker-in-Docker doesn't work:
```bash
sudo chmod 666 /var/run/docker.sock
```

### Go Tools Not Working

Reinstall Go tools:
```bash
go install golang.org/x/tools/gopls@latest
go install golang.org/x/tools/cmd/goimports@latest
```

### Port Already in Use

Check and stop conflicting services:
```bash
docker ps
docker-compose down
```

## Resources

- [Dev Containers Documentation](https://code.visualstudio.com/docs/devcontainers/containers)
- [GitHub Codespaces Documentation](https://docs.github.com/en/codespaces)
- [Go in VS Code](https://code.visualstudio.com/docs/languages/go)
