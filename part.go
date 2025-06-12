// Package a2a implements the A2A protocol operations and data structures
// This file implements the different part types (TextPart, FilePart, DataPart) as defined in the A2A schema
package a2a

// Part represents a content part in a message or artifact
type Part interface {
    GetType() string
    GetMetadata() map[string]interface{}
}

// TextPart represents a text part
type TextPart struct {
    Type     string                 `json:"type"`
    Text     string                 `json:"text"`
    Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// GetType returns the part type
func (t TextPart) GetType() string {
    return t.Type
}

// GetMetadata returns the part metadata
func (t TextPart) GetMetadata() map[string]interface{} {
    return t.Metadata
}

// NewTextPart creates a new text part
func NewTextPart(text string) TextPart {
    return TextPart{
        Type: "text",
        Text: text,
    }
}

// WithMetadata adds metadata to the text part
func (t TextPart) WithMetadata(metadata map[string]interface{}) TextPart {
    t.Metadata = metadata
    return t
}

// FilePart represents a file part
type FilePart struct {
    Type     string                 `json:"type"`
    File     FileContent            `json:"file"`
    Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// GetType returns the part type
func (f FilePart) GetType() string {
    return f.Type
}

// GetMetadata returns the part metadata
func (f FilePart) GetMetadata() map[string]interface{} {
    return f.Metadata
}

// FileContent represents the content of a file, either as base64 encoded bytes or a URI.
// Either 'bytes' or 'uri' must be provided, but not both.
type FileContent struct {
    Name     string  `json:"name,omitempty"`
    MimeType string  `json:"mimeType,omitempty"`
    Bytes    string  `json:"bytes,omitempty"`
    URI      string  `json:"uri,omitempty"`
}

// NewFilePart creates a new file part
func NewFilePart(file FileContent) FilePart {
    return FilePart{
        Type: "file",
        File: file,
    }
}

// ValidateFileContent validates that the FileContent follows the constraints
// (either 'bytes' or 'uri' must be provided, but not both)
func ValidateFileContent(file FileContent) bool {
    // Must have either bytes or URI but not both
    hasBoth := file.Bytes != "" && file.URI != ""
    hasNeither := file.Bytes == "" && file.URI == ""
    
    return !hasBoth && !hasNeither
}

// NewFileContentWithBytes creates a new FileContent with base64 encoded data
func NewFileContentWithBytes(name, mimeType, bytes string) FileContent {
    return FileContent{
        Name:     name,
        MimeType: mimeType,
        Bytes:    bytes,
    }
}

// NewFileContentWithURI creates a new FileContent with a URI reference
func NewFileContentWithURI(name, mimeType, uri string) FileContent {
    return FileContent{
        Name:     name,
        MimeType: mimeType,
        URI:      uri,
    }
}

// WithMetadata adds metadata to the file part
func (f FilePart) WithMetadata(metadata map[string]interface{}) FilePart {
    f.Metadata = metadata
    return f
}

// DataPart represents a data part
type DataPart struct {
    Type     string                 `json:"type"`
    Data     map[string]interface{} `json:"data"`
    Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// GetType returns the part type
func (d DataPart) GetType() string {
    return d.Type
}

// GetMetadata returns the part metadata
func (d DataPart) GetMetadata() map[string]interface{} {
    return d.Metadata
}

// NewDataPart creates a new data part
func NewDataPart(data map[string]interface{}) DataPart {
    return DataPart{
        Type: "data",
        Data: data,
    }
}

// WithMetadata adds metadata to the data part
func (d DataPart) WithMetadata(metadata map[string]interface{}) DataPart {
    d.Metadata = metadata
    return d
}