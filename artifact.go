// Package a2a implements the A2A protocol operations and data structures
// This file implements the Artifact structure as defined in the A2A schema
package a2a

import "encoding/json"

// Artifact represents a task artifact in the A2A protocol
// This corresponds to the Artifact definition in the schema
type Artifact struct {
    Name        *string                `json:"name,omitempty"`
    Description *string                `json:"description,omitempty"`
    Parts       []Part                 `json:"parts"`
    Index       int                    `json:"index,omitempty"`
    Append      *bool                  `json:"append,omitempty"`
    LastChunk   *bool                  `json:"lastChunk,omitempty"`
    Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// NewArtifact creates a new artifact with the given parts
func NewArtifact(parts []Part) *Artifact {
    return &Artifact{
        Parts: parts,
        Index: 0,
    }
}

// WithName adds a name to the artifact
func (a *Artifact) WithName(name string) *Artifact {
    a.Name = &name
    return a
}

// WithDescription adds a description to the artifact
func (a *Artifact) WithDescription(description string) *Artifact {
    a.Description = &description
    return a
}

// WithIndex sets the index of the artifact
func (a *Artifact) WithIndex(index int) *Artifact {
    a.Index = index
    return a
}

// WithAppend sets the append flag of the artifact
func (a *Artifact) WithAppend(append bool) *Artifact {
    a.Append = &append
    return a
}

// WithLastChunk sets the lastChunk flag of the artifact
func (a *Artifact) WithLastChunk(lastChunk bool) *Artifact {
    a.LastChunk = &lastChunk
    return a
}

// WithMetadata adds metadata to the artifact
func (a *Artifact) WithMetadata(metadata map[string]interface{}) *Artifact {
    a.Metadata = metadata
    return a
}

// ToJSON converts the artifact to JSON
func (a *Artifact) ToJSON() ([]byte, error) {
    return json.Marshal(a)
}

// ArtifactFromJSON parses JSON into an artifact
func ArtifactFromJSON(data []byte) (*Artifact, error) {
    var artifact Artifact
    err := json.Unmarshal(data, &artifact)
    if err != nil {
        return nil, err
    }
    return &artifact, nil
}