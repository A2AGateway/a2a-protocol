// Package a2a implements the A2A protocol operations and data structures
// This file implements push notification functionality as defined in the A2A schema
package a2a

// PushNotificationConfig represents push notification configuration in the A2A protocol
type PushNotificationConfig struct {
    URL            string              `json:"url"`
    Token          *string             `json:"token,omitempty"`
    Authentication *AuthenticationInfo `json:"authentication,omitempty"`
}

// NewPushNotificationConfig creates a new push notification configuration
func NewPushNotificationConfig(url string) *PushNotificationConfig {
    return &PushNotificationConfig{
        URL: url,
    }
}

// WithToken adds a token to the push notification configuration
func (p *PushNotificationConfig) WithToken(token string) *PushNotificationConfig {
    p.Token = &token
    return p
}

// WithAuthentication adds authentication information to the push notification configuration
func (p *PushNotificationConfig) WithAuthentication(auth *AuthenticationInfo) *PushNotificationConfig {
    p.Authentication = auth
    return p
}

// AuthenticationInfo represents authentication information in the A2A protocol
type AuthenticationInfo struct {
    Schemes     []string `json:"schemes"`
    Credentials *string  `json:"credentials,omitempty"`
}

// NewAuthenticationInfo creates new authentication information
func NewAuthenticationInfo(schemes []string) *AuthenticationInfo {
    return &AuthenticationInfo{
        Schemes: schemes,
    }
}

// WithCredentials adds credentials to the authentication information
func (a *AuthenticationInfo) WithCredentials(credentials string) *AuthenticationInfo {
    a.Credentials = &credentials
    return a
}

// TaskPushNotificationConfig represents task push notification configuration
type TaskPushNotificationConfig struct {
    ID                   string                `json:"id"`
    PushNotificationConfig PushNotificationConfig `json:"pushNotificationConfig"`
}

// NewTaskPushNotificationConfig creates a new task push notification configuration
func NewTaskPushNotificationConfig(id string, config PushNotificationConfig) *TaskPushNotificationConfig {
    return &TaskPushNotificationConfig{
        ID:                   id,
        PushNotificationConfig: config,
    }
}