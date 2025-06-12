// saas/pkg/a2a/message_test.go
package a2a_test

import (
    "testing"
    "github.com/A2AGateway/a2agateway/saas/pkg/a2a"
)

func TestMessageSerialization(t *testing.T) {
    // Create a message
    textPart := a2a.NewTextPart("Hello, world!")
    message := a2a.NewMessage(a2a.RoleUser, []a2a.Part{textPart})
    
    // Serialize to JSON
    json, err := message.ToJSON()
    if err != nil {
        t.Fatalf("Failed to serialize message: %v", err)
    }
    
    // Deserialize from JSON
    parsedMsg, err := a2a.MessageFromJSON(json)
    if err != nil {
        t.Fatalf("Failed to deserialize message: %v", err)
    }
    
    // Check that the parsed message matches the original
    if parsedMsg.Role != message.Role {
        t.Errorf("Role mismatch: expected %s, got %s", message.Role, parsedMsg.Role)
    }
    
    // Check that we have the right number of parts
    if len(parsedMsg.Parts) != len(message.Parts) {
        t.Errorf("Parts count mismatch: expected %d, got %d", len(message.Parts), len(parsedMsg.Parts))
    }
    
    // Get the text from the first part
    parsedTextPart, ok := parsedMsg.Parts[0].(a2a.TextPart)
    if !ok {
        t.Fatalf("First part is not a TextPart")
    }
    
    if parsedTextPart.Text != "Hello, world!" {
        t.Errorf("Text mismatch: expected %s, got %s", "Hello, world!", parsedTextPart.Text)
    }
}