package build

import (
	"github.com/oneweave/event-api/lib"
)

// ArtifactBuildFailedData matches the ArtifactBuildFailedData schema.
type ArtifactBuildFailedData struct {
	BuildID           string  `json:"buildId" bson:"build_id" validate:"required,uuid"`
	ServiceID         string  `json:"serviceId" bson:"service_id" validate:"required,uuid"`
	SourceRevisionRef string  `json:"sourceRevisionRef" bson:"source_revision_ref" validate:"required,alphanum"`
	SourceRevision    string  `json:"sourceRevision" bson:"source_revision" validate:"required"`
	ServiceVersion    string  `json:"serviceVersion" bson:"service_version" validate:"required"`
	FailureCode       *string `json:"failureCode,omitempty" bson:"failure_code,omitempty" validate:"omitempty"`
	FailureMessage    *string `json:"failureMessage,omitempty" bson:"failure_message,omitempty" validate:"omitempty"`
}

func NewArtifactBuildFailedData() ArtifactBuildFailedData {
	return ArtifactBuildFailedData{}
}

type ArtifactBuildFailedCloudEvent struct {
	lib.Envelope
	Type string                  `json:"type" bson:"type" validate:"required,eq=artifact.build.failed.v1"`
	Data ArtifactBuildFailedData `json:"data" bson:"data" validate:"required,dive"`
}

func NewArtifactBuildFailedCloudEvent() ArtifactBuildFailedCloudEvent {
	return ArtifactBuildFailedCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactBuildFailedV1Type,
		Data:     NewArtifactBuildFailedData(),
	}
}
