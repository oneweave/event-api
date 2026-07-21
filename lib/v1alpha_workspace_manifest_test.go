package lib

import (
	"encoding/json"
	"testing"

	a "github.com/stretchr/testify/assert"
)

func workspaceStrPtr(value string) *string {
	return &value
}

func TestWorkspaceManifestDefaults(t *testing.T) {
	assert := a.New(t)

	manifest := NewWorkspaceManifest()

	assert.Equal(v1alphaPluginManifestAPIVersion, manifest.APIVersion)
	assert.Equal(workspaceManifestKind, manifest.Kind)
	assert.Equal(PluginManifestMetadata{}, manifest.Metadata)
	assert.Empty(manifest.Spec.Services)
}

func TestWorkspaceServiceDefaults(t *testing.T) {
	assert := a.New(t)

	service := NewWorkspaceService()

	assert.Equal(".", service.Path)
	assert.NotNil(service.ManifestFilePath)
	if assert.NotNil(service.ManifestFilePath) {
		assert.Equal("weave.yaml", *service.ManifestFilePath)
	}
	assert.Equal("", service.Name)
}

func TestWorkspaceManifestValidation_Variants(t *testing.T) {
	assert := a.New(t)

	tests := []struct {
		name     string
		manifest WorkspaceManifest
	}{
		{
			name: "minimal manifest",
			manifest: func() WorkspaceManifest {
				manifest := NewWorkspaceManifest()
				manifest.Metadata.Namespace = "test-ns"
				manifest.Metadata.Name = "test-workspace"
				manifest.Metadata.Version = "1.0.0"
				return manifest
			}(),
		},
		{
			name: "manifest with services",
			manifest: func() WorkspaceManifest {
				manifest := NewWorkspaceManifest()
				manifest.Metadata.Namespace = "test-ns"
				manifest.Metadata.Name = "test-workspace"
				manifest.Metadata.Version = "1.0.0"
				manifest.Spec.Services = []WorkspaceService{
					{
						Name:             "app",
						Path:             ".",
						ManifestFilePath: workspaceStrPtr("weave.yaml"),
					},
					{
						Name:             "worker",
						Path:             "services/worker",
						ManifestFilePath: workspaceStrPtr("weave.yaml"),
					},
				}
				return manifest
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.manifest)
			assert.NoError(err)

			var deserialized WorkspaceManifest
			err = json.Unmarshal(data, &deserialized)
			assert.NoError(err)

			err = ValidateStruct(&deserialized)
			assert.NoError(err)
		})
	}
}

func TestWorkspaceManifestValidation_Invalid(t *testing.T) {
	assert := a.New(t)

	tests := []struct {
		name     string
		manifest WorkspaceManifest
	}{
		{
			name: "missing metadata namespace",
			manifest: func() WorkspaceManifest {
				manifest := NewWorkspaceManifest()
				manifest.Metadata.Name = "test-workspace"
				manifest.Metadata.Version = "1.0.0"
				return manifest
			}(),
		},
		{
			name: "missing kind",
			manifest: func() WorkspaceManifest {
				manifest := NewWorkspaceManifest()
				manifest.Metadata.Namespace = "test-ns"
				manifest.Metadata.Name = "test-workspace"
				manifest.Metadata.Version = "1.0.0"
				manifest.Kind = ""
				return manifest
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.manifest)
			assert.NoError(err)

			var deserialized WorkspaceManifest
			err = json.Unmarshal(data, &deserialized)
			assert.NoError(err)

			err = ValidateStruct(&deserialized)
			assert.Error(err)
		})
	}
}
