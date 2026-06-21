package build

import (
	"github.com/oneweave/event-api/lib"
)

// ArtifactBuildSucceededData matches the ArtifactBuildSucceededData schema.
type ArtifactBuildSucceededData struct {
	ArtifactBuildBaseData `json:",inline" yaml:",inline"`
	ReleaseSource         lib.ReleaseSource  `json:"releaseSource" bson:"release_source" validate:"required"`
	ReleaseTarget         lib.ReleaseTarget  `json:"releaseTarget" bson:"release_target" validate:"required"`
	Manifest              lib.PluginManifest `json:"manifest" bson:"manifest" validate:"required"`
}

func NewArtifactBuildSucceededData() ArtifactBuildSucceededData {
	return ArtifactBuildSucceededData{
		ArtifactBuildBaseData: NewArtifactBuildBaseData(),
		ReleaseSource:         lib.NewReleaseSource(),
		ReleaseTarget:         lib.NewReleaseTarget(),
		Manifest:              lib.NewPluginManifest(),
	}
}

type ArtifactBuildSucceededCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                     `json:"type" bson:"type" validate:"required,eq=artifact.build.succeeded.v1"`
	Data         ArtifactBuildSucceededData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactBuildSucceededCloudEvent() ArtifactBuildSucceededCloudEvent {
	return ArtifactBuildSucceededCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactBuildSucceededV1Type,
		Data:     NewArtifactBuildSucceededData(),
	}
}
