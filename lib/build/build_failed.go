package build

import (
	"github.com/oneweave/event-api/lib"
)

// ArtifactBuildFailedData matches the ArtifactBuildFailedData schema.
type ArtifactBuildFailedData struct {
	ArtifactBuildBaseData `json:",inline" yaml:",inline"`
	lib.EventFailure      `json:",inline" yaml:",inline"`
}

func NewArtifactBuildFailedData() ArtifactBuildFailedData {
	return ArtifactBuildFailedData{
		ArtifactBuildBaseData: NewArtifactBuildBaseData(),
		EventFailure:          lib.NewEventFailure(),
	}
}

type ArtifactBuildFailedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                  `json:"type" bson:"type" validate:"required,eq=artifact.build.failed.v1"`
	Data         ArtifactBuildFailedData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactBuildFailedCloudEvent() ArtifactBuildFailedCloudEvent {
	return ArtifactBuildFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactBuildFailedV1Type,
		Data:     NewArtifactBuildFailedData(),
	}
}
