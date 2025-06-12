// saas/pkg/a2a/protocol_test.go
package a2a_test

import (
    "testing"
    "github.com/A2AGateway/a2agateway/saas/pkg/a2a"
)

func TestCreateSendTaskRequest(t *testing.T) {
    // Create a protocol instance
    protocol := a2a.NewProtocol()
    
    // Create a message for the task
    textPart := a2a.NewTextPart("Hello from the test!")
    message := a2a.NewMessage(a2a.RoleUser, []a2a.Part{textPart})
    
    // Create parameters for the request
    params := a2a.TaskSendParams{
        ID:      "test-task-123",
        Message: *message,
    }
    
    // Create the request
    request := protocol.CreateSendTaskRequest("req-123", params)
    
    // Verify the request
    if request.JSONRPC != "2.0" {
        t.Errorf("JSONRPC version mismatch: expected %s, got %s", "2.0", request.JSONRPC)
    }
    
    if request.Method != "tasks/send" {
        t.Errorf("Method mismatch: expected %s, got %s", "tasks/send", request.Method)
    }
    
    if request.ID != "req-123" {
        t.Errorf("ID mismatch: expected %s, got %s", "req-123", request.ID)
    }
    
    // Serialize to JSON to verify it produces valid JSON
    json, err := request.ToJSON()
    if err != nil {
        t.Fatalf("Failed to serialize request: %v", err)
    }
    
    // Deserialize to verify the round trip
    parsedRequest, err := a2a.RequestFromJSON(json)
    if err != nil {
        t.Fatalf("Failed to deserialize request: %v", err)
    }
    
    if parsedRequest.Method != "tasks/send" {
        t.Errorf("Parsed method mismatch: expected %s, got %s", "tasks/send", parsedRequest.Method)
    }
}