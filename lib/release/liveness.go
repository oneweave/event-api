package release

import (
	"github.com/oneweave/event-api/lib"
)

const (
	ArtifactRegistryLivenessRequestedType = "artifact.registry.liveness.requested"
	ArtifactRegistryLivenessPassedType    = "artifact.registry.liveness.passed"
	ArtifactRegistryLivenessFailedType    = "artifact.registry.liveness.failed"
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

type ArtifactRegistryLivenessRequestedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                `json:"type" bson:"type" validate:"required,eq=artifact.registry.liveness.requested"`
	Data         LivenessRequestedData `json:"data" bson:"data"`
}

func NewArtifactRegistryLivenessRequestedCloudEvent() ArtifactRegistryLivenessRequestedCloudEvent {
	return ArtifactRegistryLivenessRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactRegistryLivenessRequestedType,
		Data:     NewLivenessRequestedData(),
	}
}

type ArtifactRegistryLivenessPassedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string             `json:"type" bson:"type" validate:"required,eq=artifact.registry.liveness.passed"`
	Data         LivenessPassedData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactRegistryLivenessPassedCloudEvent() ArtifactRegistryLivenessPassedCloudEvent {
	return ArtifactRegistryLivenessPassedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactRegistryLivenessPassedType,
		Data:     NewLivenessPassedData(),
	}
}

type ArtifactRegistryLivenessFailedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string             `json:"type" bson:"type" validate:"required,eq=artifact.registry.liveness.failed"`
	Data         LivenessFailedData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactRegistryLivenessFailedCloudEvent(errStr string) ArtifactRegistryLivenessFailedCloudEvent {
	return ArtifactRegistryLivenessFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactRegistryLivenessFailedType,
		Data:     NewLivenessFailedData(errStr),
	}
}
