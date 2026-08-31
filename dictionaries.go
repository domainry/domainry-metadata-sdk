package metadatasdk

import "context"

type LocalizedTextMap map[string]map[string]string

type Dictionary struct {
	Key         string           `json:"key"`
	Name        string           `json:"name,omitempty"`
	Description string           `json:"description,omitempty"`
	I18n        LocalizedTextMap `json:"i18n,omitempty"`
	Source      string           `json:"source,omitempty"`
	Items       []DictionaryItem `json:"items,omitempty"`
	Config      map[string]any   `json:"config,omitempty"`
}

type DictionaryItem struct {
	Key         string           `json:"key"`
	Label       string           `json:"label,omitempty"`
	Description string           `json:"description,omitempty"`
	I18n        LocalizedTextMap `json:"i18n,omitempty"`
	Value       string           `json:"value,omitempty"`
	SortOrder   int              `json:"sort_order,omitempty"`
	Locale      string           `json:"locale,omitempty"`
	Status      string           `json:"status,omitempty"`
	ParentKey   string           `json:"parent_key,omitempty"`
	Color       string           `json:"color,omitempty"`
	Icon        string           `json:"icon,omitempty"`
	Tags        []string         `json:"tags,omitempty"`
	UI          map[string]any   `json:"ui,omitempty"`
	Metadata    map[string]any   `json:"metadata,omitempty"`
	Config      map[string]any   `json:"config,omitempty"`
}

type DictionaryItemsQuery struct {
	WorkspaceID   string
	DictionaryKey string
	Locale        string
}

type DictionaryItems struct {
	DictionaryKey string           `json:"dictionary_key"`
	Locale        string           `json:"locale,omitempty"`
	Version       int              `json:"version"`
	Items         []DictionaryItem `json:"items"`
	Cached        bool             `json:"cached"`
	CacheTTLMS    int64            `json:"cache_ttl_ms"`
}

type Dictionaries interface {
	Items(context.Context, DictionaryItemsQuery) (DictionaryItems, error)
}
