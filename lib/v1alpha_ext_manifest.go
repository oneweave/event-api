package lib

import (
	"encoding/json"
	"fmt"
)

const (
	extPluginManifestKind = "Manifest"
)

// When adding new fields to the manifest,
// - ensure that the default values are set in NewExtPluginManifestSpec() and that validation is added in ValidateStruct().
// - add the fields also to the InternalManifest v1alphamanifest.go

type ExtPluginManifestSpec struct {
	Compatibility   *PluginManifestCompatibility `json:"compatibility,omitempty" bson:"compatibility,omitempty"`
	Interfaces      *PluginManifestInterfaces    `json:"interfaces,omitempty" bson:"interfaces,omitempty"`
	Dependencies    *PluginManifestDependencies  `json:"dependencies,omitempty" bson:"dependencies,omitempty"`
	Permissions     *PluginManifestPermissions   `json:"permissions,omitempty" bson:"permissions,omitempty"`
	Configuration   *PluginManifestConfiguration `json:"configuration,omitempty" bson:"configuration,omitempty"`
	ImagePullTarget *ReleaseImagePullTarget      `json:"imagePullTarget,omitempty" bson:"image_pull_target,omitempty"`
	Observability   *PluginManifestObservability `json:"observability,omitempty" bson:"observability,omitempty"`
}

func NewExtPluginManifestSpec() ExtPluginManifestSpec {
	compatibility := NewPluginManifestCompatibility()
	interfaces := NewPluginManifestInterfaces()
	dependencies := NewPluginManifestDependencies()
	permissions := NewPluginManifestPermissions()
	configuration := NewPluginManifestConfiguration()
	observability := NewPluginManifestObservability()
	return ExtPluginManifestSpec{
		Compatibility: &compatibility,
		Interfaces:    &interfaces,
		Dependencies:  &dependencies,
		Permissions:   &permissions,
		Configuration: &configuration,
		Observability: &observability,
	}
}

type ExtPluginManifest struct {
	APIVersion string                 `json:"apiVersion" bson:"api_version" validate:"required,eq=oneweave/v1alpha"`
	Kind       string                 `json:"kind" bson:"kind" validate:"required,eq=Manifest"`
	Metadata   PluginManifestMetadata `json:"metadata" bson:"metadata" validate:"required"`
	Spec       ExtPluginManifestSpec  `json:"spec" bson:"spec" validate:"required"`
}

func NewV1ExtAlphaPluginManifest() ExtPluginManifest {
	return ExtPluginManifest{
		APIVersion: v1alphaPluginManifestAPIVersion,
		Kind:       extPluginManifestKind,
		Metadata:   NewPluginManifestMetadata(),
		Spec:       NewExtPluginManifestSpec(),
	}
}

func NewExtPluginManifest() ExtPluginManifest {
	return NewV1ExtAlphaPluginManifest()
}

// ToInternal converts the external ExtPluginManifest to the internal PluginManifest.
// It uses a JSON round-trip marshalling trick to automatically map matching fields
// and merge them over the internal model's default values (such as pre-initialized slices and configuration defaults).
func (ext *ExtPluginManifest) toInternal() (PluginManifest, error) {
	internal := NewPluginManifest()

	// Marshal external manifest to JSON
	data, err := json.Marshal(ext)
	if err != nil {
		return internal, fmt.Errorf("failed to marshal external manifest: %w", err)
	}

	// Unmarshal JSON onto the default-initialized internal manifest structure
	if err := json.Unmarshal(data, &internal); err != nil {
		return internal, fmt.Errorf("failed to unmarshal into internal manifest: %w", err)
	}

	// Ensure the Kind is set to the internal kind
	internal.Kind = pluginManifestKind

	return internal, nil
}
