package release

import (
	"github.com/oneweave/event-api/lib"
)

type ArtifactReleaseRequestedData struct {
	ArtifactReleaseBaseData
	ReleaseSource lib.ReleaseSource `json:"releaseSource" bson:"release_source" validate:"required"`
	ReleaseTarget lib.ReleaseTarget `json:"releaseTarget" bson:"release_target" validate:"required"`
}

func NewArtifactReleaseRequestData() ArtifactReleaseRequestedData {
	return ArtifactReleaseRequestedData{
		ArtifactReleaseBaseData: NewArtifactReleaseBaseData(),
		ReleaseSource:           lib.NewReleaseSource(),
		ReleaseTarget:           lib.NewReleaseTarget(),
	}
}

type ArtifactReleaseRequestedCloudEvent struct {
	lib.Envelope
	Type string                       `json:"type" bson:"type" validate:"required,eq=artifact.release.requested.v1"`
	Data ArtifactReleaseRequestedData `json:"data" bson:"data" validate:"required,dive"`
}

func NewArtifactReleaseRequestedCloudEvent() ArtifactReleaseRequestedCloudEvent {
	return ArtifactReleaseRequestedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ReleaseRequestedV1Type,
		Data:     NewArtifactReleaseRequestData(),
	}
}
