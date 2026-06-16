package controller

import (
	"github.com/oneweave/event-api/lib"
)

type ControllerUpdateFailedEventData struct {
	ControllerUpdatedEventBaseData
	lib.EventFailure
}

func NewControllerUpdateFailedEventData() ControllerUpdateFailedEventData {
	return ControllerUpdateFailedEventData{
		ControllerUpdatedEventBaseData: ControllerUpdatedEventBaseData{},
		EventFailure:                   lib.NewEventFailure(),
	}
}

type ControllerUpdateFailedCloudEvent struct {
	lib.Envelope
	Type string                          `json:"type" bson:"type" validate:"required,eq=controller.update.failed.v1"`
	Data ControllerUpdateFailedEventData `json:"data" bson:"data" validate:"required"`
}

func NewControllerUpdateFailedCloudEvent() ControllerUpdateFailedCloudEvent {
	return ControllerUpdateFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ControllerUpdateFailedV1Type,
		Data:     NewControllerUpdateFailedEventData(),
	}
}
