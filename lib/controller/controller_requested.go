package controller

import (
	"github.com/oneweave/event-api/lib"
)

type ControllerUpdateRequestedCloudEvent struct {
	lib.Envelope
	Type string                        `json:"type" bson:"type" validate:"required,eq=controller.update.requested.v1"`
	Data ControllerUpdateEventBaseData `json:"data" bson:"data" validate:"required,dive"`
}

func NewControllerUpdateRequestedCloudEvent() ControllerUpdateRequestedCloudEvent {
	return ControllerUpdateRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ControllerUpdateRequestedV1Type,
		Data:     ControllerUpdateEventBaseData{},
	}
}
