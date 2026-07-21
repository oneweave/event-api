package lib

import (
	"encoding/json"
	"testing"

	a "github.com/stretchr/testify/assert"
)

func strPtr(value string) *string {
	return &value
}

func boolPtr(value bool) *bool {
	return &value
}

func TestExtPluginManifestDefaults(t *testing.T) {
	assert := a.New(t)

	manifest := NewExtPluginManifest()

	assert.Equal(v1alphaPluginManifestAPIVersion, manifest.APIVersion)
	assert.Equal(extPluginManifestKind, manifest.Kind)
	assert.Equal(PluginManifestMetadata{Annotations: map[string]string{}}, manifest.Metadata)
	assert.NotNil(manifest.Spec.Compatibility)
	assert.NotNil(manifest.Spec.Interfaces)
	assert.NotNil(manifest.Spec.Dependencies)
	assert.NotNil(manifest.Spec.Permissions)
	assert.NotNil(manifest.Spec.Configuration)
	assert.NotNil(manifest.Spec.Observability)
}

func TestExtPluginManifestSpecDefaults(t *testing.T) {
	assert := a.New(t)

	spec := NewExtPluginManifestSpec()

	assert.NotNil(spec.Compatibility)
	if assert.NotNil(spec.Compatibility) {
		assert.Equal(v1alphaPluginManifestAPIVersion, spec.Compatibility.PluginAPIVersion)
	}
	assert.NotNil(spec.Interfaces)
	assert.NotNil(spec.Dependencies)
	assert.NotNil(spec.Permissions)
	assert.NotNil(spec.Configuration)
	assert.NotNil(spec.Observability)
}

func TestExtPluginManifestValidation_Variants(t *testing.T) {
	assert := a.New(t)

	tests := []struct {
		name     string
		manifest ExtPluginManifest
	}{
		{
			name: "minimal manifest",
			manifest: func() ExtPluginManifest {
				manifest := NewExtPluginManifest()
				manifest.Metadata.Namespace = "test-ns"
				manifest.Metadata.Name = "test-plugin"
				manifest.Metadata.Version = "1.0.0"
				return manifest
			}(),
		},
		{
			name: "manifest with populated configuration",
			manifest: func() ExtPluginManifest {
				manifest := NewExtPluginManifest()
				manifest.Metadata.Namespace = "test-ns"
				manifest.Metadata.Name = "test-plugin"
				manifest.Metadata.Version = "1.0.0"
				manifest.Spec.Interfaces.REST.Public = []PluginManifestRestEndpoint{}
				manifest.Spec.Interfaces.REST.Internal = []PluginManifestRestEndpoint{}
				manifest.Spec.Interfaces.Events.Publishes = []PluginManifestEventDescriptor{}
				manifest.Spec.Interfaces.Events.Consumes = []PluginManifestEventDescriptor{}
				manifest.Spec.Dependencies.Services = []PluginManifestDependency{}
				manifest.Spec.Configuration.EnvironmentVariables = []PluginManifestEnvironmentVariable{}
				manifest.Spec.Configuration.EnvironmentVariablesFromSecrets = []PluginManifestEnvironmentVariableFromSecret{}
				return manifest
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.manifest)
			assert.NoError(err)

			var deserialized ExtPluginManifest
			err = json.Unmarshal(data, &deserialized)
			assert.NoError(err)

			err = ValidateStruct(&deserialized)
			assert.NoError(err)
		})
	}
}

