package broker

import (
	"github.com/oneweave/event-api/lib"
)

type BrokerStateRejectedData struct {
	RequestID       string               `json:"requestId" bson:"request_id" validate:"required,uuid"`
	RejectionReason *string              `json:"rejectionReason,omitempty" bson:"rejection_reason,omitempty"`
	Details         *map[string]any      `json:"details,omitempty" bson:"details,omitempty"`
	Payload         BrokerStateEventData `json:"payload" bson:"payload" validate:"required"`
}

func NewBrokerStateRejectedData() BrokerStateRejectedData {
	return BrokerStateRejectedData{Payload: NewBrokerStateEventData()}
}

type BrokerStateRejectedCloudEvent struct {
	lib.Envelope
	Type string                  `json:"type" bson:"type" validate:"required,eq=broker.update.rejected.v1"`
	Data BrokerStateRejectedData `json:"data" bson:"data" validate:"required,dive"`
}

func NewBrokerStateRejectedCloudEvent() BrokerStateRejectedCloudEvent {
	return BrokerStateRejectedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     BrokerUpdateRejectedV1Type,
		Data:     NewBrokerStateRejectedData(),
	}
}
