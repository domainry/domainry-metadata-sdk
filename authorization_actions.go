package metadatasdk

// Metadata Action keys are stable executable identities. HTTP paths are
// bindings and may change without renaming role grants.
const (
	ActionMetadataDefinitionsList          = "metadata.definitions.list"
	ActionMetadataDefinitionsGet           = "metadata.definitions.get"
	ActionMetadataLocalizedTextsList       = "metadata.localized_texts.list"
	ActionMetadataLocalizedTextsCoverage   = "metadata.localized_texts.coverage"
	ActionMetadataLocalizedTextsExportCSV  = "metadata.localized_texts.export_csv"
	ActionMetadataLocalizedTextsExportXLSX = "metadata.localized_texts.export_xlsx"
	ActionMetadataDictionaryItemsList      = "metadata.dictionary_items.list"
)

const (
	CapabilityMetadataDefinitions  = "metadata.definitions"
	CapabilityMetadataLocalization = "metadata.localization"
	CapabilityMetadataDictionaries = "metadata.dictionaries"
)
