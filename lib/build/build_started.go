package build

import (
	"github.com/oneweave/event-api/lib"
)

// ArtifactBuildStartedData matches the ArtifactBuildStartedData schema.
type ArtifactBuildStartedData struct {
	BuildID           string `json:"buildId" bson:"build_id" validate:"required,uuid"`
	ServiceID         string `json:"serviceId" bson:"service_id" validate:"required,uuid"`
	SourceRevisionRef string `json:"sourceRevisionRef" bson:"source_revision_ref" validate:"required,alphanum"`
	SourceRevision    string `json:"sourceRevision" bson:"source_revision" validate:"required"`
	ServiceVersion    string `json:"serviceVersion" bson:"service_version" validate:"required"`
}

func NewArtifactBuildStartedData() ArtifactBuildStartedData {
	return ArtifactBuildStartedData{}
}

type ArtifactBuildStartedCloudEvent struct {
	lib.Envelope
	Type string                   `json:"type" bson:"type" validate:"required,eq=artifact.build.started.v1"`
	Data ArtifactBuildStartedData `json:"data" bson:"data" validate:"required,dive"`
}

func NewArtifactBuildStartedCloudEvent() ArtifactBuildStartedCloudEvent {
	return ArtifactBuildStartedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactBuildStartedV1Type,
		Data:     NewArtifactBuildStartedData(),
	}
}
