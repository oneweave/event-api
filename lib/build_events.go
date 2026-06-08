package lib

// define types for each event and wir them into the constructors for better maintainability and readability
const (
	ArtifactBuildStartedV1Type   = "artifact.build.started.v1"
	ArtifactBuildSucceededV1Type = "artifact.build.succeeded.v1"
	ArtifactBuildFailedV1Type    = "artifact.build.failed.v1"
)

// ArtifactBuildStartedCloudEvent matches the ArtifactBuildStartedCloudEvent schema.
type ArtifactBuildStartedCloudEvent struct {
	Envelope
	Type string                   `json:"type" bson:"type" validate:"required,eq=artifact.build.started.v1"`
	Data ArtifactBuildStartedData `json:"data" bson:"data" validate:"required,dive"`
}

func NewArtifactBuildStartedCloudEvent() ArtifactBuildStartedCloudEvent {
	return ArtifactBuildStartedCloudEvent{
		Envelope: NewBaseEvent(),
		Type:     ArtifactBuildStartedV1Type,
		Data:     NewArtifactBuildStartedData(),
	}
}

// ArtifactBuildSucceededCloudEvent matches the ArtifactBuildSucceededCloudEvent schema.
type ArtifactBuildSucceededCloudEvent struct {
	Envelope
	Type string                     `json:"type" bson:"type" validate:"required,eq=artifact.build.succeeded.v1"`
	Data ArtifactBuildSucceededData `json:"data" bson:"data" validate:"required,dive"`
}

func NewArtifactBuildSucceededCloudEvent() ArtifactBuildSucceededCloudEvent {
	return ArtifactBuildSucceededCloudEvent{
		Envelope: NewBaseEvent(),
		Type:     ArtifactBuildSucceededV1Type,
		Data:     NewArtifactBuildSucceededData(),
	}
}

// ArtifactBuildFailedCloudEvent matches the ArtifactBuildFailedCloudEvent schema.
type ArtifactBuildFailedCloudEvent struct {
	Envelope
	Type string                  `json:"type" bson:"type" validate:"required,eq=artifact.build.failed.v1"`
	Data ArtifactBuildFailedData `json:"data" bson:"data" validate:"required,dive"`
}

func NewArtifactBuildFailedCloudEvent() ArtifactBuildFailedCloudEvent {
	return ArtifactBuildFailedCloudEvent{
		Envelope: NewBaseEvent(),
		Type:     ArtifactBuildFailedV1Type,
		Data:     NewArtifactBuildFailedData(),
	}
}
