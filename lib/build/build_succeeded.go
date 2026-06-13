package build

import (
	"github.com/oneweave/event-api/lib"
)

// ArtifactBuildSucceededData matches the ArtifactBuildSucceededData schema.
type ArtifactBuildSucceededData struct {
	ArtifactBuildBaseData
	ImagePullTargets []lib.ReleaseImagePullTarget `json:"imagePullTargets" bson:"image_pull_targets" validate:"required,min=1,dive"`
}

func NewArtifactBuildSucceededData() ArtifactBuildSucceededData {
	return ArtifactBuildSucceededData{
		ArtifactBuildBaseData: NewArtifactBuildBaseData(),
	}
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
