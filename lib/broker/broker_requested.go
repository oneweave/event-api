package broker

import (
	"github.com/oneweave/event-api/lib"
)

type BrokerUpdateRequestedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                 `json:"type" bson:"type" validate:"required,eq=broker.update.requested.v1"`
	Data         BrokerUpdatedEventData `json:"data" bson:"data" validate:"required"`
}

func NewBrokerUpdateRequestedCloudEvent() BrokerUpdateRequestedCloudEvent {
	return BrokerUpdateRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     BrokerUpdateRequestedV1Type,
		Data:     NewBrokerUpdatedEventData(),
	}
}
