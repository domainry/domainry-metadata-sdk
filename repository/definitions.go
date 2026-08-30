package repository

import (
	"context"
	"database/sql"
	"encoding/json"
)

type Definition struct {
	ResourceType, Key, ObjectKey, Name string
	SchemaHash                         string
	Payload                            json.RawMessage
}

type Snapshot struct {
	SchemaVersion, SourceKind, SourceID string
	Definitions                         []Definition
}

type DefinitionRepository interface {
	SyncDefinitions(context.Context, Snapshot) error
	DefinitionSnapshot(context.Context) (Snapshot, error)
}

// QueryExecutor lets an embedded Metadata module read through the host's
// transaction without exposing Metadata-owned table names to the host.
type QueryExecutor interface {
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
}

type ExecutionExecutor interface {
	QueryExecutor
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

type StoredDefinition struct {
	Definition
	SchemaVersion string
	SourceKind    string
	SourceID      string
	DisabledAt    string
	CreatedAt     string
	UpdatedAt     string
}

type ReplaceResult struct {
	Replaced    bool
	CurrentHash string
}

type DefinitionVersion struct {
	ResourceType  string
	ResourceKey   string
	SchemaVersion string
	SchemaHash    string
	Payload       json.RawMessage
	CreatedAt     string
}

type ExecutorSnapshotRepository interface {
	DefinitionSnapshotWithExecutor(context.Context, QueryExecutor) (Snapshot, error)
}

// ExecutorDefinitionRepository owns DML for Metadata definition tables while
// participating in a transaction opened by the embedding host.
type ExecutorDefinitionRepository interface {
	GetDefinitionWithExecutor(context.Context, QueryExecutor, string, string) (StoredDefinition, bool, error)
	ListDefinitionsWithExecutor(context.Context, QueryExecutor, string, string) ([]StoredDefinition, error)
	ReplaceDefinitionWithExecutor(context.Context, ExecutionExecutor, StoredDefinition, *string) (ReplaceResult, error)
	DisableDefinitionWithExecutor(context.Context, ExecutionExecutor, string, string, string, *string) (bool, error)
	CountDefinitionVersionsWithExecutor(context.Context, QueryExecutor, string, string) (int, error)
	InsertDefinitionVersionWithExecutor(context.Context, ExecutionExecutor, DefinitionVersion) error
	ListDefinitionVersionsWithExecutor(context.Context, QueryExecutor, string, string) ([]DefinitionVersion, error)
	GetDefinitionVersionWithExecutor(context.Context, QueryExecutor, string, string, string) (DefinitionVersion, bool, error)
}

type Binding interface {
	DefinitionRepository() DefinitionRepository
}
