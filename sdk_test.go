package metadatasdk

import "testing"

func TestDescriptorRequiresCompleteBusinessBinding(t *testing.T) {
	valid := Descriptor{ProtocolVersion: ProtocolVersionV1, Mode: DeploymentModeModule, Capabilities: []string{"definitions", "localization", "dictionaries", "projection"}}
	if err := valid.Validate(); err != nil {
		t.Fatal(err)
	}
	invalid := valid
	invalid.Capabilities = invalid.Capabilities[:3]
	if err := invalid.Validate(); err == nil {
		t.Fatal("descriptor without projection capability was accepted")
	}
}

func TestApplicationRefRequiresStableInstallationIdentity(t *testing.T) {
	if err := (ApplicationRef{}).Validate(); err == nil {
		t.Fatal("empty installation identity was accepted")
	}
	if err := (ApplicationRef{InstallationID: "runtime"}).Validate(); err != nil {
		t.Fatal(err)
	}
}
