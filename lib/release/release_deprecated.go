package release

import (
	"github.com/oneweave/event-api/lib"
)

type ArtifactReleaseDeprecatedData struct {
	ArtifactReleaseBaseData
	Reason string `json:"reason" bson:"reason" validate:"required"`
}

func NewArtifactReleaseDeprecatedData() ArtifactReleaseDeprecatedData {
	return ArtifactReleaseDeprecatedData{
		ArtifactReleaseBaseData: NewArtifactReleaseBaseData(),
	}
}

type ArtifactReleaseDeprecatedCloudEvent struct {
	lib.Envelope
	Type string                        `json:"type" bson:"type" validate:"required,eq=artifact.release.deprecated.v1"`
	Data ArtifactReleaseDeprecatedData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactReleaseDeprecatedCloudEvent() ArtifactReleaseDeprecatedCloudEvent {
	return ArtifactReleaseDeprecatedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ReleaseDeprecatedV1Type,
		Data:     NewArtifactReleaseDeprecatedData(),
	}
}
