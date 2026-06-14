package broker

import (
	"github.com/oneweave/event-api/lib"
)

type BrokerUpdateSucceededData struct {
	BrokerUpdateEventData
	Manifest lib.PluginManifest `json:"manifest" bson:"manifest,omitempty" validate:"omitempty,dive"`
}

type BrokerUpdateSucceededCloudEvent struct {
	lib.Envelope
	Type string                    `json:"type" bson:"type" validate:"required,eq=broker.update.succeeded.v1"`
	Data BrokerUpdateSucceededData `json:"data" bson:"data" validate:"required,dive"`
}

func NewBrokerUpdateSucceededCloudEvent() BrokerUpdateSucceededCloudEvent {
	return BrokerUpdateSucceededCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     BrokerUpdateSucceededV1Type,
		Data: BrokerUpdateSucceededData{
			BrokerUpdateEventData: NewBrokerUpdateEventData(),
			Manifest:              lib.NewPluginManifest(),
		},
	}
}
