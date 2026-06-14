package build

import (
	"github.com/oneweave/event-api/lib"
)

// ArtifactBuildFailedData matches the ArtifactBuildFailedData schema.
type ArtifactBuildFailedData struct {
	ArtifactBuildBaseData
	lib.EventFailure
}

func NewArtifactBuildFailedData() ArtifactBuildFailedData {
	return ArtifactBuildFailedData{
		ArtifactBuildBaseData: NewArtifactBuildBaseData(),
		EventFailure:          lib.NewEventFailure(),
	}
}

type ArtifactBuildFailedCloudEvent struct {
	lib.Envelope
	Type string                  `json:"type" bson:"type" validate:"required,eq=artifact.build.failed.v1"`
	Data ArtifactBuildFailedData `json:"data" bson:"data" validate:"required,dive"`
}

func NewArtifactBuildFailedCloudEvent() ArtifactBuildFailedCloudEvent {
	return ArtifactBuildFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactBuildFailedV1Type,
		Data:     NewArtifactBuildFailedData(),
	}
}
