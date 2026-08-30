package metadatasdk

import "context"

const ProtocolVersionV1 = "domainry-metadata-protocol-v1"

type ApplicationRef struct{ InstallationID string }
type Descriptor struct{ ProtocolVersion, Mode string }

type Binding interface {
	Descriptor() Descriptor
	Close(context.Context) error
}
