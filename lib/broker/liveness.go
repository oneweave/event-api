package broker

import (
	"github.com/oneweave/event-api/lib"
)

const (
	ServiceBrokerLivenessRequestedType = "service.broker.liveness.requested"
	ServiceBrokerLivenessPassedType    = "service.broker.liveness.passed"
	ServiceBrokerLivenessFailedType    = "service.broker.liveness.failed"
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

type ServiceBrokerLivenessRequestedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                `json:"type" bson:"type" validate:"required,eq=service.broker.liveness.requested"`
	Data         LivenessRequestedData `json:"data" bson:"data"`
}

func NewServiceBrokerLivenessRequestedCloudEvent() ServiceBrokerLivenessRequestedCloudEvent {
	return ServiceBrokerLivenessRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ServiceBrokerLivenessRequestedType,
		Data:     NewLivenessRequestedData(),
	}
}

type ServiceBrokerLivenessPassedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string             `json:"type" bson:"type" validate:"required,eq=service.broker.liveness.passed"`
	Data         LivenessPassedData `json:"data" bson:"data" validate:"required"`
}

func NewServiceBrokerLivenessPassedCloudEvent() ServiceBrokerLivenessPassedCloudEvent {
	return ServiceBrokerLivenessPassedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ServiceBrokerLivenessPassedType,
		Data:     NewLivenessPassedData(),
	}
}

type ServiceBrokerLivenessFailedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string             `json:"type" bson:"type" validate:"required,eq=service.broker.liveness.failed"`
	Data         LivenessFailedData `json:"data" bson:"data" validate:"required"`
}

func NewServiceBrokerLivenessFailedCloudEvent(errStr string) ServiceBrokerLivenessFailedCloudEvent {
	return ServiceBrokerLivenessFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ServiceBrokerLivenessFailedType,
		Data:     NewLivenessFailedData(errStr),
	}
}
