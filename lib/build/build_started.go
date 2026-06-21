package build

import (
	"github.com/oneweave/event-api/lib"
)

// ArtifactBuildStartedData matches the ArtifactBuildStartedData schema.
type ArtifactBuildStartedData struct {
	ArtifactBuildBaseData `json:",inline" yaml:",inline"`
	ReleaseSource         lib.ReleaseSource  `json:"releaseSource" bson:"release_source" validate:"required"`
	ReleaseTarget         lib.ReleaseTarget  `json:"releaseTarget" bson:"release_target" validate:"required"`
	Manifest              lib.PluginManifest `json:"manifest" bson:"manifest,omitempty" validate:"required"`
}

func NewArtifactBuildStartedData() ArtifactBuildStartedData {
	return ArtifactBuildStartedData{
		ArtifactBuildBaseData: NewArtifactBuildBaseData(),
		ReleaseSource:         lib.NewReleaseSource(),
		ReleaseTarget:         lib.NewReleaseTarget(),
		Manifest:              lib.NewPluginManifest(),
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
