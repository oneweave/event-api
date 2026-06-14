package broker

import (
	"github.com/oneweave/event-api/lib"
)

type BrokerStateSucceededCloudEvent struct {
	lib.Envelope
	Type     string               `json:"type" bson:"type" validate:"required,eq=broker.update.succeeded.v1"`
	Data     BrokerStateEventData `json:"data" bson:"data" validate:"required,dive"`
	Manifest lib.PluginManifest   `json:"manifest" bson:"manifest,omitempty" validate:"omitempty,dive"`
}

func NewBrokerStateSucceededCloudEvent() BrokerStateSucceededCloudEvent {
	return BrokerStateSucceededCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     BrokerUpdateSucceededV1Type,
		Data:     NewBrokerStateEventData(),
		Manifest: lib.NewPluginManifest(),
	}
}
