package metadatasdk

import (
	"context"
	"fmt"
	"strings"

	"github.com/domainry/domainry-metadata-sdk/modulehost"
)

const ProtocolVersionV1 = "domainry-metadata-protocol-v1"

type DeploymentMode string

const DeploymentModeModule DeploymentMode = "module"

type ApplicationRef struct{ InstallationID string }

func (r ApplicationRef) Validate() error {
	if strings.TrimSpace(r.InstallationID) == "" {
		return &Error{StatusCode: 500, Code: "metadata.installation_identity_required"}
	}
	return nil
}

type Descriptor struct {
	ProtocolVersion string
	Mode            DeploymentMode
	Capabilities    []string
}

func (d Descriptor) Validate() error {
	if d.ProtocolVersion != ProtocolVersionV1 {
		return fmt.Errorf("Metadata protocol version %q is unsupported", d.ProtocolVersion)
	}
	if d.Mode != DeploymentModeModule {
		return fmt.Errorf("Metadata deployment mode %q is unsupported", d.Mode)
	}
	required := map[string]bool{"definitions": false, "definition_store": false, "localization": false, "dictionaries": false, "projection": false}
	for _, capability := range d.Capabilities {
		if _, ok := required[strings.TrimSpace(capability)]; ok {
			required[strings.TrimSpace(capability)] = true
		}
	}
	for capability, found := range required {
		if !found {
			return fmt.Errorf("Metadata capability %q is unavailable", capability)
		}
	}
	return nil
}

// Factory is selected by the outer application composition root. Embedded
// Metadata receives only host infrastructure capabilities and remains the
// sole owner of its schema and migrations.
type Factory interface {
	OpenModule(context.Context, ApplicationRef, modulehost.Host) (Binding, error)
}

type Binding interface {
	Descriptor() Descriptor
	Definitions() Definitions
	DefinitionStore() DefinitionStore
	Localization() Localization
	Dictionaries() Dictionaries
	Projection() Projection
	Close(context.Context) error
}
