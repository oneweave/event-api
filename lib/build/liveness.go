package build

import (
	"github.com/oneweave/event-api/lib"
)

const (
	ArtifactBuilderLivenessRequestedType = "artifact.builder.liveness.requested"
	ArtifactBuilderLivenessPassedType    = "artifact.builder.liveness.passed"
	ArtifactBuilderLivenessFailedType    = "artifact.builder.liveness.failed"
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

type ArtifactBuilderLivenessRequestedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                `json:"type" bson:"type" validate:"required,eq=artifact.builder.liveness.requested"`
	Data         LivenessRequestedData `json:"data" bson:"data"`
}

func NewArtifactBuilderLivenessRequestedCloudEvent() ArtifactBuilderLivenessRequestedCloudEvent {
	return ArtifactBuilderLivenessRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactBuilderLivenessRequestedType,
		Data:     NewLivenessRequestedData(),
	}
}

type ArtifactBuilderLivenessPassedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string             `json:"type" bson:"type" validate:"required,eq=artifact.builder.liveness.passed"`
	Data         LivenessPassedData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactBuilderLivenessPassedCloudEvent() ArtifactBuilderLivenessPassedCloudEvent {
	return ArtifactBuilderLivenessPassedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactBuilderLivenessPassedType,
		Data:     NewLivenessPassedData(),
	}
}

type ArtifactBuilderLivenessFailedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string             `json:"type" bson:"type" validate:"required,eq=artifact.builder.liveness.failed"`
	Data         LivenessFailedData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactBuilderLivenessFailedCloudEvent(errStr string) ArtifactBuilderLivenessFailedCloudEvent {
	return ArtifactBuilderLivenessFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactBuilderLivenessFailedType,
		Data:     NewLivenessFailedData(errStr),
	}
}
