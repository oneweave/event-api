package controller

import (
	"github.com/oneweave/event-api/lib"
)

type ControllerUpdateRejectedData struct {
	ControllerUpdateBaseData `json:",inline" yaml:",inline"`
	ControllerUpdateID       string `json:"controllerUpdateId" bson:"controller_update_id" validate:"required,eventid"`
}

type ControllerUpdateRejectedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                       `json:"type" bson:"type" validate:"required,eq=controller.update.rejected.v1"`
	Data         ControllerUpdateRejectedData `json:"data" bson:"data" validate:"required"`
}

func NewControllerUpdateRejectedCloudEvent() ControllerUpdateRejectedCloudEvent {
	return ControllerUpdateRejectedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ControllerUpdateRejectedV1Type,
		Data:     ControllerUpdateRejectedData{},
	}
}
