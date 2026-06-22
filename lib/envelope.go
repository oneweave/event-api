package lib

import (
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	eventid "github.com/oneweave/event-id"
)

const (
	Prefix                    = "evt"
	EnvelopeSpecVersion       = "1.0"
	DataContentTypeJSON       = "application/json"
	CorrelationIdExtensionKey = "correlationid"
	CausationIdExtensionKey   = "causationid"
)

type Envelope struct {
	SpecVersion     string `json:"specversion" bson:"spec_version" validate:"required,eq=1.0"`
	ID              string `json:"id" bson:"id" validate:"required,eventid"`
	Source          string `json:"source" bson:"source" validate:"required"`
	Subject         string `json:"subject" bson:"subject" validate:"required"`
	Time            string `json:"time" bson:"time" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	DataContentType string `json:"datacontenttype" bson:"data_content_type" validate:"required,eq=application/json"`
	Dataschema      string `json:"dataschema" bson:"data_schema" validate:"omitempty,uri"`
	// correlation for cross-service tracing, reuse correlationid from cloudevents extensions
	CorrelationID string `json:"correlationid" bson:"correlation_id" validate:"required,eventid"`
	// causation for event sourcing and debugging, use event ID as causation ID for traceability
	CausationID string `json:"causationid" bson:"causation_id" validate:"required,eventid"`
}

func NewEnvelope() Envelope {
	now := time.Now().UTC().Format(time.RFC3339)
	id, _ := eventid.New(Prefix)
	return Envelope{
		ID:              id,
		SpecVersion:     EnvelopeSpecVersion,
		DataContentType: DataContentTypeJSON,
		Time:            now,
	}
}

// NewEnvelopeFromCloudEvent creates a new Envelope
// by copying relevant fields from the provided CloudEvent.
// Use this when you want to create a new event
// based on an existing/incoming CloudEvent
// that is causally related to an existing event
func NewEnvelopeFromCloudEvent(event cloudevents.Event) Envelope {
	extensions := event.Extensions()
	correlationID := extensions[CorrelationIdExtensionKey]

	envelope := NewEnvelope()
	envelope.CorrelationID = correlationID.(string)
	envelope.CausationID = event.ID()
	// we expect type to be set by the sender
	return envelope
}

// NewEnvelopeFromEnvelope creates a new Envelope by copying
// relevant fields from the provided Envelope.
// Use this when you already have an existing Envelope
// and want to create a new event that is causally related to an existing event
func NewEnvelopeFromEnvelope(fromEnvelope Envelope) Envelope {
	envelope := NewEnvelope()
	envelope.CorrelationID = fromEnvelope.CorrelationID
	envelope.CausationID = fromEnvelope.CausationID
	envelope.Dataschema = fromEnvelope.Dataschema
	envelope.Source = fromEnvelope.Source
	envelope.Subject = fromEnvelope.Subject
	// we expect type to be set by the sender
	return envelope
}
