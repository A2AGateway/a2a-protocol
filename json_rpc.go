// Package a2a implements the A2A protocol operations and data structures
// This file implements the base JSON-RPC 2.0 request/response structures as defined in the A2A schema
package a2a

import "encoding/json"

// JSONRPCVersion is the version of JSON-RPC used
const JSONRPCVersion = "2.0"

// JSONRPCMessage represents a JSON-RPC 2.0 message
type JSONRPCMessage struct {
    JSONRPC string      `json:"jsonrpc"`
    ID      interface{} `json:"id,omitempty"`
}

// NewJSONRPCMessage creates a new JSON-RPC message
func NewJSONRPCMessage(id interface{}) *JSONRPCMessage {
    return &JSONRPCMessage{
        JSONRPC: JSONRPCVersion,
        ID:      id,
    }
}

// JSONRPCRequest represents a JSON-RPC 2.0 request
type JSONRPCRequest struct {
    JSONRPC string      `json:"jsonrpc"`
    ID      interface{} `json:"id,omitempty"`
    Method  string      `json:"method"`
    Params  interface{} `json:"params,omitempty"`
}

// NewJSONRPCRequest creates a new JSON-RPC request
func NewJSONRPCRequest(id interface{}, method string, params interface{}) *JSONRPCRequest {
    return &JSONRPCRequest{
        JSONRPC: JSONRPCVersion,
        ID:      id,
        Method:  method,
        Params:  params,
    }
}

// ToJSON converts the request to JSON
func (r *JSONRPCRequest) ToJSON() ([]byte, error) {
    return json.Marshal(r)
}

// RequestFromJSON parses JSON into a request
func RequestFromJSON(data []byte) (*JSONRPCRequest, error) {
    var request JSONRPCRequest
    err := json.Unmarshal(data, &request)
    if err != nil {
        return nil, err
    }
    return &request, nil
}

// JSONRPCResponse represents a JSON-RPC 2.0 response
type JSONRPCResponse struct {
    JSONRPC string       `json:"jsonrpc"`
    ID      interface{}  `json:"id,omitempty"`
    Result  interface{}  `json:"result,omitempty"`
    Error   *JSONRPCError `json:"error,omitempty"`
}

// NewJSONRPCResponse creates a new JSON-RPC response
func NewJSONRPCResponse(id interface{}, result interface{}) *JSONRPCResponse {
    return &JSONRPCResponse{
        JSONRPC: JSONRPCVersion,
        ID:      id,
        Result:  result,
    }
}

// NewJSONRPCErrorResponse creates a new JSON-RPC error response
func NewJSONRPCErrorResponse(id interface{}, error *JSONRPCError) *JSONRPCResponse {
    return &JSONRPCResponse{
        JSONRPC: JSONRPCVersion,
        ID:      id,
        Error:   error,
    }
}

// ToJSON converts the response to JSON
func (r *JSONRPCResponse) ToJSON() ([]byte, error) {
    return json.Marshal(r)
}

// ResponseFromJSON parses JSON into a response
func ResponseFromJSON(data []byte) (*JSONRPCResponse, error) {
    var response JSONRPCResponse
    err := json.Unmarshal(data, &response)
    if err != nil {
        return nil, err
    }
    return &response, nil
}

// IsError checks if the response is an error
func (r *JSONRPCResponse) IsError() bool {
    return r.Error != nil
}