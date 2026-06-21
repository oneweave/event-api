package release

import (
	"github.com/oneweave/event-api/lib"
)

type ArtifactReleasePublishedData struct {
	ArtifactReleaseBaseData `json:",inline" yaml:",inline"`
	ReleaseTarget           lib.ReleaseTarget `json:"releaseTarget" bson:"release_target" validate:"required"`
}

func NewArtifactReleasePublishedData() ArtifactReleasePublishedData {
	return ArtifactReleasePublishedData{
		ArtifactReleaseBaseData: NewArtifactReleaseBaseData(),
		ReleaseTarget:           lib.NewReleaseTarget(),
	}
}

type ArtifactReleasePublishedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                       `json:"type" bson:"type" validate:"required,eq=artifact.release.published.v1"`
	Data         ArtifactReleasePublishedData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactReleasePublishedCloudEvent() ArtifactReleasePublishedCloudEvent {
	return ArtifactReleasePublishedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactReleasePublishedV1Type,
		Data:     NewArtifactReleasePublishedData(),
	}
}
