package broker

import (
	"github.com/oneweave/event-api/lib"
)

type BrokerUpdateRequestedData struct {
	BrokerUpdateBaseData `json:",inline" yaml:",inline"`
	Manifest             lib.PluginManifest `json:"manifest" bson:"manifest" validate:"required"`
}

func NewBrokerUpdateRequestedData() BrokerUpdateRequestedData {
	return BrokerUpdateRequestedData{
		BrokerUpdateBaseData: NewBrokerUpdateBaseData(),
		Manifest:             lib.NewPluginManifest(),
	}
}

type BrokerUpdateRequestedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                    `json:"type" bson:"type" validate:"required,eq=broker.update.requested.v1"`
	Data         BrokerUpdateRequestedData `json:"data" bson:"data" validate:"required"`
}

func NewBrokerUpdateRequestedCloudEvent() BrokerUpdateRequestedCloudEvent {
	return BrokerUpdateRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     BrokerUpdateRequestedV1Type,
		Data:     NewBrokerUpdateRequestedData(),
	}
}
