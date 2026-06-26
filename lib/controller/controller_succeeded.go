package controller

import (
	"github.com/oneweave/event-api/lib"
)

type ControllerUpdateSucceededData struct {
	ControllerUpdateBaseData `json:",inline" yaml:",inline"`

	Tags        []string `json:"tags" bson:"tags" validate:"required,dive,required"`
	MainUri     string   `json:"mainUri" bson:"main_uri" validate:"required,url,startswith=http|startswith=https"`
	Uris        []string `json:"uris" bson:"uris" validate:"required,dive,url,startswith=http|startswith=https"`
	ServiceName string   `json:"serviceName" bson:"service_name" validate:"required"`
}

type ControllerUpdateSucceededCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                        `json:"type" bson:"type" validate:"required,eq=controller.update.succeeded.v1"`
	Data         ControllerUpdateSucceededData `json:"data" bson:"data" validate:"required"`
}

func NewControllerUpdateSucceededCloudEvent() ControllerUpdateSucceededCloudEvent {
	return ControllerUpdateSucceededCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ControllerUpdateSucceededV1Type,
		Data:     ControllerUpdateSucceededData{},
	}
}
