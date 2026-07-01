package registry

import (
	"github.com/oneweave/event-api/lib"
)

const (
	ServiceRegistryLivenessRequestedType = "service.registry.liveness.requested"
	ServiceRegistryLivenessPassedType    = "service.registry.liveness.passed"
	ServiceRegistryLivenessFailedType    = "service.registry.liveness.failed"
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

type ServiceRegistryLivenessRequestedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                `json:"type" bson:"type" validate:"required,eq=service.registry.liveness.requested"`
	Data         LivenessRequestedData `json:"data" bson:"data"`
}

func NewServiceRegistryLivenessRequestedCloudEvent() ServiceRegistryLivenessRequestedCloudEvent {
	return ServiceRegistryLivenessRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ServiceRegistryLivenessRequestedType,
		Data:     NewLivenessRequestedData(),
	}
}

type ServiceRegistryLivenessPassedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string             `json:"type" bson:"type" validate:"required,eq=service.registry.liveness.passed"`
	Data         LivenessPassedData `json:"data" bson:"data" validate:"required"`
}

func NewServiceRegistryLivenessPassedCloudEvent() ServiceRegistryLivenessPassedCloudEvent {
	return ServiceRegistryLivenessPassedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ServiceRegistryLivenessPassedType,
		Data:     NewLivenessPassedData(),
	}
}

type ServiceRegistryLivenessFailedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string             `json:"type" bson:"type" validate:"required,eq=service.registry.liveness.failed"`
	Data         LivenessFailedData `json:"data" bson:"data" validate:"required"`
}

func NewServiceRegistryLivenessFailedCloudEvent(errStr string) ServiceRegistryLivenessFailedCloudEvent {
	return ServiceRegistryLivenessFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ServiceRegistryLivenessFailedType,
		Data:     NewLivenessFailedData(errStr),
	}
}
