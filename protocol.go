// Package a2a implements the A2A protocol operations and data structures
// This package implements the specification defined in a2a.json schema
package a2a

import (
    "encoding/json"
    "errors"
)

// Task-related request/response structs

// TaskSendParams represents parameters for tasks/send method
type TaskSendParams struct {
    ID              string                 `json:"id"`
    SessionID       string                 `json:"sessionId,omitempty"`
    Message         Message                `json:"message"`
    PushNotification *PushNotificationConfig `json:"pushNotification,omitempty"`
    HistoryLength   *int                   `json:"historyLength,omitempty"`
    Metadata        map[string]interface{} `json:"metadata,omitempty"`
}

// TaskIdParams represents parameters for methods requiring only a task ID
type TaskIdParams struct {
    ID       string                 `json:"id"`
    Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// TaskQueryParams represents parameters for tasks/get method
type TaskQueryParams struct {
    ID            string                 `json:"id"`
    HistoryLength *int                   `json:"historyLength,omitempty"`
    Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

// SendTaskRequest represents a JSON-RPC request for the tasks/send method
type SendTaskRequest struct {
    JSONRPC string        `json:"jsonrpc"`
    ID      interface{}   `json:"id,omitempty"`
    Method  string        `json:"method"`
    Params  TaskSendParams `json:"params"`
}

// SendTaskResponse represents a JSON-RPC response for the tasks/send method
type SendTaskResponse struct {
    JSONRPC string        `json:"jsonrpc"`
    ID      interface{}   `json:"id,omitempty"`
    Result  *Task         `json:"result,omitempty"`
    Error   *JSONRPCError `json:"error,omitempty"`
}

// GetTaskRequest represents a JSON-RPC request for the tasks/get method
type GetTaskRequest struct {
    JSONRPC string          `json:"jsonrpc"`
    ID      interface{}     `json:"id,omitempty"`
    Method  string          `json:"method"`
    Params  TaskQueryParams `json:"params"`
}

// GetTaskResponse represents a JSON-RPC response for the tasks/get method
type GetTaskResponse struct {
    JSONRPC string        `json:"jsonrpc"`
    ID      interface{}   `json:"id,omitempty"`
    Result  *Task         `json:"result,omitempty"`
    Error   *JSONRPCError `json:"error,omitempty"`
}

// CancelTaskRequest represents a JSON-RPC request for the tasks/cancel method
type CancelTaskRequest struct {
    JSONRPC string        `json:"jsonrpc"`
    ID      interface{}   `json:"id,omitempty"`
    Method  string        `json:"method"`
    Params  TaskIdParams  `json:"params"`
}

// CancelTaskResponse represents a JSON-RPC response for the tasks/cancel method
type CancelTaskResponse struct {
    JSONRPC string        `json:"jsonrpc"`
    ID      interface{}   `json:"id,omitempty"`
    Result  *Task         `json:"result,omitempty"`
    Error   *JSONRPCError `json:"error,omitempty"`
}

// SetTaskPushNotificationRequest represents a JSON-RPC request for the tasks/pushNotification/set method
type SetTaskPushNotificationRequest struct {
    JSONRPC string                    `json:"jsonrpc"`
    ID      interface{}               `json:"id,omitempty"`
    Method  string                    `json:"method"`
    Params  TaskPushNotificationConfig `json:"params"`
}

// SetTaskPushNotificationResponse represents a JSON-RPC response for the tasks/pushNotification/set method
type SetTaskPushNotificationResponse struct {
    JSONRPC string                    `json:"jsonrpc"`
    ID      interface{}               `json:"id,omitempty"`
    Result  *TaskPushNotificationConfig `json:"result,omitempty"`
    Error   *JSONRPCError             `json:"error,omitempty"`
}

// GetTaskPushNotificationRequest represents a JSON-RPC request for the tasks/pushNotification/get method
type GetTaskPushNotificationRequest struct {
    JSONRPC string        `json:"jsonrpc"`
    ID      interface{}   `json:"id,omitempty"`
    Method  string        `json:"method"`
    Params  TaskIdParams  `json:"params"`
}

// GetTaskPushNotificationResponse represents a JSON-RPC response for the tasks/pushNotification/get method
type GetTaskPushNotificationResponse struct {
    JSONRPC string                    `json:"jsonrpc"`
    ID      interface{}               `json:"id,omitempty"`
    Result  *TaskPushNotificationConfig `json:"result,omitempty"`
    Error   *JSONRPCError             `json:"error,omitempty"`
}

// SendTaskStreamingRequest represents a JSON-RPC request for the tasks/sendSubscribe method
type SendTaskStreamingRequest struct {
    JSONRPC string        `json:"jsonrpc"`
    ID      interface{}   `json:"id,omitempty"`
    Method  string        `json:"method"`
    Params  TaskSendParams `json:"params"`
}

// SendTaskStreamingResponse represents a JSON-RPC response for the tasks/sendSubscribe method
type SendTaskStreamingResponse struct {
    JSONRPC string        `json:"jsonrpc"`
    ID      interface{}   `json:"id,omitempty"`
    Result  interface{}   `json:"result,omitempty"`
    Error   *JSONRPCError `json:"error,omitempty"`
}

// TaskResubscriptionRequest represents a JSON-RPC request for the tasks/resubscribe method
type TaskResubscriptionRequest struct {
    JSONRPC string          `json:"jsonrpc"`
    ID      interface{}     `json:"id,omitempty"`
    Method  string          `json:"method"`
    Params  TaskQueryParams `json:"params"`
}

// Protocol implements the A2A protocol
type Protocol struct {
    // Dependencies can be added here
}

// NewProtocol creates a new A2A protocol instance
func NewProtocol() *Protocol {
    return &Protocol{}
}

// CreateSendTaskRequest creates a JSON-RPC request for tasks/send
func (p *Protocol) CreateSendTaskRequest(id interface{}, params TaskSendParams) *SendTaskRequest {
    return &SendTaskRequest{
        JSONRPC: JSONRPCVersion,
        ID:      id,
        Method:  "tasks/send",
        Params:  params,
    }
}

// CreateGetTaskRequest creates a JSON-RPC request for tasks/get
func (p *Protocol) CreateGetTaskRequest(id interface{}, params TaskQueryParams) *GetTaskRequest {
    return &GetTaskRequest{
        JSONRPC: JSONRPCVersion,
        ID:      id,
        Method:  "tasks/get",
        Params:  params,
    }
}

// CreateCancelTaskRequest creates a JSON-RPC request for tasks/cancel
func (p *Protocol) CreateCancelTaskRequest(id interface{}, params TaskIdParams) *CancelTaskRequest {
    return &CancelTaskRequest{
        JSONRPC: JSONRPCVersion,
        ID:      id,
        Method:  "tasks/cancel",
        Params:  params,
    }
}

// CreateTaskStreamingRequest creates a JSON-RPC request for tasks/sendSubscribe
func (p *Protocol) CreateTaskStreamingRequest(id interface{}, params TaskSendParams) *SendTaskStreamingRequest {
    return &SendTaskStreamingRequest{
        JSONRPC: JSONRPCVersion,
        ID:      id,
        Method:  "tasks/sendSubscribe",
        Params:  params,
    }
}

// CreateTaskResubscriptionRequest creates a JSON-RPC request for tasks/resubscribe
func (p *Protocol) CreateTaskResubscriptionRequest(id interface{}, params TaskQueryParams) *TaskResubscriptionRequest {
    return &TaskResubscriptionRequest{
        JSONRPC: JSONRPCVersion,
        ID:      id,
        Method:  "tasks/resubscribe",
        Params:  params,
    }
}

// CreateSetTaskPushNotificationRequest creates a JSON-RPC request for tasks/pushNotification/set
func (p *Protocol) CreateSetTaskPushNotificationRequest(id interface{}, params TaskPushNotificationConfig) *SetTaskPushNotificationRequest {
    return &SetTaskPushNotificationRequest{
        JSONRPC: JSONRPCVersion,
        ID:      id,
        Method:  "tasks/pushNotification/set",
        Params:  params,
    }
}

// CreateGetTaskPushNotificationRequest creates a JSON-RPC request for tasks/pushNotification/get
func (p *Protocol) CreateGetTaskPushNotificationRequest(id interface{}, params TaskIdParams) *GetTaskPushNotificationRequest {
    return &GetTaskPushNotificationRequest{
        JSONRPC: JSONRPCVersion,
        ID:      id,
        Method:  "tasks/pushNotification/get",
        Params:  params,
    }
}

// ParseResponse parses a JSON-RPC response
func (p *Protocol) ParseResponse(data []byte) (*JSONRPCResponse, error) {
    var response JSONRPCResponse
    err := json.Unmarshal(data, &response)
    if err != nil {
        return nil, err
    }
    
    if response.Error != nil {
        return &response, errors.New(response.Error.Message)
    }
    
    return &response, nil
}

// ParseSendTaskResponse parses a SendTaskResponse from JSON
func (p *Protocol) ParseSendTaskResponse(data []byte) (*SendTaskResponse, error) {
    var response SendTaskResponse
    err := json.Unmarshal(data, &response)
    if err != nil {
        return nil, err
    }
    
    if response.Error != nil {
        return &response, errors.New(response.Error.Message)
    }
    
    return &response, nil
}

// ParseGetTaskResponse parses a GetTaskResponse from JSON
func (p *Protocol) ParseGetTaskResponse(data []byte) (*GetTaskResponse, error) {
    var response GetTaskResponse
    err := json.Unmarshal(data, &response)
    if err != nil {
        return nil, err
    }
    
    if response.Error != nil {
        return &response, errors.New(response.Error.Message)
    }
    
    return &response, nil
}

// ParseCancelTaskResponse parses a CancelTaskResponse from JSON
func (p *Protocol) ParseCancelTaskResponse(data []byte) (*CancelTaskResponse, error) {
    var response CancelTaskResponse
    err := json.Unmarshal(data, &response)
    if err != nil {
        return nil, err
    }
    
    if response.Error != nil {
        return &response, errors.New(response.Error.Message)
    }
    
    return &response, nil
}

// ParseSetTaskPushNotificationResponse parses a SetTaskPushNotificationResponse from JSON
func (p *Protocol) ParseSetTaskPushNotificationResponse(data []byte) (*SetTaskPushNotificationResponse, error) {
    var response SetTaskPushNotificationResponse
    err := json.Unmarshal(data, &response)
    if err != nil {
        return nil, err
    }
    
    if response.Error != nil {
        return &response, errors.New(response.Error.Message)
    }
    
    return &response, nil
}

// ParseGetTaskPushNotificationResponse parses a GetTaskPushNotificationResponse from JSON
func (p *Protocol) ParseGetTaskPushNotificationResponse(data []byte) (*GetTaskPushNotificationResponse, error) {
    var response GetTaskPushNotificationResponse
    err := json.Unmarshal(data, &response)
    if err != nil {
        return nil, err
    }
    
    if response.Error != nil {
        return &response, errors.New(response.Error.Message)
    }
    
    return &response, nil
}

// ParseTask parses a Task from a JSON-RPC response
func (p *Protocol) ParseTask(response *JSONRPCResponse) (*Task, error) {
    if response.Error != nil {
        return nil, errors.New(response.Error.Message)
    }
    
    taskBytes, err := json.Marshal(response.Result)
    if err != nil {
        return nil, err
    }
    
    var task Task
    err = json.Unmarshal(taskBytes, &task)
    if err != nil {
        return nil, err
    }
    
    return &task, nil
}

// ParseStreamingResponse parses a SendTaskStreamingResponse from JSON
func (p *Protocol) ParseStreamingResponse(data []byte) (*SendTaskStreamingResponse, error) {
    var response SendTaskStreamingResponse
    err := json.Unmarshal(data, &response)
    if err != nil {
        return nil, err
    }
    
    if response.Error != nil {
        return &response, errors.New(response.Error.Message)
    }
    
    return &response, nil
}

// ParseTaskStatusUpdate parses a TaskStatusUpdateEvent from a streaming response
func (p *Protocol) ParseTaskStatusUpdate(response *SendTaskStreamingResponse) (*TaskStatusUpdateEvent, error) {
    if response.Error != nil {
        return nil, errors.New(response.Error.Message)
    }
    
    eventBytes, err := json.Marshal(response.Result)
    if err != nil {
        return nil, err
    }
    
    var event TaskStatusUpdateEvent
    err = json.Unmarshal(eventBytes, &event)
    if err != nil {
        return nil, err
    }
    
    return &event, nil
}

// ParseTaskArtifactUpdate parses a TaskArtifactUpdateEvent from a streaming response
func (p *Protocol) ParseTaskArtifactUpdate(response *SendTaskStreamingResponse) (*TaskArtifactUpdateEvent, error) {
    if response.Error != nil {
        return nil, errors.New(response.Error.Message)
    }
    
    eventBytes, err := json.Marshal(response.Result)
    if err != nil {
        return nil, err
    }
    
    var event TaskArtifactUpdateEvent
    err = json.Unmarshal(eventBytes, &event)
    if err != nil {
        return nil, err
    }
    
    return &event, nil
}