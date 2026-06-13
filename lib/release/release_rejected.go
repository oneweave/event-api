package release

import (
	"github.com/oneweave/event-api/lib"
)

type ArtifactReleaseRejectedData struct {
	ArtifactReleaseBaseData
	Reason string `json:"reason" bson:"reason" validate:"required"`
}

func NewArtifactReleaseRejectedData() ArtifactReleaseRejectedData {
	return ArtifactReleaseRejectedData{
		ArtifactReleaseBaseData: NewArtifactReleaseBaseData(),
	}
}

type ArtifactReleaseRejectedCloudEvent struct {
	lib.Envelope
	Type string                      `json:"type" bson:"type" validate:"required,eq=artifact.release.rejected.v1"`
	Data ArtifactReleaseRejectedData `json:"data" bson:"data" validate:"required,dive"`
}

func NewArtifactReleaseRejectedCloudEvent() ArtifactReleaseRejectedCloudEvent {
	return ArtifactReleaseRejectedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ReleaseRejectedV1Type,
		Data:     NewArtifactReleaseRejectedData(),
	}
}
