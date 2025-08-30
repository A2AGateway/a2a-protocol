# A2A Protocol - Go Implementation

The official Go implementation of the Agent2Agent (A2A) communication protocol for building distributed agent systems with seamless inter-agent communication.

## 🚀 Quick Start

### Installation

```bash
go get github.com/A2AGateway/a2a-protocol
```

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/A2AGateway/a2a-protocol"
)

func main() {
    // Create agent card
    agent := &a2aprotocol.AgentCard{
        Name:        "My Agent",
        Description: "A helpful AI agent",
        URL:         "http://localhost:8080",
        Version:     "1.0.0",
        Capabilities: a2aprotocol.AgentCapabilities{
            Streaming:               true,
            PushNotifications:       true,
            StateTransitionHistory:  true,
        },
    }

    // Create protocol handler
    handler := a2aprotocol.NewProtocolHandler(agent)

    // Handle incoming A2A messages
    handler.HandleTaskSend(func(ctx context.Context, params *a2aprotocol.TaskSendParams) (*a2aprotocol.Task, error) {
        // Process the task
        task := &a2aprotocol.Task{
            ID:     params.ID,
            Status: a2aprotocol.TaskStatus{
                State:     "completed",
                Timestamp: time.Now(),
            },
        }
        return task, nil
    })

    // Start A2A server
    log.Fatal(handler.ListenAndServe(":8080"))
}
```

## 🔌 Agent Communication

### Sending Messages to Other Agents

```go
package main

import (
    "context"
    "github.com/A2AGateway/a2a-protocol"
)

func sendMessageToAgent() {
    client := a2aprotocol.NewClient()
    
    // Create message
    message := &a2aprotocol.Message{
        Role: "user",
        Parts: []a2aprotocol.Part{
            &a2aprotocol.TextPart{
                Type: "text",
                Text: "Hello from another agent!",
            },
        },
    }
    
    // Send to agent
    task, err := client.SendTask(context.Background(), &a2aprotocol.TaskSendParams{
        ID:      "task-123",
        Message: message,
    }, "http://target-agent:8080")
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Task sent: %s, Status: %s\n", task.ID, task.Status.State)
}
```

## 📋 Core Features

- **JSON-RPC 2.0** - Standardized message transport
- **Task Management** - Complete task lifecycle support
- **Agent Discovery** - Dynamic agent registration and lookup
- **Streaming Support** - Real-time response streaming  
- **Push Notifications** - Event-driven communication
- **Type Safety** - Comprehensive Go type definitions
- **Error Handling** - Robust error propagation
- **Testing Tools** - Built-in testing utilities

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run benchmarks
go test -bench=. ./...
```

## 🔧 Advanced Configuration

```go
// Custom server configuration
config := &a2aprotocol.ServerConfig{
    Host:           "0.0.0.0",
    Port:           8080,
    ReadTimeout:    30 * time.Second,
    WriteTimeout:   30 * time.Second,
    MaxMessageSize: 10 * 1024 * 1024, // 10MB
    EnableCORS:     true,
}

server := a2aprotocol.NewServer(config, handler)
```

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md).

## 🔗 Related Projects

- **[A2A Python SDK](https://github.com/A2AGateway/a2a-python-sdk)** - Python implementation with AI integrations
- **[A2A Connector](https://github.com/A2AGateway/a2a-connector)** - Enterprise system integrations  
- **[A2A Gateway](https://github.com/A2AGateway/a2agateway)** - Complete SaaS platform

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🆘 Support

- **Documentation**: https://docs.a2agateway.com/protocol
- **GitHub Issues**: https://github.com/A2AGateway/a2a-protocol/issues
- **Community Discord**: https://discord.gg/a2agateway
- **Email**: protocol@a2agateway.com

---

Built with ❤️ by the A2AGateway team