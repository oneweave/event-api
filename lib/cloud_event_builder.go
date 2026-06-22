package lib

import (
	"fmt"
	"log"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func NewCloudEventFromEnvelope(envelope *Envelope, eventType string, payload interface{}) (*cloudevents.Event, error) {
	event := cloudevents.NewEvent()
	event.SetID(envelope.ID)
	event.SetSource(envelope.Source)
	event.SetSpecVersion(envelope.SpecVersion)
	event.SetType(eventType)
	event.SetSubject(envelope.Subject)
	if envelope.Time != "" {
		eventTime, err := time.Parse(time.RFC3339, envelope.Time)
		if err != nil {
			log.Printf("failed to parse event time type=%s id=%s time=%s: %v", eventType, envelope.ID, envelope.Time, err)
			return nil, fmt.Errorf("parse %s event time: %w", eventType, err)
		}
		event.SetTime(eventTime.UTC())
	}
	event.SetDataContentType(envelope.DataContentType)
	event.SetDataSchema(envelope.Dataschema)
	event.SetExtension(CorrelationIdExtensionKey, envelope.CorrelationID)
	event.SetExtension(CausationIdExtensionKey, envelope.CausationID)
	if err := event.SetData(envelope.DataContentType, payload); err != nil {
		log.Printf("failed to set event data type=%s id=%s: %v", eventType, envelope.ID, err)
		return nil, fmt.Errorf("set %s data: %w", eventType, err)
	}
	if err := event.Validate(); err != nil {
		log.Printf("failed to validate event type=%s id=%s: %v", eventType, envelope.ID, err)
		return nil, fmt.Errorf("validate %s event: %w", eventType, err)
	}
	return &event, nil
}

type CloudEventBuilder struct {
	dataschema  string
	envelope    *Envelope
	eventSource string
	eventType   string
	data        interface{}
}

func NewCloudEventBuilder(envelope *Envelope) CloudEventBuilder {
	return CloudEventBuilder{
		envelope: envelope,
	}
}

func (b CloudEventBuilder) WithPayload(data interface{}) CloudEventBuilder {
	b.data = data
	return b
}

func (b CloudEventBuilder) WithEventType(eventType string) CloudEventBuilder {
	b.eventType = eventType
	return b
}

func (b CloudEventBuilder) WithDataSchema(dataschema string) CloudEventBuilder {
	b.dataschema = dataschema
	return b
}

func (b CloudEventBuilder) WithEventSource(eventSource string) CloudEventBuilder {
	b.eventSource = eventSource
	return b
}

func (b CloudEventBuilder) Build() (*cloudevents.Event, error) {
	if b.envelope == nil {
		return nil, fmt.Errorf("envelope is required")
	}
	if b.eventType == "" {
		return nil, fmt.Errorf("event type is required")
	}
	if b.data == nil {
		return nil, fmt.Errorf("event data is required")
	}
	envelope := *b.envelope
	if b.eventSource != "" {
		envelope.Source = b.eventSource
	}
	if b.dataschema != "" {
		envelope.Dataschema = b.dataschema
	}
	return NewCloudEventFromEnvelope(&envelope, b.eventType, b.data)
}
