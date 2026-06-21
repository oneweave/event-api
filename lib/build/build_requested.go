package build

import (
	"github.com/oneweave/event-api/lib"
)

// ArtifactBuildRequestedData matches the ArtifactBuildRequestedData schema.
type ArtifactBuildRequestedData struct {
	ArtifactBuildBaseData `json:",inline" yaml:",inline"`
	ReleaseSource         lib.ReleaseSource `json:"releaseSource" bson:"release_source" validate:"required"`
	ReleaseTarget         lib.ReleaseTarget `json:"releaseTarget" bson:"release_target" validate:"required"`
}

func NewArtifactBuildRequestedData() ArtifactBuildRequestedData {
	return ArtifactBuildRequestedData{
		ArtifactBuildBaseData: NewArtifactBuildBaseData(),
		ReleaseSource:         lib.NewReleaseSource(),
		ReleaseTarget:         lib.NewReleaseTarget(),
	}
}

type ArtifactBuildRequestedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                     `json:"type" bson:"type" validate:"required,eq=artifact.build.requested.v1"`
	Data         ArtifactBuildRequestedData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactBuildRequestedCloudEvent() ArtifactBuildRequestedCloudEvent {
	return ArtifactBuildRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactBuildRequestedV1Type,
		Data:     NewArtifactBuildRequestedData(),
	}
}
