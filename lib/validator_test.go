package lib

import (
	"encoding/json"
	"testing"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	eventid "github.com/oneweave/event-id"
	a "github.com/stretchr/testify/assert"
)

type TestStruct struct {
	ID    string `json:"id" validate:"required,eventid"`
	Email string `json:"email" validate:"required,email"`
}

func TestValidateEventID(t *testing.T) {
	assert := a.New(t)

	// Generate valid event-id
	validID, err := eventid.New("evt")
	assert.NoError(err)
	assert.NotEmpty(validID)

	// Valid cases
	t.Run("valid event-id struct validation", func(t *testing.T) {
		s := TestStruct{
			ID:    validID,
			Email: "test@example.com",
		}
		err := ValidateStruct(s)
		assert.NoError(err)
	})

	// Invalid cases
	t.Run("invalid prefix or format fails validation", func(t *testing.T) {
		invalidIDs := []string{
			"invalid-uuid-format",
			"evt_short",
			"evt_069rz3kw7dyyz2gj28t5cy4tqgi", // too long (contains i)
			"evt-069rz3kw7dyyz2gj28t5cy4tqg",  // hyphen separator instead of underscore
		}

		for _, id := range invalidIDs {
			s := TestStruct{
				ID:    id,
				Email: "test@example.com",
			}
			err := ValidateStruct(s)
			assert.Error(err, "should fail validation for ID: %s", id)
		}
	})
}

func TestParseAndValidate(t *testing.T) {
	assert := a.New(t)

	validID, err := eventid.New("evt")
	assert.NoError(err)

	t.Run("successful parse and validate", func(t *testing.T) {
		payload := TestStruct{
			ID:    validID,
			Email: "hello@oneweave.io",
		}
		event := cloudevents.NewEvent()
		event.SetID(validID)
		event.SetSource("test-source")
		event.SetType("test.event.v1")
		err := event.SetData(cloudevents.ApplicationJSON, payload)
		assert.NoError(err)

		s, err := ParseAndValidate[TestStruct](event)
		assert.NoError(err)
		assert.Equal(validID, s.ID)
		assert.Equal("hello@oneweave.io", s.Email)
	})

	t.Run("unmarshal failure", func(t *testing.T) {
		event := cloudevents.NewEvent()
		event.SetID(validID)
		event.SetSource("test-source")
		event.SetType("test.event.v1")
		_ = event.SetData(cloudevents.ApplicationJSON, "just a string")

		_, err := ParseAndValidate[TestStruct](event)
		assert.Error(err)
	})

	t.Run("validation failure", func(t *testing.T) {
		payload := TestStruct{
			ID:    "invalid_id",
			Email: "not-an-email",
		}
		event := cloudevents.NewEvent()
		event.SetID(validID)
		event.SetSource("test-source")
		event.SetType("test.event.v1")
		_ = event.SetData(cloudevents.ApplicationJSON, payload)

		_, err = ParseAndValidate[TestStruct](event)
		assert.Error(err)
	})
}

func TestPluginManifestValidation_EmptyArrays(t *testing.T) {
	assert := a.New(t)

	// Create a default valid manifest
	manifest := NewPluginManifest()
	manifest.Metadata.Namespace = "test-ns"
	manifest.Metadata.Name = "test-svc"
	manifest.Metadata.Version = "1.0.0"

	manifest.Spec.ImagePullTarget.Name = "test-image"
	manifest.Spec.ImagePullTarget.BaseURL = "gcr.io/"
	manifest.Spec.ImagePullTarget.Namespace = "my-project/"
	manifest.Spec.ImagePullTarget.Tags = []string{"latest"}

	manifest.Spec.Interfaces.REST.Public = []PluginManifestRestEndpoint{}
	manifest.Spec.Interfaces.REST.Internal = []PluginManifestRestEndpoint{}
	manifest.Spec.Interfaces.Events.Publishes = []PluginManifestEventDescriptor{}
	manifest.Spec.Interfaces.Events.Consumes = []PluginManifestEventDescriptor{}
	manifest.Spec.Dependencies.Services = []PluginManifestDependency{}

	manifest.Spec.Observability.CorrelationIdHeader = "X-Correlation-ID"

	t.Run("validates correctly as Go struct directly", func(t *testing.T) {
		err := ValidateStruct(&manifest)
		assert.NoError(err)
	})

	t.Run("validates correctly after JSON roundtrip with empty arrays", func(t *testing.T) {
		data, err := json.Marshal(manifest)
		assert.NoError(err)

		var deserialized PluginManifest
		err = json.Unmarshal(data, &deserialized)
		assert.NoError(err)

		err = ValidateStruct(&deserialized)
		assert.NoError(err)
	})
}
