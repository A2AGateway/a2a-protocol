// Package a2a implements the A2A protocol operations and data structures
// This file implements the agent card functionality as defined in the A2A schema
package a2a

import "encoding/json"

// AgentAuthentication represents the authentication requirements of an agent
type AgentAuthentication struct {
    Schemes    []string `json:"schemes"`
    Credentials *string  `json:"credentials,omitempty"`
}

// AgentCapabilities represents the capabilities of an agent
type AgentCapabilities struct {
    Streaming            bool `json:"streaming,omitempty"`
    PushNotifications    bool `json:"pushNotifications,omitempty"`
    StateTransitionHistory bool `json:"stateTransitionHistory,omitempty"`
}

// AgentProvider represents the provider of an agent
type AgentProvider struct {
    Organization string  `json:"organization"`
    URL          *string `json:"url,omitempty"`
}

// AgentSkill represents a skill that an agent can perform
type AgentSkill struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Description *string  `json:"description,omitempty"`
    Tags        []string `json:"tags,omitempty"`
    Examples    []string `json:"examples,omitempty"`
    InputModes  []string `json:"inputModes,omitempty"`
    OutputModes []string `json:"outputModes,omitempty"`
}

// AgentCard represents an agent's capabilities according to the A2A protocol
type AgentCard struct {
    Name             string             `json:"name"`
    Description      *string            `json:"description,omitempty"`
    URL              string             `json:"url"`
    Provider         *AgentProvider     `json:"provider,omitempty"`
    Version          string             `json:"version"`
    DocumentationURL *string            `json:"documentationUrl,omitempty"`
    Capabilities     AgentCapabilities  `json:"capabilities"`
    Authentication   *AgentAuthentication `json:"authentication,omitempty"`
    DefaultInputModes []string          `json:"defaultInputModes,omitempty"`
    DefaultOutputModes []string         `json:"defaultOutputModes,omitempty"`
    Skills           []AgentSkill       `json:"skills"`
}

// NewAgentCard creates a new agent card with required fields
func NewAgentCard(name, url, version string, capabilities AgentCapabilities, skills []AgentSkill) *AgentCard {
    return &AgentCard{
        Name:               name,
        URL:                url,
        Version:            version,
        Capabilities:       capabilities,
        Skills:             skills,
        DefaultInputModes:  []string{"text"},
        DefaultOutputModes: []string{"text"},
    }
}

// ToJSON converts the agent card to JSON
func (a *AgentCard) ToJSON() ([]byte, error) {
    return json.Marshal(a)
}

// FromJSON parses JSON into an agent card
func FromJSON(data []byte) (*AgentCard, error) {
    var card AgentCard
    err := json.Unmarshal(data, &card)
    if err != nil {
        return nil, err
    }
    return &card, nil
}

// Validate checks if the agent card is valid according to the A2A schema
func (a *AgentCard) Validate() bool {
    // Required fields check
    if a.Name == "" || a.URL == "" || a.Version == "" || len(a.Skills) == 0 {
        return false
    }
    
    // Skills validation
    for _, skill := range a.Skills {
        if skill.ID == "" || skill.Name == "" {
            return false
        }
    }
    
    return true
}

// WithDescription adds a description to the agent card
func (a *AgentCard) WithDescription(description string) *AgentCard {
    a.Description = &description
    return a
}

// WithProvider adds provider information to the agent card
func (a *AgentCard) WithProvider(organization string, url *string) *AgentCard {
    a.Provider = &AgentProvider{
        Organization: organization,
        URL:          url,
    }
    return a
}

/*
card := NewAgentCard("MyAgent", "https://example.com/agent", "1.0", capabilities, skills).
    WithDescription("This is my agent").
    WithProvider("My Company", nil)
	*/
