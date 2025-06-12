// Package a2a implements the A2A protocol operations and data structures
// This file implements the specific error types defined in the A2A schema
package a2a

// Error codes as defined in the A2A schema
const (
    ErrCodeParseError                  = -32700
    ErrCodeInvalidRequest              = -32600
    ErrCodeMethodNotFound              = -32601
    ErrCodeInvalidParams               = -32602
    ErrCodeInternalError               = -32603
    ErrCodeTaskNotFound                = -32001
    ErrCodeTaskNotCancelable           = -32002
    ErrCodePushNotificationNotSupported = -32003
    ErrCodeUnsupportedOperation        = -32004
)

// JSONRPCError represents a JSON-RPC 2.0 error
type JSONRPCError struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

// NewJSONRPCError creates a new JSON-RPC error
func NewJSONRPCError(code int, message string) *JSONRPCError {
    return &JSONRPCError{
        Code:    code,
        Message: message,
    }
}

// WithData adds data to the error
func (e *JSONRPCError) WithData(data interface{}) *JSONRPCError {
    e.Data = data
    return e
}

// JSONParseError creates a JSON parse error
func JSONParseError() *JSONRPCError {
    return &JSONRPCError{
        Code:    ErrCodeParseError,
        Message: "Invalid JSON payload",
    }
}

// InvalidRequestError creates an invalid request error
func InvalidRequestError() *JSONRPCError {
    return &JSONRPCError{
        Code:    ErrCodeInvalidRequest,
        Message: "Request payload validation error",
    }
}

// MethodNotFoundError creates a method not found error
func MethodNotFoundError() *JSONRPCError {
    return &JSONRPCError{
        Code:    ErrCodeMethodNotFound,
        Message: "Method not found",
    }
}

// InvalidParamsError creates an invalid parameters error
func InvalidParamsError() *JSONRPCError {
    return &JSONRPCError{
        Code:    ErrCodeInvalidParams,
        Message: "Invalid parameters",
    }
}

// InternalError creates an internal error
func InternalError() *JSONRPCError {
    return &JSONRPCError{
        Code:    ErrCodeInternalError,
        Message: "Internal error",
    }
}

// TaskNotFoundError creates a task not found error
func TaskNotFoundError() *JSONRPCError {
    return &JSONRPCError{
        Code:    ErrCodeTaskNotFound,
        Message: "Task not found",
    }
}

// TaskNotCancelableError creates a task not cancelable error
func TaskNotCancelableError() *JSONRPCError {
    return &JSONRPCError{
        Code:    ErrCodeTaskNotCancelable,
        Message: "Task cannot be canceled",
    }
}

// PushNotificationNotSupportedError creates a push notification not supported error
func PushNotificationNotSupportedError() *JSONRPCError {
    return &JSONRPCError{
        Code:    ErrCodePushNotificationNotSupported,
        Message: "Push Notification is not supported",
    }
}

// UnsupportedOperationError creates an unsupported operation error
func UnsupportedOperationError() *JSONRPCError {
    return &JSONRPCError{
        Code:    ErrCodeUnsupportedOperation,
        Message: "This operation is not supported",
    }
}
