package broker

import (
	"github.com/oneweave/event-api/lib"
)

type BrokerUpdateRequestedCloudEvent struct {
	lib.Envelope
	Type string                `json:"type" bson:"type" validate:"required,eq=broker.update.requested.v1"`
	Data BrokerUpdateEventData `json:"data" bson:"data" validate:"required"`
}

func NewBrokerUpdateRequestedCloudEvent() BrokerUpdateRequestedCloudEvent {
	return BrokerUpdateRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     BrokerUpdateRequestedV1Type,
		Data:     NewBrokerUpdateEventData(),
	}
}
