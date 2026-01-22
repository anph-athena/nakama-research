# GitHub Copilot Configuration

This directory contains configurations for GitHub Copilot to provide enhanced AI assistance for the Nakama research project.

## Overview

The Copilot configuration system consists of three main components:

1. **Skills** (`skills.md`) - Specialized knowledge areas
2. **Subagents** (`subagents.md`) - Focused task specialists
3. **Integration** - How skills and subagents work together

## Structure

```
.github/copilot/
├── README.md       # This file
├── skills.md       # Skill definitions
└── subagents.md    # Subagent definitions
```

## Skills vs Subagents

### Skills
**What**: Broad areas of expertise
**Purpose**: Provide domain knowledge
**Activation**: Based on file context and task type
**Examples**: 
- Nakama Development
- Go Server Development
- Docker Deployment

### Subagents
**What**: Specialized workers under main agents
**Purpose**: Handle specific subtasks
**Activation**: When parent agent is active and task matches specialty
**Examples**:
- Go Concurrency Subagent (under Go Expert)
- Matchmaking Subagent (under Nakama Expert)
- Load Testing Subagent (under Testing & QA)

## How It Works

### 1. Context Detection

Copilot analyzes:
- Current file(s) being edited
- Recent files in workspace
- Task description or prompt
- Related agents in `.github/agents/`

### 2. Skill Activation

Skills activate automatically when:
```
User edits: main.go
→ Detects: Go code + Nakama imports
→ Activates: go-server-dev + nakama-dev skills
→ References: go-expert.md + nakama-expert.md agents
```

### 3. Subagent Delegation

When specific task identified:
```
User task: "Optimize matchmaking performance"
→ Parent Agent: nakama-expert
→ Subagent: matchmaking-subagent
→ Additional: go-performance-subagent
→ Output: Optimized code with benchmarks
```

## Using the System

### Implicit Usage

Simply work on files, and the system activates automatically:

```bash
# Editing main.go
# → go-server-dev skill active
# → nakama-dev skill active
# → go-expert agent provides guidance

# Editing docker-compose.yml
# → docker-deployment skill active
# → docker-devops agent provides guidance

# Writing tests
# → testing-qa skill active
# → testing-qa agent provides guidance
```

### Explicit Requests

Request specific skills or subagents:

```
"Use the matchmaking-subagent to implement skill-based matching"
"Apply the go-performance-subagent to optimize this function"
"Use the load-testing-subagent to create stress tests"
```

## Skill-Agent Mapping

| Skill | Related Agent | Subagents |
|-------|---------------|-----------|
| nakama-dev | nakama-expert | matchmaking, storage, real-time |
| go-server-dev | go-expert | concurrency, testing, performance |
| matchmaking-system | nakama-expert | matchmaking |
| docker-deployment | docker-devops | container-optimization, orchestration, database |
| testing-qa | testing-qa | load-testing, integration-testing, test-automation |

## Example Workflows

### Workflow 1: Implement New Matchmaking Feature

```
1. User opens main.go
   → go-server-dev + nakama-dev skills activate

2. User: "Add skill-based matchmaking"
   → matchmaking-subagent activates
   → Implements matchmaking logic
   → go-testing-subagent adds tests
   → api-documentation-subagent documents API

3. User opens docker-compose.yml
   → docker-deployment skill activates
   → Validates deployment configuration
```

### Workflow 2: Fix Performance Issue

```
1. User: "Matchmaking is slow under load"
   → go-performance-subagent analyzes code
   → Identifies bottlenecks
   → Suggests optimizations

2. User applies optimizations
   → load-testing-subagent creates benchmarks
   → Validates improvements

3. Documentation subagent updates docs
   → Notes performance characteristics
```

### Workflow 3: Add Integration Tests

```
1. User: "Add end-to-end tests for matchmaking"
   → testing-qa skill activates
   → integration-testing-subagent designs tests
   → go-testing-subagent implements tests

2. User: "Run tests in CI"
   → test-automation-subagent updates workflow
   → Configures GitHub Actions

3. User: "Document test procedures"
   → documentation agent updates TESTING.md
```

## Configuration

### Adding New Skills

Edit `skills.md`:
```markdown
### N. New Skill Name
**ID**: `new-skill-id`
**Description**: What this skill does
**Capabilities**:
- Capability 1
- Capability 2
**Context Files**: Relevant file patterns
**Related Agents**: agent names
```

### Adding New Subagents

Edit `subagents.md`:
```markdown
### N. New Subagent
**Parent Agent**: parent-agent-name
**Specialty**: Specific area
**Responsibilities**: What it handles
**Use Cases**: When to use
```

## Integration with Agents

Copilot configuration works alongside `.github/agents/`:

```
.github/
├── agents/              # Agent personalities and expertise
│   ├── go-expert.md
│   ├── nakama-expert.md
│   └── ...
└── copilot/            # Copilot-specific configuration
    ├── skills.md       # Skill definitions
    └── subagents.md    # Subagent hierarchy
```

**Agents** provide the knowledge and personality
**Skills** define what knowledge to apply
**Subagents** handle specific subtasks

## Best Practices

### 1. Keep Skills Broad
- Skills should cover general domains
- Don't create too many overlapping skills
- Link skills to agents for detailed guidance

### 2. Make Subagents Focused
- Each subagent has one clear specialty
- Subagents handle specific tasks well
- Multiple subagents can collaborate

### 3. Maintain Consistency
- Use consistent naming conventions
- Keep documentation updated
- Link related skills, agents, and subagents

### 4. Test Configuration
- Verify skills activate correctly
- Check subagent delegation works
- Ensure agent references are valid

## Troubleshooting

### Skill Not Activating

1. Check file pattern matches
2. Verify agent reference exists
3. Ensure task description is clear

### Subagent Not Working

1. Verify parent agent is active
2. Check specialty matches task
3. Ensure clear task description

### Agent Reference Not Found

1. Check `.github/agents/` directory
2. Verify agent file exists
3. Fix references in skills.md

## Resources

- [GitHub Copilot Documentation](https://docs.github.com/en/copilot)
- [Copilot Best Practices](https://docs.github.com/en/copilot/using-github-copilot/getting-started-with-github-copilot)
- [Agent Configuration Guide](../.github/agents/README.md)

## Feedback

If skills or subagents aren't working as expected:
1. Review this documentation
2. Check agent definitions
3. Update configurations as needed
4. Test with various tasks
