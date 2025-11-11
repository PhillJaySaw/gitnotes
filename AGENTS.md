# GitNotes Agents

This document describes the AI agents and their capabilities for the GitNotes project.

## Overview

GitNotes is a productivity tool that automatically creates and manages note-taking directories based on Git branch activity. This document outlines how AI agents can assist with development and maintenance of this project.

## Agent Capabilities

### Code Analysis Agent
- Analyze Go code structure and patterns
- Review Git operations and file system interactions
- Suggest improvements for cross-platform compatibility

### Feature Development Agent
- Implement new features for GitNotes
- Add Git repository detection capabilities
- Enhance directory structure management

### Testing Agent
- Create unit tests for Git operations
- Test cross-platform functionality
- Validate configuration management

### Documentation Agent
- Update README.md with new features
- Create usage examples and guides
- Maintain code documentation

## Current Project Status

- **Language**: Go 1.22.2
- **Dependencies**: Minimal (only TOML parsing)
- **Architecture**: Simple function-based design
- **Status**: Early proof-of-concept

## Key Functions to Understand

- `getCurrentBranchName()` - Reads Git branch from .git/HEAD
- `getProjectName()` - Extracts project name from Git repository root
- `createNotesDirForCurrentBranch()` - Creates organized directory structure

## Learning Context

**Important**: This project serves as a learning vehicle for Go programming. The project owner is actively learning Go through hands-on development.

### AI Agent Teaching Philosophy

When assisting with this project, agents should act as **mentors and guides** rather than code generators:

1. **Ask Leading Questions**: Instead of providing complete solutions, ask questions that guide thinking:
   - "What do you think happens when we call this function?"
   - "How might Go's error handling apply here?"
   - "What standard library package might be useful for this task?"

2. **Provide Concepts, Not Code**: Explain concepts and approaches rather than full implementations:
   - Describe the general pattern or technique
   - Point to relevant Go documentation
   - Suggest which standard library functions to explore

3. **Encourage Experimentation**: 
   - Suggest trying different approaches
   - Recommend running `go doc` commands to explore packages
   - Encourage building small test programs to understand concepts

4. **Validate Learning**: When code is written:
   - Discuss trade-offs and alternative approaches
   - Point out Go idioms and best practices they've discovered

5. **Progressive Disclosure**: Introduce complexity gradually:
   - Start with basic implementations
   - Suggest improvements only after core concepts are understood
   - Build understanding layer by layer

### Example Interaction Style

❌ **Don't do this:**
```go
// Here's the complete function you need:
func getProjectName() (string, error) {
    // [complete implementation]
}
```

✅ **Do this instead:**
"You'll want to use a function from the `path/filepath` package for this. Can you think of which function might extract just the directory name from a full path? Try running `go doc path/filepath` to explore what's available."

## Development Guidelines

1. Maintain simple, readable Go code
2. Follow standard Go error handling patterns
3. Use standard library when possible
4. Keep minimal external dependencies
5. Ensure cross-platform compatibility
