package metadatasdk

import (
	"context"
	"fmt"
	"strings"

	"github.com/domainry/domainry-foundation/modulecapability"
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
	required := map[string]bool{"definitions": false, "localization": false, "dictionaries": false, "projection": false}
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

type Error struct {
	StatusCode int
	Code       string
	Cause      error
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if strings.TrimSpace(e.Code) != "" {
		return e.Code
	}
	if e.Cause != nil {
		return e.Cause.Error()
	}
	return "metadata.error"
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Cause
}

type Binding interface {
	modulecapability.Binding
	Descriptor() Descriptor
	Definitions() Definitions
	Localization() Localization
	Dictionaries() Dictionaries
	Projection() Projection
	Close(context.Context) error
}
