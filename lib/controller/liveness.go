package controller

import (
	"github.com/oneweave/event-api/lib"
)

const (
	ServiceControllerLivenessRequestedType = "service.controller.liveness.requested"
	ServiceControllerLivenessPassedType    = "service.controller.liveness.passed"
	ServiceControllerLivenessFailedType    = "service.controller.liveness.failed"
)

// LivenessRequestedData is the payload for liveness request events.
type LivenessRequestedData struct{}

// NewLivenessRequestedData returns a new LivenessRequestedData.
func NewLivenessRequestedData() LivenessRequestedData {
	return LivenessRequestedData{}
}

// LivenessPassedData is the payload for successful liveness check events.
type LivenessPassedData struct {
	Status string `json:"status" bson:"status" validate:"required,eq=healthy"`
}

// NewLivenessPassedData returns a new LivenessPassedData.
func NewLivenessPassedData() LivenessPassedData {
	return LivenessPassedData{
		Status: "healthy",
	}
}

// LivenessFailedData is the payload for failed liveness check events.
type LivenessFailedData struct {
	Status string `json:"status" bson:"status" validate:"required,eq=unhealthy"`
	Error  string `json:"error" bson:"error" validate:"required"`
}

// NewLivenessFailedData returns a new LivenessFailedData.
func NewLivenessFailedData(errStr string) LivenessFailedData {
	return LivenessFailedData{
		Status: "unhealthy",
		Error:  errStr,
	}
}

type ServiceControllerLivenessRequestedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                `json:"type" bson:"type" validate:"required,eq=service.controller.liveness.requested"`
	Data         LivenessRequestedData `json:"data" bson:"data"`
}

func NewServiceControllerLivenessRequestedCloudEvent() ServiceControllerLivenessRequestedCloudEvent {
	return ServiceControllerLivenessRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ServiceControllerLivenessRequestedType,
		Data:     NewLivenessRequestedData(),
	}
}

type ServiceControllerLivenessPassedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string             `json:"type" bson:"type" validate:"required,eq=service.controller.liveness.passed"`
	Data         LivenessPassedData `json:"data" bson:"data" validate:"required"`
}

func NewServiceControllerLivenessPassedCloudEvent() ServiceControllerLivenessPassedCloudEvent {
	return ServiceControllerLivenessPassedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ServiceControllerLivenessPassedType,
		Data:     NewLivenessPassedData(),
	}
}

type ServiceControllerLivenessFailedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string             `json:"type" bson:"type" validate:"required,eq=service.controller.liveness.failed"`
	Data         LivenessFailedData `json:"data" bson:"data" validate:"required"`
}

func NewServiceControllerLivenessFailedCloudEvent(errStr string) ServiceControllerLivenessFailedCloudEvent {
	return ServiceControllerLivenessFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ServiceControllerLivenessFailedType,
		Data:     NewLivenessFailedData(errStr),
	}
}
