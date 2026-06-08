package lib

// define types for each event and wir them into the constructors for better maintainability and readability
const (
	ArtifactBuildStartedV1Type   = "artifact.build.started.v1"
	ArtifactBuildSucceededV1Type = "artifact.build.succeeded.v1"
	ArtifactBuildFailedV1Type    = "artifact.build.failed.v1"
)

// ArtifactBuildStartedCloudEvent matches the ArtifactBuildStartedCloudEvent schema.
type ArtifactBuildStartedCloudEvent struct {
	BaseEvent
	Type string                   `json:"type" validate:"required,eq=artifact.build.started.v1"`
	Data ArtifactBuildStartedData `json:"data" validate:"required,dive"`
}

func NewArtifactBuildStartedCloudEvent() ArtifactBuildStartedCloudEvent {
	return ArtifactBuildStartedCloudEvent{
		BaseEvent: NewBaseEvent(),
		Type:      ArtifactBuildStartedV1Type,
		Data:      NewArtifactBuildStartedData(),
	}
}

// ArtifactBuildSucceededCloudEvent matches the ArtifactBuildSucceededCloudEvent schema.
type ArtifactBuildSucceededCloudEvent struct {
	BaseEvent
	Type string                     `json:"type" validate:"required,eq=artifact.build.succeeded.v1"`
	Data ArtifactBuildSucceededData `json:"data" validate:"required,dive"`
}

func NewArtifactBuildSucceededCloudEvent() ArtifactBuildSucceededCloudEvent {
	return ArtifactBuildSucceededCloudEvent{
		BaseEvent: NewBaseEvent(),
		Type:      ArtifactBuildSucceededV1Type,
		Data:      NewArtifactBuildSucceededData(),
	}
}

// ArtifactBuildFailedCloudEvent matches the ArtifactBuildFailedCloudEvent schema.
type ArtifactBuildFailedCloudEvent struct {
	BaseEvent
	Type string                  `json:"type" validate:"required,eq=artifact.build.failed.v1"`
	Data ArtifactBuildFailedData `json:"data" validate:"required,dive"`
}

func NewArtifactBuildFailedCloudEvent() ArtifactBuildFailedCloudEvent {
	return ArtifactBuildFailedCloudEvent{
		BaseEvent: NewBaseEvent(),
		Type:      ArtifactBuildFailedV1Type,
		Data:      NewArtifactBuildFailedData(),
	}
}
