package release

import (
	"github.com/oneweave/event-api/lib"
)

type ArtifactReleaseRejectedData struct {
	ArtifactReleaseBaseData `json:",inline" yaml:",inline"`
	Reason                  string `json:"reason" bson:"reason" validate:"required"`
}

func NewArtifactReleaseRejectedData() ArtifactReleaseRejectedData {
	return ArtifactReleaseRejectedData{
		ArtifactReleaseBaseData: NewArtifactReleaseBaseData(),
	}
}

type ArtifactReleaseRejectedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                      `json:"type" bson:"type" validate:"required,eq=artifact.release.rejected.v1"`
	Data         ArtifactReleaseRejectedData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactReleaseRejectedCloudEvent() ArtifactReleaseRejectedCloudEvent {
	return ArtifactReleaseRejectedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactReleaseRejectedV1Type,
		Data:     NewArtifactReleaseRejectedData(),
	}
}
