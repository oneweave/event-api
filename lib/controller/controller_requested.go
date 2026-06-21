package controller

import (
	"github.com/oneweave/event-api/lib"
)

type ControllerUpdateRequestedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                   `json:"type" bson:"type" validate:"required,eq=controller.update.requested.v1"`
	Data         ControllerUpdateBaseData `json:"data" bson:"data" validate:"required"`
}

func NewControllerUpdateRequestedCloudEvent() ControllerUpdateRequestedCloudEvent {
	return ControllerUpdateRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ControllerUpdateRequestedV1Type,
		Data:     ControllerUpdateBaseData{},
	}
}
