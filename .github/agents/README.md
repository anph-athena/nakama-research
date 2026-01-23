# GitHub Copilot Agentic System

This directory contains agent configurations for GitHub Copilot to provide specialized assistance for different aspects of the nakama-research project.

## Available Agents

### 1. Go Expert (`go-expert.md`)
**Specialization**: Go programming language and best practices

**Use for**:
- Writing and reviewing Go code
- Optimizing Go performance
- Managing Go modules and dependencies
- Implementing Go testing
- Following Go idioms and conventions

### 2. Nakama Expert (`nakama-expert.md`)
**Specialization**: Nakama game server development

**Use for**:
- Implementing Nakama runtime modules
- Designing matchmaking systems
- Configuring Nakama server
- Implementing real-time multiplayer features
- Using Nakama APIs effectively

### 3. Docker & DevOps (`docker-devops.md`)
**Specialization**: Containerization and deployment

**Use for**:
- Docker and Docker Compose configuration
- Database setup and management
- Service orchestration
- Health checks and monitoring
- Production deployment strategies

### 4. Testing & QA (`testing-qa.md`)
**Specialization**: Testing and quality assurance

**Use for**:
- Writing unit and integration tests
- Implementing test automation
- Load testing matchmaking systems
- Ensuring code quality
- Performance testing

### 5. Documentation (`documentation.md`)
**Specialization**: Technical documentation

**Use for**:
- Writing and updating README files
- Creating API documentation
- Writing tutorials and guides
- Code comments and inline documentation
- Maintaining documentation quality

## How to Use

When working on specific tasks, GitHub Copilot can reference these agent profiles to provide more contextual and specialized assistance:

1. **Go Code Development**: Reference `go-expert.md` for Go-specific guidance
2. **Nakama Integration**: Reference `nakama-expert.md` for Nakama-specific patterns
3. **Deployment Changes**: Reference `docker-devops.md` for infrastructure guidance
4. **Adding Tests**: Reference `testing-qa.md` for testing strategies
5. **Documentation Updates**: Reference `documentation.md` for writing standards

## Agent Structure

Each agent file contains:
- **Expertise**: Areas of specialization
- **Responsibilities**: What the agent should focus on
- **Context**: Project-specific information
- **Guidelines**: Best practices and patterns
- **Examples**: Code snippets and patterns

## Maintaining Agents

When the project evolves, update the relevant agent files to:
- Add new patterns and best practices
- Update context with new components
- Include lessons learned
- Add common pitfalls to avoid
- Update guidelines based on team decisions

## Project Context

This project is a Nakama game server implementation with:
- **Language**: Go
- **Server**: Nakama 3.21.1
- **Database**: CockroachDB
- **Deployment**: Docker Compose
- **Focus**: Level-based matchmaking system

## Contributing

When adding new agents or updating existing ones:
1. Follow the established structure
2. Provide clear, actionable guidance
3. Include relevant code examples
4. Keep context up-to-date
5. Focus on project-specific needs
