package broker

import (
	"github.com/oneweave/event-api/lib"
)

type BrokerUpdateFailedData struct {
	BrokerUpdateBaseData `json:",inline" yaml:",inline"`
	lib.EventFailure     `json:",inline" yaml:",inline"`
}

func NewBrokerUpdateFailedData() BrokerUpdateFailedData {
	return BrokerUpdateFailedData{
		BrokerUpdateBaseData: NewBrokerUpdateBaseData(),
		EventFailure:         lib.NewEventFailure(),
	}
}

type BrokerUpdateFailedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                 `json:"type" bson:"type" validate:"required,eq=broker.update.failed.v1"`
	Data         BrokerUpdateFailedData `json:"data" bson:"data" validate:"required"`
}

func NewBrokerUpdateFailedCloudEvent() BrokerUpdateFailedCloudEvent {
	return BrokerUpdateFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     BrokerUpdateFailedV1Type,
		Data:     NewBrokerUpdateFailedData(),
	}
}
