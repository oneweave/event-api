package broker

import (
	"github.com/oneweave/event-api/lib"
)

type BrokerUpdateRejectedData struct {
	RequestID       string               `json:"requestId" bson:"request_id" validate:"required,uuid"`
	RejectionReason *string              `json:"rejectionReason,omitempty" bson:"rejection_reason,omitempty"`
	Details         *map[string]any      `json:"details,omitempty" bson:"details,omitempty"`
	Payload         BrokerUpdateBaseData `json:"payload" bson:"payload" validate:"required"`
}

func NewBrokerUpdateRejectedData() BrokerUpdateRejectedData {
	return BrokerUpdateRejectedData{Payload: NewBrokerUpdateBaseData()}
}

type BrokerUpdateRejectedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                   `json:"type" bson:"type" validate:"required,eq=broker.update.rejected.v1"`
	Data         BrokerUpdateRejectedData `json:"data" bson:"data" validate:"required"`
}

func NewBrokerUpdateRejectedCloudEvent() BrokerUpdateRejectedCloudEvent {
	return BrokerUpdateRejectedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     BrokerUpdateRejectedV1Type,
		Data:     NewBrokerUpdateRejectedData(),
	}
}
