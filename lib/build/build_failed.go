package build

import (
	"github.com/oneweave/event-api/lib"
)

// ArtifactBuildFailedData matches the ArtifactBuildFailedData schema.
type ArtifactBuildFailedData struct {
	ArtifactBuildBaseData
	FailureCode    string `json:"failureCode" bson:"failure_code" validate:"required"`
	FailureMessage string `json:"failureMessage" bson:"failure_message" validate:"required"`
}

func NewArtifactBuildFailedData() ArtifactBuildFailedData {
	return ArtifactBuildFailedData{
		ArtifactBuildBaseData: NewArtifactBuildBaseData(),
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
