// Package a2a implements the A2A protocol operations and data structures
// This file defines message structures used for communication between agents according to the A2A protocol
package a2a

import (
    "encoding/json"
	"fmt"
)

// MessageRole defines the sender role
type MessageRole string

const (
    RoleUser  MessageRole = "user"
    RoleAgent MessageRole = "agent"
)

// Message represents a communication between agents according to A2A protocol
type Message struct {
    Role     MessageRole             `json:"role"`
    Parts    []Part                  `json:"parts"`
    Metadata map[string]interface{}  `json:"metadata,omitempty"`
}

type rawPart struct {
    Type string `json:"type"`
}

// NewMessage creates a new message
func NewMessage(role MessageRole, parts []Part) *Message {
    return &Message{
        Role:  role,
        Parts: parts,
    }
}

// WithMetadata adds metadata to the message
func (m *Message) WithMetadata(metadata map[string]interface{}) *Message {
    m.Metadata = metadata
    return m
}

// ToJSON converts the message to JSON
func (m *Message) ToJSON() ([]byte, error) {
    return json.Marshal(m)
}

// Marshall message
func (m *Message) MarshalJSON() ([]byte, error) {
    type MessageAlias Message
    return json.Marshal((*MessageAlias)(m))
}

// MessageFromJSON parses JSON into a message
func MessageFromJSON(data []byte) (*Message, error) {
    var message struct {
        Role     MessageRole             `json:"role"`
        Parts    []json.RawMessage       `json:"parts"`
        Metadata map[string]interface{}  `json:"metadata,omitempty"`
    }
    
    err := json.Unmarshal(data, &message)
    if err != nil {
        return nil, err
    }
    
    result := &Message{
        Role:     message.Role,
        Metadata: message.Metadata,
        Parts:    make([]Part, 0, len(message.Parts)),
    }
    
    for _, rawPartData := range message.Parts {
        // First determine the type
        var r rawPart
        err := json.Unmarshal(rawPartData, &r)
        if err != nil {
            return nil, err
        }
        
        // Then unmarshal to the appropriate concrete type
        var part Part
        switch r.Type {
        case "text":
            var textPart TextPart
            err = json.Unmarshal(rawPartData, &textPart)
            part = textPart
        case "file":
            var filePart FilePart
            err = json.Unmarshal(rawPartData, &filePart)
            part = filePart
        case "data":
            var dataPart DataPart
            err = json.Unmarshal(rawPartData, &dataPart)
            part = dataPart
        default:
            return nil, fmt.Errorf("unknown part type: %s", r.Type)
        }
        
        if err != nil {
            return nil, err
        }
        
        result.Parts = append(result.Parts, part)
    }
    
    return result, nil
}