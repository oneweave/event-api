package release

import (
	"github.com/oneweave/event-api/lib"
)

type ArtifactReleasePublishedData struct {
	ArtifactReleaseBaseData
	ReleaseTarget lib.ReleaseTarget `json:"releaseTarget" bson:"release_target" validate:"required"`
}

func NewArtifactReleasePublishedData() ArtifactReleasePublishedData {
	return ArtifactReleasePublishedData{
		ArtifactReleaseBaseData: NewArtifactReleaseBaseData(),
		ReleaseTarget:           lib.NewReleaseTarget(),
	}
}

type ArtifactReleasePublishedCloudEvent struct {
	lib.Envelope
	Type string                       `json:"type" bson:"type" validate:"required,eq=artifact.release.published.v1"`
	Data ArtifactReleasePublishedData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactReleasePublishedCloudEvent() ArtifactReleasePublishedCloudEvent {
	return ArtifactReleasePublishedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ReleasePublishedV1Type,
		Data:     NewArtifactReleasePublishedData(),
	}
}
