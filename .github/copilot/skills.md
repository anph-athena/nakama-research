# GitHub Copilot Skills Configuration

This file defines custom skills for GitHub Copilot to enhance its capabilities for this Nakama research project.

## Available Skills

### 1. Nakama Development Skill
**ID**: `nakama-dev`
**Description**: Expert knowledge in Nakama game server development
**Capabilities**:
- Implement Nakama runtime modules in Go
- Design matchmaking systems
- Configure Nakama server settings
- Integrate client-server communication
- Implement authentication and authorization

**Context Files**:
- `main.go` - Runtime module implementation
- `local.yml` - Nakama configuration
- `docker-compose.yml` - Deployment setup

**Related Agents**: nakama-expert

---

### 2. Go Server Development Skill
**ID**: `go-server-dev`
**Description**: Expert knowledge in Go server-side development
**Capabilities**:
- Write idiomatic Go code
- Implement concurrent systems
- Optimize performance
- Handle errors properly
- Write comprehensive tests

**Context Files**:
- `*.go` - All Go source files
- `go.mod`, `go.sum` - Dependencies

**Related Agents**: go-expert

---

### 3. Matchmaking System Skill
**ID**: `matchmaking-system`
**Description**: Specialized knowledge in game matchmaking algorithms
**Capabilities**:
- Design fair matchmaking algorithms
- Implement level-based matching
- Handle concurrent matchmaking requests
- Optimize matchmaking performance
- Balance match quality vs wait time

**Context Files**:
- `main.go` - Matchmaking implementation
- `local.yml` - Matchmaker configuration

**Related Agents**: nakama-expert, testing-qa

---

### 4. Docker Deployment Skill
**ID**: `docker-deployment`
**Description**: Expert knowledge in containerized deployments
**Capabilities**:
- Configure Docker Compose services
- Set up database containers
- Implement health checks
- Manage volumes and networks
- Optimize container images

**Context Files**:
- `docker-compose.yml` - Service definitions
- `Dockerfile` - Container builds (if any)

**Related Agents**: docker-devops

---

### 5. Testing & QA Skill
**ID**: `testing-qa`
**Description**: Expert knowledge in software testing and quality assurance
**Capabilities**:
- Write unit tests
- Implement integration tests
- Create load tests
- Validate functionality
- Measure code coverage

**Context Files**:
- `*_test.go` - Test files
- `TESTING.md` - Testing documentation

**Related Agents**: testing-qa

---

## Skill Usage

Skills are automatically activated based on:
1. **File context** - Which files you're working on
2. **Task type** - What you're trying to accomplish
3. **Agent selection** - Which agent is most relevant

### Example Scenarios

**Scenario 1: Implementing new matchmaking logic**
- Active Skills: `matchmaking-system`, `nakama-dev`, `go-server-dev`
- Active Agents: nakama-expert, go-expert
- Context: `main.go`, `local.yml`

**Scenario 2: Fixing deployment issues**
- Active Skills: `docker-deployment`
- Active Agents: docker-devops
- Context: `docker-compose.yml`

**Scenario 3: Adding tests**
- Active Skills: `testing-qa`, `go-server-dev`
- Active Agents: testing-qa, go-expert
- Context: `*_test.go`, Go source files

---

## Skill Configuration

Each skill has the following properties:

```yaml
skill:
  id: "skill-identifier"
  name: "Human Readable Name"
  description: "What this skill provides"
  capabilities:
    - "Specific capability 1"
    - "Specific capability 2"
  context_files:
    - "Relevant file patterns"
  related_agents:
    - "agent-name"
  priority: "high|medium|low"
```

---

## Adding New Skills

To add a new skill:

1. Define the skill ID and description
2. List specific capabilities
3. Identify relevant context files
4. Link to related agents
5. Set priority level
6. Document usage scenarios

---

## Skill Prioritization

Skills are prioritized based on:
- **Context relevance** - How related to current files
- **Task alignment** - How well it matches the task
- **Agent activation** - Whether related agents are active
- **Manual override** - User can explicitly request skills
