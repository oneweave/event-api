package broker

import (
	"github.com/oneweave/event-api/lib"
)

type BrokerUpdateSucceededData struct {
	BrokerUpdatedEventData `json:",inline" yaml:",inline"`
	Manifest               lib.PluginManifest `json:"manifest" bson:"manifest,omitempty" validate:"omitempty"`
}

type BrokerUpdateSucceededCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                    `json:"type" bson:"type" validate:"required,eq=broker.update.succeeded.v1"`
	Data         BrokerUpdateSucceededData `json:"data" bson:"data" validate:"required"`
}

func NewBrokerUpdateSucceededCloudEvent() BrokerUpdateSucceededCloudEvent {
	return BrokerUpdateSucceededCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     BrokerUpdateSucceededV1Type,
		Data: BrokerUpdateSucceededData{
			BrokerUpdatedEventData: NewBrokerUpdatedEventData(),
			Manifest:               lib.NewPluginManifest(),
		},
	}
}
