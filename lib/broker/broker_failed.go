package broker

import (
	"github.com/oneweave/event-api/lib"
)

type BrokerStateFailedData struct {
	BrokerStateEventData
	lib.EventFailure
}

func NewBrokerStateFailedData() BrokerStateFailedData {
	return BrokerStateFailedData{
		BrokerStateEventData: NewBrokerStateEventData(),
		EventFailure:         lib.NewEventFailure(),
	}
}

type BrokerStateFailedCloudEvent struct {
	lib.Envelope
	Type string                `json:"type" bson:"type" validate:"required,eq=broker.update.failed.v1"`
	Data BrokerStateFailedData `json:"data" bson:"data" validate:"required,dive"`
}

func NewBrokerStateFailedCloudEvent() BrokerStateFailedCloudEvent {
	return BrokerStateFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     BrokerUpdateFailedV1Type,
		Data:     NewBrokerStateFailedData(),
	}
}
