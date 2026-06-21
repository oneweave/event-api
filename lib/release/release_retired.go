package release

import (
	"github.com/oneweave/event-api/lib"
)

type ArtifactReleaseRetiredCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                  `json:"type" bson:"type" validate:"required,eq=artifact.release.retired.v1"`
	Data         ArtifactReleaseBaseData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactReleaseRetiredCloudEvent() ArtifactReleaseRetiredCloudEvent {
	return ArtifactReleaseRetiredCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactReleaseRetiredV1Type,
		Data:     NewArtifactReleaseBaseData(),
	}
}
