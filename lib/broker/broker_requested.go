package broker

import (
	"github.com/oneweave/event-api/lib"
)

type BrokerStateRequestedCloudEvent struct {
	lib.Envelope
	Type string               `json:"type" bson:"type" validate:"required,eq=broker.update.requested.v1"`
	Data BrokerStateEventData `json:"data" bson:"data" validate:"required,dive"`
}

func NewBrokerStateRequestedCloudEvent() BrokerStateRequestedCloudEvent {
	return BrokerStateRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     BrokerUpdateRequestedV1Type,
		Data:     NewBrokerStateEventData(),
	}
}