func TestExtPluginManifestValidation_FullyConfigured(t *testing.T) {
	assert := a.New(t)

	manifest := NewExtPluginManifest()
	manifest.Metadata.Namespace = "example"
	manifest.Metadata.Name = "example-service"
	manifest.Metadata.Version = "v1"
	manifest.Metadata.Description = strPtr("Example plugin manifest with optional fields populated")
	manifest.Metadata.Owner = strPtr("platform-team")
	manifest.Metadata.Annotations = map[string]string{
		"env":  "dev",
		"team": "core",
	}

	manifest.Spec.Compatibility.PluginAPIVersion = v1alphaPluginManifestAPIVersion
	manifest.Spec.Interfaces.REST.Public = []PluginManifestRestEndpoint{
		{
			Path:           "/healthz",
			Methods:        []string{"GET"},
			AllowAnonymous: boolPtr(true),
		},
	}
	manifest.Spec.Interfaces.REST.Internal = []PluginManifestRestEndpoint{
		{
			Path:    "/internal/status",
			Methods: []string{"GET"},
		},
	}
	manifest.Spec.Interfaces.Events.Publishes = []PluginManifestEventDescriptor{
		{
			Name:                 "artifact.release.published.v1",
			RequiresCapabilities: []string{"oneweave.core.write"},
			Required:             boolPtr(true),
		},
	}
	manifest.Spec.Interfaces.Events.Consumes = []PluginManifestEventDescriptor{
		{
			Name:                 "artifact.release.requested.v1",
			RequiresCapabilities: []string{"oneweave.core.read"},
			Required:             boolPtr(true),
		},
	}
	manifest.Spec.Dependencies.Services = []PluginManifestDependency{
		{
			Name:     "api-gateway",
			Version:  strPtr("v1"),
			Required: true,
		},
	}
	manifest.Spec.Permissions.DataAccess.Owns = []string{"releases"}
	manifest.Spec.Permissions.DataAccess.DependsOn = []string{"artifacts"}
	manifest.Spec.Configuration.Health = PluginManifestHealthProbes{
		Liveness:  "/live",
		Readiness: "/ready",
		Startup:   "/startup",
	}
	manifest.Spec.Configuration.Replicas = ReplicaConfiguration{
		MinReplicas: 1,
		MaxReplicas: 3,
	}
	manifest.Spec.Configuration.Resources = PluginManifestResources{
		CPU:    "1",
		Memory: "512Mi",
	}
	manifest.Spec.Configuration.EnvironmentVariables = []PluginManifestEnvironmentVariable{
		{
			Key:         "LOG_LEVEL",
			Value:       strPtr("info"),
			Description: strPtr("Log verbosity for the service"),
		},
	}
	manifest.Spec.Configuration.EnvironmentVariablesFromSecrets = []PluginManifestEnvironmentVariableFromSecret{
		{
			Key:         "DB_PASSWORD",
			Secret:      "db-password-secret",
			Version:     strPtr("latest"),
			Description: strPtr("Database password secret reference"),
		},
	}
	manifest.Spec.Observability.Logs = "stdout"
	manifest.Spec.Observability.Metrics = "/metrics"
	manifest.Spec.Observability.Tracing = "otel"
	manifest.Spec.Observability.CorrelationIdHeader = "x-correlation-id"

	data, err := json.Marshal(manifest)
	assert.NoError(err)

	var deserialized ExtPluginManifest
	err = json.Unmarshal(data, &deserialized)
	assert.NoError(err)

	assert.Equal("example", deserialized.Metadata.Namespace)
	assert.Equal("example-service", deserialized.Metadata.Name)
	assert.Equal("v1", deserialized.Metadata.Version)
	if assert.NotNil(deserialized.Metadata.Description) {
		assert.Equal("Example plugin manifest with optional fields populated", *deserialized.Metadata.Description)
	}
	if assert.NotNil(deserialized.Metadata.Owner) {
		assert.Equal("platform-team", *deserialized.Metadata.Owner)
	}
	assert.Equal(map[string]string{"env": "dev", "team": "core"}, deserialized.Metadata.Annotations)

	err = ValidateStruct(&deserialized)
	assert.NoError(err)
}

func TestExtPluginManifestValidation_Invalid(t *testing.T) {
	assert := a.New(t)

	tests := []struct {
		name     string
		manifest ExtPluginManifest
	}{
		{
			name: "missing metadata namespace",
			manifest: func() ExtPluginManifest {
				manifest := NewExtPluginManifest()
				manifest.Metadata.Name = "test-plugin"
				manifest.Metadata.Version = "1.0.0"
				return manifest
			}(),
		},
		{
			name: "missing kind",
			manifest: func() ExtPluginManifest {
				manifest := NewExtPluginManifest()
				manifest.Metadata.Namespace = "test-ns"
				manifest.Metadata.Name = "test-plugin"
				manifest.Metadata.Version = "1.0.0"
				manifest.Kind = ""
				return manifest
			}(),
		},
		{
			name: "nil metadata annotations",
			manifest: func() ExtPluginManifest {
				manifest := NewExtPluginManifest()
				manifest.Metadata.Namespace = "test-ns"
				manifest.Metadata.Name = "test-plugin"
				manifest.Metadata.Version = "1.0.0"
				manifest.Metadata.Annotations = nil
				return manifest
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.manifest)
			assert.NoError(err)

			var deserialized ExtPluginManifest
			err = json.Unmarshal(data, &deserialized)
			assert.NoError(err)

			err = ValidateStruct(&deserialized)
			assert.Error(err)
		})
	}
}
