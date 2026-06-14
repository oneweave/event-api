package controller

import (
	"github.com/oneweave/event-api/lib"
)

type ControllerUpdateRejectedCloudEvent struct {
	lib.Envelope
	Type string                        `json:"type" bson:"type" validate:"required,eq=controller.update.rejected.v1"`
	Data ControllerUpdateEventBaseData `json:"data" bson:"data" validate:"required,dive"`
}

func NewControllerUpdateRejectedCloudEvent() ControllerUpdateRejectedCloudEvent {
	return ControllerUpdateRejectedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ControllerUpdateRejectedV1Type,
		Data:     ControllerUpdateEventBaseData{},
	}
}
