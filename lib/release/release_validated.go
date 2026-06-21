package release

import (
	"github.com/oneweave/event-api/lib"
)

type ArtifactReleaseValidatedData struct {
	ArtifactReleaseBaseData `json:",inline" yaml:",inline"`
	Manifest                lib.PluginManifest `json:"manifest" bson:"manifest" validate:"required"`
}

func NewArtifactReleaseValidatedData() ArtifactReleaseValidatedData {
	return ArtifactReleaseValidatedData{
		ArtifactReleaseBaseData: NewArtifactReleaseBaseData(),
		Manifest:                lib.NewPluginManifest(),
	}
}

type ArtifactReleaseValidatedCloudEvent struct {
	lib.Envelope `json:",inline" yaml:",inline"`
	Type         string                       `json:"type" bson:"type" validate:"required,eq=artifact.release.validated.v1"`
	Data         ArtifactReleaseValidatedData `json:"data" bson:"data" validate:"required"`
}

func NewArtifactReleaseValidatedCloudEvent() ArtifactReleaseValidatedCloudEvent {
	return ArtifactReleaseValidatedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactReleaseValidatedV1Type,
		Data:     NewArtifactReleaseValidatedData(),
	}
}
