package lib

import (
	"os"
	"path/filepath"
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestParseAndNormalizeManifest_WeaveMin(t *testing.T) {
	assert := a.New(t)

	// Read the minimal example config
	path := filepath.Join("..", "examples", "weave-min.yaml")
	data, err := os.ReadFile(path)
	assert.NoError(err, "should read weave-min.yaml successfully")

	// Print raw data
	t.Logf("Raw YAML: %s", string(data))

	// Parse and normalize
	internal, err := ParseAndConvertManifest(data)
	assert.NoError(err, "should parse and validate weave-min.yaml successfully")

	// Verify Kind and APIVersion
	assert.Equal("InternalManifest", internal.Kind)
	assert.Equal("oneweave/v1alpha", internal.APIVersion)

	// Verify Metadata fields
	assert.Equal("example", internal.Metadata.Namespace)
	assert.Equal("example-service", internal.Metadata.Name)
	assert.Equal("v1", internal.Metadata.Version)

	// Verify that arrays/slices are non-nil (normalized to empty arrays instead of nil)
	assert.NotNil(internal.Spec.Interfaces.REST.Public)
	assert.Len(internal.Spec.Interfaces.REST.Public, 0)
	assert.NotNil(internal.Spec.Interfaces.REST.Internal)
	assert.Len(internal.Spec.Interfaces.REST.Internal, 0)
	assert.NotNil(internal.Spec.Interfaces.Events.Publishes)
	assert.Len(internal.Spec.Interfaces.Events.Publishes, 0)
	assert.NotNil(internal.Spec.Interfaces.Events.Consumes)
	assert.Len(internal.Spec.Interfaces.Events.Consumes, 0)
	assert.NotNil(internal.Spec.Dependencies.Services)
	assert.Len(internal.Spec.Dependencies.Services, 0)
	assert.NotNil(internal.Spec.Permissions.DataAccess.Owns)
	assert.Len(internal.Spec.Permissions.DataAccess.Owns, 0)
	assert.NotNil(internal.Spec.Permissions.DataAccess.DependsOn)
	assert.Len(internal.Spec.Permissions.DataAccess.DependsOn, 0)
	assert.NotNil(internal.Spec.Configuration.EnvironmentVariables)
	assert.Len(internal.Spec.Configuration.EnvironmentVariables, 0)
	assert.NotNil(internal.Spec.Configuration.EnvironmentVariablesFromSecrets)
	assert.Len(internal.Spec.Configuration.EnvironmentVariablesFromSecrets, 0)

	// Verify that defaults are populated
	assert.Equal(0, internal.Spec.Configuration.Replicas.MinReplicas)
	assert.Equal(5, internal.Spec.Configuration.Replicas.MaxReplicas)
	assert.Equal("1", internal.Spec.Configuration.Resources.CPU)
	assert.Equal("512Mi", internal.Spec.Configuration.Resources.Memory)
	assert.Equal("/live", internal.Spec.Configuration.Health.Liveness)
	assert.Equal("/ready", internal.Spec.Configuration.Health.Readiness)
	assert.Equal("/startup", internal.Spec.Configuration.Health.Startup)
}

func TestParseAndNormalizeManifest_WeaveInternal(t *testing.T) {
	assert := a.New(t)

	// Read the fully configured example
	path := filepath.Join("..", "examples", "weave-max.yaml")
	data, err := os.ReadFile(path)
	assert.NoError(err, "should read weave-max.yaml successfully")

	// Parse and normalize
	internal, err := ParseAndConvertManifest(data)
	assert.NoError(err, "should parse and validate weave-max.yaml successfully")

	// Verify Kind and APIVersion
	assert.Equal("InternalManifest", internal.Kind)
	assert.Equal("oneweave/v1alpha", internal.APIVersion)

	// Verify Metadata fields are populated correctly
	assert.Equal("example", internal.Metadata.Namespace)
	assert.Equal("example-service", internal.Metadata.Name)
	assert.Equal("v1", internal.Metadata.Version)
	if assert.NotNil(internal.Metadata.Description) {
		assert.Equal("Example plugin manifest with optional fields populated", *internal.Metadata.Description)
	}
	if assert.NotNil(internal.Metadata.Owner) {
		assert.Equal("platform-team", *internal.Metadata.Owner)
	}
	assert.Equal("dev", internal.Metadata.Annotations["env"])
	assert.Equal("core", internal.Metadata.Annotations["team"])

	// Verify interfaces and dependencies are parsed
	assert.Len(internal.Spec.Interfaces.REST.Public, 1)
	assert.Equal("/healthz", internal.Spec.Interfaces.REST.Public[0].Path)
	assert.Equal([]string{"GET"}, internal.Spec.Interfaces.REST.Public[0].Methods)
	assert.True(*internal.Spec.Interfaces.REST.Public[0].AllowAnonymous)

	assert.Len(internal.Spec.Interfaces.REST.Internal, 1)
	assert.Equal("/internal/status", internal.Spec.Interfaces.REST.Internal[0].Path)
	assert.Equal([]string{"GET"}, internal.Spec.Interfaces.REST.Internal[0].Methods)

	assert.Len(internal.Spec.Interfaces.Events.Publishes, 1)
	assert.Equal("artifact.release.published.v1", internal.Spec.Interfaces.Events.Publishes[0].Name)
	assert.Equal([]string{"oneweave.core.write"}, internal.Spec.Interfaces.Events.Publishes[0].RequiresCapabilities)
	assert.True(*internal.Spec.Interfaces.Events.Publishes[0].Required)

	assert.Len(internal.Spec.Interfaces.Events.Consumes, 1)
	assert.Equal("artifact.release.requested.v1", internal.Spec.Interfaces.Events.Consumes[0].Name)
	assert.Equal([]string{"oneweave.core.read"}, internal.Spec.Interfaces.Events.Consumes[0].RequiresCapabilities)
	assert.True(*internal.Spec.Interfaces.Events.Consumes[0].Required)

	assert.Len(internal.Spec.Dependencies.Services, 1)
	assert.Equal("api-gateway", internal.Spec.Dependencies.Services[0].Name)
	assert.Equal("v1", *internal.Spec.Dependencies.Services[0].Version)
	assert.True(internal.Spec.Dependencies.Services[0].Required)

	assert.Equal([]string{"releases"}, internal.Spec.Permissions.DataAccess.Owns)
	assert.Equal([]string{"artifacts"}, internal.Spec.Permissions.DataAccess.DependsOn)

	// Verify overridden defaults in configuration
	assert.Equal(1, internal.Spec.Configuration.Replicas.MinReplicas)
	assert.Equal(3, internal.Spec.Configuration.Replicas.MaxReplicas)
	assert.Equal("1", internal.Spec.Configuration.Resources.CPU)
	assert.Equal("512Mi", internal.Spec.Configuration.Resources.Memory)

	assert.Len(internal.Spec.Configuration.EnvironmentVariables, 1)
	assert.Equal("LOG_LEVEL", internal.Spec.Configuration.EnvironmentVariables[0].Key)
	assert.Equal("info", *internal.Spec.Configuration.EnvironmentVariables[0].Value)

	assert.Len(internal.Spec.Configuration.EnvironmentVariablesFromSecrets, 1)
	assert.Equal("DB_PASSWORD", internal.Spec.Configuration.EnvironmentVariablesFromSecrets[0].Key)
	assert.Equal("db-password-secret", internal.Spec.Configuration.EnvironmentVariablesFromSecrets[0].Secret)
	if assert.NotNil(internal.Spec.Configuration.EnvironmentVariablesFromSecrets[0].Version) {
		assert.Equal("latest", *internal.Spec.Configuration.EnvironmentVariablesFromSecrets[0].Version)
	}
	if assert.NotNil(internal.Spec.Configuration.EnvironmentVariablesFromSecrets[0].Description) {
		assert.Equal("Database password secret reference", *internal.Spec.Configuration.EnvironmentVariablesFromSecrets[0].Description)
	}

	// Verify observability
	assert.Equal("stdout", internal.Spec.Observability.Logs)
	assert.Equal("/metrics", internal.Spec.Observability.Metrics)
	assert.Equal("otel", internal.Spec.Observability.Tracing)
	assert.Equal("x-correlation-id", internal.Spec.Observability.CorrelationIdHeader)
}

func TestParseAndNormalizeManifest_InvalidYAML(t *testing.T) {
	assert := a.New(t)

	// 1. Completely invalid YAML syntax
	invalidYAML := []byte(`
apiVersion: oneweave/v1alpha
kind: Manifest
metadata:
  namespace: [invalid
`)
	_, err := ParseAndConvertManifest(invalidYAML)
	assert.Error(err)

	// 2. Missing required metadata namespace
	missingNamespace := []byte(`
apiVersion: oneweave/v1alpha
kind: Manifest
metadata:
  name: example-service
  version: v1
spec: {}
`)
	_, err = ParseAndConvertManifest(missingNamespace)
	assert.Error(err)
	assert.Contains(err.Error(), "Namespace")

	// 3. Invalid replica settings (MinReplicas > MaxReplicas)
	invalidReplicas := []byte(`
apiVersion: oneweave/v1alpha
kind: Manifest
metadata:
  namespace: example
  name: example-service
  version: v1
spec:
  configuration:
    replicas:
      minReplicas: 5
      maxReplicas: 3
`)
	_, err = ParseAndConvertManifest(invalidReplicas)
	assert.Error(err)

	// 4. Invalid annotations (not a map of strings, e.g. integer value)
	invalidAnnotations := []byte(`
apiVersion: oneweave/v1alpha
kind: Manifest
metadata:
  namespace: example
  name: example-service
  version: v1
  annotations:
    env: 123
`)
	_, err = ParseAndConvertManifest(invalidAnnotations)
	assert.Error(err)
}
