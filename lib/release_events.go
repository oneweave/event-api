package lib

const (
	ReleaseRequestedV1Type = "artifact.release.requested.v1"
	ReleaseValidatedV1Type = "artifact.release.validated.v1"
)

type ArtifactReleaseRequestedCloudEvent struct {
	BaseEvent
	Type string                   `json:"type" validate:"required,eq=artifact.release.requested.v1"`
	Data ArtifactBuildRequestData `json:"data" validate:"required,dive"`
}

func NewArtifactReleaseRequestedCloudEvent() ArtifactReleaseRequestedCloudEvent {
	return ArtifactReleaseRequestedCloudEvent{
		BaseEvent: NewBaseEvent(),
		Type:      ReleaseRequestedV1Type,
		Data:      NewArtifactBuildRequestData(),
	}
}

type ArtifactReleaseValidatedCloudEvent struct {
	BaseEvent
	Type string                        `json:"type" validate:"required,eq=artifact.release.validated.v1"`
	Data ArtifactReleaseArtifactIDData `json:"data" validate:"required,dive"`
}

func NewArtifactReleaseValidatedCloudEvent() ArtifactReleaseValidatedCloudEvent {
	return ArtifactReleaseValidatedCloudEvent{
		BaseEvent: NewBaseEvent(),
		Type:      ReleaseValidatedV1Type,
		Data:      NewArtifactReleaseArtifactIDData(),
	}
}
