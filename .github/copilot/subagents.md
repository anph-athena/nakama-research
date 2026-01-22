# GitHub Copilot Subagents Configuration

This file defines subagents that work under the main agents to provide specialized, focused assistance.

## Subagent Architecture

```
Main Agent
  └── Subagent 1
      └── Specialized Task A
  └── Subagent 2
      └── Specialized Task B
```

---

## Go Expert Subagents

### 1. Go Concurrency Subagent
**Parent Agent**: go-expert
**Specialty**: Goroutines, channels, and concurrent patterns
**Responsibilities**:
- Design concurrent systems
- Implement goroutine pools
- Handle race conditions
- Optimize channel usage
- Debug concurrency issues

**Use Cases**:
- Matchmaker concurrent ticket processing
- Real-time match state updates
- Parallel player validation

---

### 2. Go Testing Subagent
**Parent Agent**: go-expert
**Specialty**: Testing frameworks and test design
**Responsibilities**:
- Write table-driven tests
- Create test fixtures
- Mock dependencies
- Measure coverage
- Benchmark performance

**Use Cases**:
- Unit tests for matchmaking logic
- Integration tests for Nakama APIs
- Performance benchmarks

---

### 3. Go Performance Subagent
**Parent Agent**: go-expert
**Specialty**: Performance optimization
**Responsibilities**:
- Profile CPU and memory
- Optimize hot paths
- Reduce allocations
- Improve throughput
- Minimize latency

**Use Cases**:
- Optimize matchmaker performance
- Reduce server startup time
- Improve database query performance

---

## Nakama Expert Subagents

### 1. Matchmaking Subagent
**Parent Agent**: nakama-expert
**Specialty**: Matchmaking algorithms and implementation
**Responsibilities**:
- Design matchmaking criteria
- Implement custom matching logic
- Balance fairness vs speed
- Handle edge cases
- Monitor matchmaking metrics

**Use Cases**:
- Level-based matchmaking
- Skill-based matchmaking
- Region-based matching
- Party matchmaking

**Context**:
```go
// Matchmaking properties
numericProperties := map[string]float64{
    "level": playerLevel,
    "skill_rating": skillRating,
    "latency": avgLatency,
}
```

---

### 2. Storage Subagent
**Parent Agent**: nakama-expert
**Specialty**: Nakama storage and collections
**Responsibilities**:
- Design storage schema
- Implement CRUD operations
- Manage permissions
- Query collections efficiently
- Handle storage events

**Use Cases**:
- Store player profiles
- Save match history
- Manage leaderboards
- Cache game state

---

### 3. Real-time Subagent
**Parent Agent**: nakama-expert
**Specialty**: Real-time features (matches, chat, notifications)
**Responsibilities**:
- Implement match handlers
- Handle real-time messages
- Manage presence
- Broadcast updates
- Handle disconnections

**Use Cases**:
- Real-time match state
- In-game chat
- Player notifications
- Live leaderboard updates

---

## Docker & DevOps Subagents

### 1. Container Optimization Subagent
**Parent Agent**: docker-devops
**Specialty**: Container optimization and security
**Responsibilities**:
- Minimize image size
- Implement multi-stage builds
- Configure security settings
- Optimize layer caching
- Scan for vulnerabilities

**Use Cases**:
- Production-ready images
- Secure deployments
- Fast builds

---

### 2. Orchestration Subagent
**Parent Agent**: docker-devops
**Specialty**: Service orchestration and scaling
**Responsibilities**:
- Configure service dependencies
- Implement health checks
- Set up load balancing
- Manage scaling
- Handle failover

**Use Cases**:
- Multi-instance Nakama
- Database replication
- High availability setup

---

### 3. Database Management Subagent
**Parent Agent**: docker-devops
**Specialty**: CockroachDB configuration and management
**Responsibilities**:
- Configure database clusters
- Optimize queries
- Manage migrations
- Set up backups
- Monitor performance

**Use Cases**:
- Database optimization
- Migration scripts
- Backup strategies

---

## Testing & QA Subagents

### 1. Load Testing Subagent
**Parent Agent**: testing-qa
**Specialty**: Performance and load testing
**Responsibilities**:
- Design load test scenarios
- Simulate concurrent users
- Measure throughput
- Identify bottlenecks
- Generate performance reports

**Use Cases**:
- Matchmaker stress testing
- Server capacity planning
- Performance benchmarking

**Tools**:
- k6 for load testing
- Go benchmarks
- Custom test clients

---

### 2. Integration Testing Subagent
**Parent Agent**: testing-qa
**Specialty**: End-to-end and integration testing
**Responsibilities**:
- Test component integration
- Validate API contracts
- Test database operations
- Verify workflows
- Check error handling

**Use Cases**:
- Client-server integration
- Matchmaking flow testing
- Authentication testing

---

### 3. Test Automation Subagent
**Parent Agent**: testing-qa
**Specialty**: CI/CD test automation
**Responsibilities**:
- Configure test pipelines
- Automate regression tests
- Set up test environments
- Generate test reports
- Track test metrics

**Use Cases**:
- GitHub Actions workflows
- Automated PR testing
- Nightly test runs

---

## Documentation Subagents

### 1. API Documentation Subagent
**Parent Agent**: documentation
**Specialty**: API reference documentation
**Responsibilities**:
- Document RPC functions
- Describe request/response formats
- Provide code examples
- Explain error codes
- Generate API specs

**Use Cases**:
- Matchmaking API docs
- Custom RPC documentation
- Client integration guides

---

### 2. Tutorial Subagent
**Parent Agent**: documentation
**Specialty**: Tutorials and guides
**Responsibilities**:
- Write step-by-step guides
- Create code examples
- Explain concepts clearly
- Provide troubleshooting tips
- Update tutorials

**Use Cases**:
- Getting started guides
- Matchmaking tutorials
- Deployment guides

---

## Subagent Invocation

Subagents are automatically invoked when:
1. **Specific task detected** - Task matches subagent specialty
2. **Parent agent active** - Parent agent is handling the request
3. **Context matches** - Working on relevant files/features
4. **Explicit request** - User mentions specific functionality

### Example Invocation Flow

```
User: "Optimize the matchmaking performance"
  └─> Agent: nakama-expert (activated)
      └─> Subagent: matchmaking-subagent (activated)
          └─> Subagent: go-performance-subagent (activated)
              └─> Output: Optimized matchmaking code with benchmarks
```

---

## Subagent Coordination

Multiple subagents can work together:

**Example**: Implementing and testing new matchmaking feature
1. **matchmaking-subagent** - Implements the logic
2. **go-testing-subagent** - Writes unit tests
3. **load-testing-subagent** - Creates performance tests
4. **api-documentation-subagent** - Documents the API

---

## Adding New Subagents

To add a new subagent:

1. Identify parent agent
2. Define specialized area
3. List responsibilities
4. Document use cases
5. Provide code examples
6. Link to related subagents
