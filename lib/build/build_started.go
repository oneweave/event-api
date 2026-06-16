package build

import (
	"github.com/oneweave/event-api/lib"
)

// ArtifactBuildStartedData matches the ArtifactBuildStartedData schema.
type ArtifactBuildStartedData struct {
	ArtifactBuildBaseData `json:",inline" yaml:",inline"`
}

func NewArtifactBuildStartedData() ArtifactBuildStartedData {
	return ArtifactBuildStartedData{
		ArtifactBuildBaseData: NewArtifactBuildBaseData(),
	}
}

type ArtifactBuildStartedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                   `json:"type" bson:"type" validate:"required,eq=artifact.build.started.v1"`
	Data         ArtifactBuildStartedData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactBuildStartedCloudEvent() ArtifactBuildStartedCloudEvent {
	return ArtifactBuildStartedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactBuildStartedV1Type,
		Data:     NewArtifactBuildStartedData(),
	}
}
