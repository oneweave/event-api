package broker

import (
	"github.com/oneweave/event-api/lib"
)

type BrokerUpdateFailedData struct {
	BrokerUpdateEventData
	lib.EventFailure
}

func NewBrokerUpdateFailedData() BrokerUpdateFailedData {
	return BrokerUpdateFailedData{
		BrokerUpdateEventData: NewBrokerUpdateEventData(),
		EventFailure:          lib.NewEventFailure(),
	}
}

type BrokerUpdateFailedCloudEvent struct {
	lib.Envelope
	Type string                 `json:"type" bson:"type" validate:"required,eq=broker.update.failed.v1"`
	Data BrokerUpdateFailedData `json:"data" bson:"data" validate:"required,dive"`
}

func NewBrokerUpdateFailedCloudEvent() BrokerUpdateFailedCloudEvent {
	return BrokerUpdateFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     BrokerUpdateFailedV1Type,
		Data:     NewBrokerUpdateFailedData(),
	}
}
