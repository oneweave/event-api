package release

import (
	"github.com/oneweave/event-api/lib"
)

type ArtifactReleaseValidatedCloudEvent struct {
	lib.Envelope
	Type string                  `json:"type" bson:"type" validate:"required,eq=artifact.release.validated.v1"`
	Data ArtifactReleaseBaseData `json:"data" bson:"data" validate:"required,dive"`
}

func NewArtifactReleaseValidatedCloudEvent() ArtifactReleaseValidatedCloudEvent {
	return ArtifactReleaseValidatedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ReleaseValidatedV1Type,
		Data:     NewArtifactReleaseBaseData(),
	}
}
