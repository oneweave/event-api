package build

import (
	"github.com/oneweave/event-api/lib"
)

// ArtifactBuildSucceededData matches the ArtifactBuildSucceededData schema.
type ArtifactBuildSucceededData struct {
	BuildID           string                       `json:"buildId" bson:"build_id" validate:"required,uuid"`
	ServiceID         string                       `json:"serviceId" bson:"service_id" validate:"required,uuid"`
	SourceRevisionRef string                       `json:"sourceRevisionRef" bson:"source_revision_ref" validate:"required,alphanum"`
	SourceRevision    string                       `json:"sourceRevision" bson:"source_revision" validate:"required"`
	ServiceVersion    string                       `json:"serviceVersion" bson:"service_version" validate:"required"`
	ImagePullTargets  []lib.ReleaseImagePullTarget `json:"imagePullTargets" bson:"image_pull_targets" validate:"required,min=1,dive"`
}

func NewArtifactBuildSucceededData() ArtifactBuildSucceededData {
	return ArtifactBuildSucceededData{}
}

type ArtifactBuildSucceededCloudEvent struct {
	lib.Envelope
	Type string                     `json:"type" bson:"type" validate:"required,eq=artifact.build.succeeded.v1"`
	Data ArtifactBuildSucceededData `json:"data" bson:"data" validate:"required,dive"`
}

func NewArtifactBuildSucceededCloudEvent() ArtifactBuildSucceededCloudEvent {
	return ArtifactBuildSucceededCloudEvent{
		Envelope: lib.NewEnvelope(),
		Type:     ArtifactBuildSucceededV1Type,
		Data:     NewArtifactBuildSucceededData(),
	}
}
