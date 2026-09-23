package metadatasdk

import (
	"context"

	shareddefinition "github.com/domainry/domainry-foundation/definition"
)

const (
	DefinitionOwnerMetadata     = shareddefinition.OwnerMetadata
	DefinitionOwnerIdentity     = shareddefinition.OwnerIdentity
	DefinitionOwnerWorkflow     = shareddefinition.OwnerWorkflow
	DefinitionOwnerAutomation   = shareddefinition.OwnerAutomation
	DefinitionOwnerIntegration  = shareddefinition.OwnerIntegration
	DefinitionOwnerReport       = shareddefinition.OwnerReport
	DefinitionOwnerAgent        = shareddefinition.OwnerAgent
	DefinitionOwnerScheduler    = shareddefinition.OwnerScheduler
	DefinitionOwnerNotification = shareddefinition.OwnerNotification
	DefinitionOwnerLifecycle    = shareddefinition.OwnerLifecycle

	DefinitionNoCurrentVersion = shareddefinition.NoCurrentVersion
)

type Error = shareddefinition.Error
type Definition = shareddefinition.Definition
type DefinitionQuery = shareddefinition.Query
type DefinitionSnapshot = shareddefinition.Snapshot
type DefinitionPublishCommand = shareddefinition.PublishCommand
type DefinitionPublishResult = shareddefinition.PublishResult
type DefinitionDisableCommand = shareddefinition.DisableCommand
type DefinitionVersionQuery = shareddefinition.VersionQuery
type DefinitionVersion = shareddefinition.Version

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

// DefinitionStore is the SDK-facing Definition port. Its projection snapshot
// retains Metadata localization for Metadata itself; ordinary modules adapt a
// Foundation store with AdaptDefinitionStore and cannot write localization.
type DefinitionStore interface {
	Definitions
	ReplaceSourceSnapshot(context.Context, ProjectionSnapshot) error
	Publish(context.Context, DefinitionPublishCommand) (DefinitionPublishResult, error)
	Disable(context.Context, DefinitionDisableCommand) error
	GetVersion(context.Context, DefinitionVersionQuery) (DefinitionVersion, bool, error)
}

func AdaptDefinitionStore(store shareddefinition.StorePort) DefinitionStore {
	return definitionStoreAdapter{store: store}
}

type definitionStoreAdapter struct{ store shareddefinition.StorePort }

func (a definitionStoreAdapter) List(ctx context.Context, query DefinitionQuery) ([]Definition, error) {
	return a.store.List(ctx, query)
}

func (a definitionStoreAdapter) Get(ctx context.Context, owner, resourceType, key string) (Definition, bool, error) {
	return a.store.Get(ctx, owner, resourceType, key)
}

func (a definitionStoreAdapter) Snapshot(ctx context.Context, query DefinitionQuery) (DefinitionSnapshot, error) {
	return a.store.Snapshot(ctx, query)
}

func (a definitionStoreAdapter) ReplaceSourceSnapshot(ctx context.Context, snapshot ProjectionSnapshot) error {
	if len(snapshot.LocalizedText) != 0 {
		return &Error{StatusCode: 400, Code: "metadata.localized_text_requires_metadata_module"}
	}
	return a.store.ReplaceSourceSnapshot(ctx, shareddefinition.SourceSnapshot{
		Owner: snapshot.Owner, SchemaVersion: snapshot.SchemaVersion,
		SourceKind: snapshot.SourceKind, SourceID: snapshot.SourceID,
		Definitions: snapshot.Definitions,
	})
}

func (a definitionStoreAdapter) Publish(ctx context.Context, command DefinitionPublishCommand) (DefinitionPublishResult, error) {
	return a.store.Publish(ctx, command)
}

func (a definitionStoreAdapter) Disable(ctx context.Context, command DefinitionDisableCommand) error {
	return a.store.Disable(ctx, command)
}

func (a definitionStoreAdapter) GetVersion(ctx context.Context, query DefinitionVersionQuery) (DefinitionVersion, bool, error) {
	return a.store.GetVersion(ctx, query)
}

var _ DefinitionStore = definitionStoreAdapter{}
