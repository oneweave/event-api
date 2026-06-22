package lib

import (
	"testing"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	eventid "github.com/oneweave/event-id"
	a "github.com/stretchr/testify/assert"
)

func TestNewCloudEventFromEnvelope_Success(t *testing.T) {
	assert := a.New(t)

	validID, err := eventid.New("evt")
	assert.NoError(err)

	correlationID, err := eventid.New("evt")
	assert.NoError(err)

	dataSchema := "https://example.com/schemas/test"
	envelope := Envelope{
		SpecVersion:     EnvelopeSpecVersion,
		ID:              validID,
		Source:          "https://example.com/events",
		Subject:         "test-subject",
		Time:            time.Now().UTC().Format(time.RFC3339),
		DataContentType: cloudevents.ApplicationJSON,
		Dataschema:      dataSchema,
		CorrelationID:   correlationID,
		CausationID:     validID,
	}

	payload := map[string]string{"message": "hello"}

	event, err := NewCloudEventFromEnvelope(&envelope, "test.event.v1", payload)
	assert.NoError(err)
	assert.NotNil(event)
	assert.Equal("test.event.v1", event.Type())
	assert.Equal("https://example.com/events", event.Source())
	assert.Equal("test-subject", event.Subject())
	assert.Equal(cloudevents.ApplicationJSON, event.DataContentType())
	assert.Equal(correlationID, event.Extensions()[CorrelationIdExtensionKey])
	assert.Equal(validID, event.Extensions()[CausationIdExtensionKey])

	var actualPayload map[string]string
	assert.NoError(event.DataAs(&actualPayload))
	assert.Equal(payload, actualPayload)
}

func TestNewCloudEventFromEnvelope_InvalidTime(t *testing.T) {
	assert := a.New(t)

	validID, err := eventid.New("evt")
	assert.NoError(err)

	dataSchema := "https://example.com/schemas/test"
	envelope := Envelope{
		SpecVersion:     EnvelopeSpecVersion,
		ID:              validID,
		Source:          "https://example.com/events",
		Subject:         "test-subject",
		Time:            "not-a-valid-time",
		DataContentType: cloudevents.ApplicationJSON,
		Dataschema:      dataSchema,
		CorrelationID:   validID,
		CausationID:     validID,
	}

	_, err = NewCloudEventFromEnvelope(&envelope, "test.event.v1", map[string]string{"message": "bad"})
	assert.Error(err)
}

func TestCloudEventBuilder_Build_UsesEnvelopeSourceWhenUnset(t *testing.T) {
	assert := a.New(t)

	validID, err := eventid.New("evt")
	assert.NoError(err)

	dataSchema := "https://example.com/schemas/test"
	envelope := Envelope{
		SpecVersion:     EnvelopeSpecVersion,
		ID:              validID,
		Source:          "https://example.com/default-source",
		Subject:         "builder-subject",
		Time:            time.Now().UTC().Format(time.RFC3339),
		DataContentType: cloudevents.ApplicationJSON,
		Dataschema:      dataSchema,
		CorrelationID:   validID,
		CausationID:     validID,
	}

	builder := NewCloudEventBuilder(&envelope).
		WithEventType("builder.event.v1").
		WithPayload(map[string]string{"status": "ok"})

	event, err := builder.Build()
	assert.NoError(err)
	assert.NotNil(event)
	assert.Equal("https://example.com/default-source", event.Source())
}

func TestCloudEventBuilder_WithEventSource_OverridesSource(t *testing.T) {
	assert := a.New(t)

	validID, err := eventid.New("evt")
	assert.NoError(err)

	dataSchema := "https://example.com/schemas/test"
	envelope := Envelope{
		SpecVersion:     EnvelopeSpecVersion,
		ID:              validID,
		Source:          "https://example.com/default-source",
		Subject:         "builder-subject",
		Time:            time.Now().UTC().Format(time.RFC3339),
		DataContentType: cloudevents.ApplicationJSON,
		Dataschema:      dataSchema,
		CorrelationID:   validID,
		CausationID:     validID,
	}

	builder := NewCloudEventBuilder(&envelope).
		WithEventSource("https://example.com/override-source").
		WithEventType("builder.event.v1").
		WithPayload(map[string]string{"status": "ok"})

	event, err := builder.Build()
	assert.NoError(err)
	assert.NotNil(event)
	assert.Equal("https://example.com/override-source", event.Source())
}
