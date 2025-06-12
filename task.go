// Package a2a implements the A2A protocol operations and data structures
// This file defines task structures and operations according to the A2A protocol
package a2a

import (
    "encoding/json"
    "time"
)

// TaskState represents the current state of a task
type TaskState string

const (
    TaskStateSubmitted    TaskState = "submitted"
    TaskStateWorking      TaskState = "working"
    TaskStateInputRequired TaskState = "input-required"
    TaskStateCompleted    TaskState = "completed"
    TaskStateCanceled     TaskState = "canceled"
    TaskStateFailed       TaskState = "failed"
    TaskStateUnknown      TaskState = "unknown"
)

// TaskStatus represents the status of a task
type TaskStatus struct {
    State     TaskState `json:"state"`
    Message   *Message  `json:"message,omitempty"`
    Timestamp time.Time `json:"timestamp"`
}

// TaskStatusUpdateEvent represents a task status update event
type TaskStatusUpdateEvent struct {
    ID       string     `json:"id"`
    Status   TaskStatus `json:"status"`
    Final    bool       `json:"final,omitempty"`
    Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// TaskArtifactUpdateEvent represents a task artifact update event
type TaskArtifactUpdateEvent struct {
    ID       string   `json:"id"`
    Artifact Artifact `json:"artifact"`
    Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Task represents a task in the A2A protocol
type Task struct {
    ID        string      `json:"id"`
    SessionID *string     `json:"sessionId,omitempty"`
    Status    TaskStatus  `json:"status"`
    Artifacts []Artifact  `json:"artifacts,omitempty"`
    History   []Message   `json:"history,omitempty"`
    Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// NewTask creates a new task
func NewTask(id string, state TaskState) *Task {
    return &Task{
        ID: id,
        Status: TaskStatus{
            State:     state,
            Timestamp: time.Now(),
        },
    }
}

// WithSessionID adds a session ID to the task
func (t *Task) WithSessionID(sessionID string) *Task {
    t.SessionID = &sessionID
    return t
}

// WithMessage adds a message to the task status
func (t *Task) WithMessage(message *Message) *Task {
    t.Status.Message = message
    return t
}

// AddToHistory adds a message to the task history
func (t *Task) AddToHistory(message Message) *Task {
    if t.History == nil {
        t.History = []Message{}
    }
    t.History = append(t.History, message)
    return t
}

// AddArtifact adds an artifact to the task
func (t *Task) AddArtifact(artifact Artifact) *Task {
    if t.Artifacts == nil {
        t.Artifacts = []Artifact{}
    }
    t.Artifacts = append(t.Artifacts, artifact)
    return t
}

// ToJSON converts the task to JSON
func (t *Task) ToJSON() ([]byte, error) {
    return json.Marshal(t)
}

// TaskFromJSON parses JSON into a task
func TaskFromJSON(data []byte) (*Task, error) {
    var task Task
    err := json.Unmarshal(data, &task)
    if err != nil {
        return nil, err
    }
    return &task, nil
}