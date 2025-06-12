# A2A Protocol

Open source implementation of the Agent-to-Agent (A2A) communication protocol.

## Overview

The A2A Protocol enables secure, structured communication between autonomous agents in distributed systems.

## Features

- **JSON-RPC messaging** - Standardized communication format
- **Agent registration** - Dynamic agent discovery and management
- **Task management** - Structured task creation and execution
- **Push notifications** - Real-time event delivery
- **Error handling** - Robust error propagation and recovery

## Core Components

- `protocol.go` - Core protocol implementation
- `message.go` - Message structure and handling
- `task.go` - Task management
- `agent_card.go` - Agent metadata and capabilities
- `json_rpc.go` - JSON-RPC transport layer

## Usage

```go
import "github.com/A2AGateway/a2a-protocol"

// Create a new message
msg := a2a.NewMessage("task.create", payload)

// Send via JSON-RPC
response, err := client.Send(msg)
```

## Testing

```bash
go test ./...
```

## License

MIT License - see LICENSE file for details.

## Contributing

Contributions welcome! Please read our contributing guidelines.