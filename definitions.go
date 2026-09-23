package metadatasdk

import (
	"context"
	"encoding/json"
)

const (
	DefinitionOwnerMetadata     = "metadata"
	DefinitionOwnerIdentity     = "identity"
	DefinitionOwnerWorkflow     = "workflow"
	DefinitionOwnerAutomation   = "automation"
	DefinitionOwnerIntegration  = "integration"
	DefinitionOwnerReport       = "report"
	DefinitionOwnerAgent        = "agent"
	DefinitionOwnerScheduler    = "scheduler"
	DefinitionOwnerNotification = "notification"
	DefinitionOwnerLifecycle    = "lifecycle"

	// DefinitionNoCurrentVersion is the compare-and-swap token required when
	// publishing a definition that must not already exist.
	DefinitionNoCurrentVersion = "none"
)

type Definition struct {
	Owner            string          `json:"owner"`
	ResourceType     string          `json:"resource_type"`
	ResourceKey      string          `json:"resource_key"`
	CurrentVersionID string          `json:"current_version_id"`
	Status           string          `json:"status"`
	ObjectKey        string          `json:"object_key,omitempty"`
	Name             string          `json:"name,omitempty"`
	Payload          json.RawMessage `json:"payload"`
	SchemaVersion    string          `json:"schema_version,omitempty"`
	SchemaHash       string          `json:"schema_hash,omitempty"`
	SourceKind       string          `json:"source_kind,omitempty"`
	SourceID         string          `json:"source_id,omitempty"`
	PublishedAt      string          `json:"published_at,omitempty"`
	PublishedBy      string          `json:"published_by,omitempty"`
	DisabledAt       string          `json:"disabled_at,omitempty"`
	DisabledBy       string          `json:"disabled_by,omitempty"`
	CreatedAt        string          `json:"created_at,omitempty"`
	UpdatedAt        string          `json:"updated_at,omitempty"`
}

type DefinitionQuery struct {
	Owner        string
	CrossOwner   bool
	ResourceType string
	SourceID     string
}

type DefinitionSnapshot struct {
	Definitions []Definition
}

type Definitions interface {
	List(context.Context, DefinitionQuery) ([]Definition, error)
	Get(context.Context, string, string, string) (Definition, bool, error)
	Snapshot(context.Context, DefinitionQuery) (DefinitionSnapshot, error)
}

type ProjectionSnapshot struct {
	Owner         string
	SchemaVersion string
	SourceKind    string
	SourceID      string
	Name          string
	DefaultLocale string
	Definitions   []Definition
	LocalizedText []LocalizedText
}

type Projection interface {
	Sync(context.Context, ProjectionSnapshot) error
}

type DefinitionPublishCommand struct {
	Owner                    string
	ResourceType             string
	ResourceKey              string
	ExpectedCurrentVersionID string
	SchemaVersion            string
	SchemaHash               string
	ObjectKey                string
	Name                     string
	Payload                  json.RawMessage
	SourceKind               string
	SourceID                 string
	PublishedBy              string
}

type DefinitionPublishResult struct {
	Definition       Definition
	CurrentVersionID string
}

type DefinitionDisableCommand struct {
	Owner                    string
	ResourceType             string
	ResourceKey              string
	ExpectedCurrentVersionID string
	DisabledBy               string
}

type DefinitionVersionQuery struct {
	Owner         string
	ResourceType  string
	ResourceKey   string
	VersionID     string
	SchemaVersion string
}

type DefinitionVersion struct {
	ID            string          `json:"id"`
	Owner         string          `json:"owner"`
	ResourceType  string          `json:"resource_type"`
	ResourceKey   string          `json:"resource_key"`
	SchemaVersion string          `json:"schema_version"`
	SchemaHash    string          `json:"schema_hash"`
	Payload       json.RawMessage `json:"payload"`
	CreatedAt     string          `json:"created_at"`
}

// DefinitionStore is the deployment-neutral host port for current reads,
// source-owned replacement, CAS publication, disable and immutable history.
type DefinitionStore interface {
	Definitions
	ReplaceSourceSnapshot(context.Context, ProjectionSnapshot) error
	Publish(context.Context, DefinitionPublishCommand) (DefinitionPublishResult, error)
	Disable(context.Context, DefinitionDisableCommand) error
	GetVersion(context.Context, DefinitionVersionQuery) (DefinitionVersion, bool, error)
}
