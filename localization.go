package metadatasdk

import "context"

type LocalizedText struct {
	WorkspaceID string `json:"workspace_id"`
	EntityType  string `json:"entity_type"`
	EntityKey   string `json:"entity_key"`
	Property    string `json:"property"`
	Locale      string `json:"locale"`
	Text        string `json:"text"`
	SourceKind  string `json:"source_kind,omitempty"`
	SourceID    string `json:"source_id,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type LocalizedTextQuery struct {
	WorkspaceID string
	EntityType  string
	EntityKey   string
	Property    string
	Locale      string
}

type LocalizedTextCoverageQuery struct {
	WorkspaceID    string
	Locale         string
	FallbackLocale string
}

type LocalizedTextCoverageItem struct {
	WorkspaceID    string `json:"workspace_id"`
	EntityType     string `json:"entity_type"`
	EntityKey      string `json:"entity_key"`
	Property       string `json:"property"`
	Locale         string `json:"locale"`
	RequestedText  string `json:"requested_text,omitempty"`
	FallbackLocale string `json:"fallback_locale,omitempty"`
	FallbackText   string `json:"fallback_text,omitempty"`
	DefaultText    string `json:"default_text,omitempty"`
	ResolvedText   string `json:"resolved_text"`
	ResolvedSource string `json:"resolved_source"`
	SourceKind     string `json:"source_kind,omitempty"`
	SourceID       string `json:"source_id,omitempty"`
	Missing        bool   `json:"missing"`
}

type LocalizedTextCoverage struct {
	WorkspaceID    string                      `json:"workspace_id"`
	Locale         string                      `json:"locale"`
	FallbackLocale string                      `json:"fallback_locale,omitempty"`
	Items          []LocalizedTextCoverageItem `json:"items"`
	MissingCount   int                         `json:"missing_count"`
	TotalCount     int                         `json:"total_count"`
}

type Localization interface {
	List(context.Context, LocalizedTextQuery) ([]LocalizedText, error)
	Coverage(context.Context, LocalizedTextCoverageQuery) (LocalizedTextCoverage, error)
}
