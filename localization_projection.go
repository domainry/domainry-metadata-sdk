package metadatasdk

import "context"

// LocalizedTextResourceSnapshot replaces the localized text projection for
// one source-owned resource inside one Workspace.
type LocalizedTextResourceSnapshot struct {
	WorkspaceID string
	EntityType  string
	EntityKey   string
	SourceKind  string
	SourceID    string
	Values      []LocalizedText
}

// LocalizationProjection is the narrow write boundary used by definition
// owners that participate in the host transaction. Metadata remains the sole
// owner of localized-text persistence.
type LocalizationProjection interface {
	ReplaceResource(context.Context, LocalizedTextResourceSnapshot) error
}

// LocalizationProjectionBinding is optional on the base Binding so older
// consumers that only read localization do not acquire a write capability.
type LocalizationProjectionBinding interface {
	LocalizationProjection() LocalizationProjection
}
