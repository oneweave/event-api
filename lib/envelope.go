package lib

import (
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/sixafter/nanoid"
)

const (
	EnvelopeSpecVersion       = "1.0"
	DataContentTypeJSON       = "application/json"
	CorrelationIdExtensionKey = "correlationid"
	CausationIdExtensionKey   = "causationid"
)

type Envelope struct {
	SpecVersion     string `json:"specversion" bson:"spec_version" validate:"required,eq=1.0"`
	ID              string `json:"id" bson:"id" validate:"required,uuid"`
	Source          string `json:"source" bson:"source" validate:"required"`
	Subject         string `json:"subject" bson:"subject" validate:"required"`
	Time            string `json:"time" bson:"time" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	DataContentType string `json:"datacontenttype" bson:"data_content_type" validate:"required,eq=application/json"`
	Dataschema      string `json:"dataschema" bson:"data_schema" validate:"required"`
	// correlation for cross-service tracing, reuse correlationid from cloudevents extensions
	CorrelationID string `json:"correlationid" bson:"correlation_id" validate:"required,uuid"`
	// causation for event sourcing and debugging, use event ID as causation ID for traceability
	CausationID string `json:"causationid" bson:"causation_id" validate:"required,uuid"`
}

func NewEnvelope() Envelope {
	now := time.Now().UTC().Format(time.RFC3339)
	nanoid, _ := nanoid.New()
	return Envelope{
		ID:              nanoid.String(),
		SpecVersion:     EnvelopeSpecVersion,
		DataContentType: DataContentTypeJSON,
		Time:            now,
	}
}

func NewEnvelopeFromCloudEvent(event cloudevents.Event) Envelope {
	extensions := event.Extensions()
	correlationID := extensions[CorrelationIdExtensionKey]
	causationID := event.ID()

	envelope := NewEnvelope()
	envelope.CorrelationID = correlationID.(string)
	envelope.CausationID = causationID
	// we expect type to be set by the sender
	return envelope
}
