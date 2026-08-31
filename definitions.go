package metadatasdk

import (
	"context"
	"encoding/json"
)

type Definition struct {
	ResourceType  string          `json:"resource_type"`
	ResourceKey   string          `json:"resource_key"`
	ObjectKey     string          `json:"object_key,omitempty"`
	Name          string          `json:"name,omitempty"`
	Payload       json.RawMessage `json:"payload"`
	SchemaVersion string          `json:"schema_version,omitempty"`
	SchemaHash    string          `json:"schema_hash,omitempty"`
	SourceKind    string          `json:"source_kind,omitempty"`
	SourceID      string          `json:"source_id,omitempty"`
	DisabledAt    string          `json:"disabled_at,omitempty"`
	CreatedAt     string          `json:"created_at,omitempty"`
	UpdatedAt     string          `json:"updated_at,omitempty"`
}

type DefinitionQuery struct {
	ResourceType string
	SourceID     string
}

type DefinitionSnapshot struct {
	Definitions []Definition
}

type Definitions interface {
	List(context.Context, DefinitionQuery) ([]Definition, error)
	Get(context.Context, string, string) (Definition, bool, error)
	Snapshot(context.Context) (DefinitionSnapshot, error)
}

type ProjectionSnapshot struct {
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
