package controller

import (
	"github.com/oneweave/event-api/lib"
)

type ControllerUpdateFailedData struct {
	ControllerUpdateBaseData `json:",inline" yaml:",inline"`
	ControllerUpdateID       string `json:"controllerUpdateId" bson:"controller_update_id" validate:"required,eventid"`
	lib.EventFailure         `json:",inline" yaml:",inline"`
}

func NewControllerUpdateFailedData() ControllerUpdateFailedData {
	return ControllerUpdateFailedData{
		ControllerUpdateBaseData: ControllerUpdateBaseData{},
		EventFailure:             lib.NewEventFailure(),
	}
}

type ControllerUpdateFailedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                     `json:"type" bson:"type" validate:"required,eq=controller.update.failed.v1"`
	Data         ControllerUpdateFailedData `json:"data" bson:"data" validate:"required"`
}

func NewControllerUpdateFailedCloudEvent() ControllerUpdateFailedCloudEvent {
	return ControllerUpdateFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ControllerUpdateFailedV1Type,
		Data:     NewControllerUpdateFailedData(),
	}
}
