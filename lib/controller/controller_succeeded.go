package controller

import (
	"github.com/oneweave/event-api/lib"
)

type ControllerUpdateSucceededEventData struct {
	ControllerUpdatedEventBaseData `json:",inline" yaml:",inline"`

	OldVersion string `json:"oldVersion" bson:"old_version" validate:"required"`
	NewVersion string `json:"newVersion" bson:"new_version" validate:"required"`
	BaseUrl    string `json:"baseUrl" bson:"base_url" validate:"required,url,startswith=http|https"`
}

type ControllerUpdateSucceededCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                             `json:"type" bson:"type" validate:"required,eq=controller.update.succeeded.v1"`
	Data         ControllerUpdateSucceededEventData `json:"data" bson:"data" validate:"required"`
}

func NewControllerUpdateSucceededCloudEvent() ControllerUpdateSucceededCloudEvent {
	return ControllerUpdateSucceededCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ControllerUpdateSucceededV1Type,
		Data:     ControllerUpdateSucceededEventData{},
	}
}
