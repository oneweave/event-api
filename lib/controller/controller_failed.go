package controller

import (
	"github.com/oneweave/event-api/lib"
)

type ControllerUpdateFailedEventData struct {
	ControllerUpdatedEventBaseData `json:",inline" yaml:",inline"`
	lib.EventFailure               `json:",inline" yaml:",inline"`
}

func NewControllerUpdateFailedEventData() ControllerUpdateFailedEventData {
	return ControllerUpdateFailedEventData{
		ControllerUpdatedEventBaseData: ControllerUpdatedEventBaseData{},
		EventFailure:                   lib.NewEventFailure(),
	}
}

type ControllerUpdateFailedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                          `json:"type" bson:"type" validate:"required,eq=controller.update.failed.v1"`
	Data         ControllerUpdateFailedEventData `json:"data" bson:"data" validate:"required"`
}

func NewControllerUpdateFailedCloudEvent() ControllerUpdateFailedCloudEvent {
	return ControllerUpdateFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ControllerUpdateFailedV1Type,
		Data:     NewControllerUpdateFailedEventData(),
	}
}
