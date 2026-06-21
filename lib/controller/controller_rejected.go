package controller

import (
	"github.com/oneweave/event-api/lib"
)

type ControllerUpdateRejectedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                   `json:"type" bson:"type" validate:"required,eq=controller.update.rejected.v1"`
	Data         ControllerUpdateBaseData `json:"data" bson:"data" validate:"required"`
}

func NewControllerUpdateRejectedCloudEvent() ControllerUpdateRejectedCloudEvent {
	return ControllerUpdateRejectedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ControllerUpdateRejectedV1Type,
		Data:     ControllerUpdateBaseData{},
	}
}
