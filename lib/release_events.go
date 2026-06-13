package lib

const (
	ReleaseRequestedV1Type = "artifact.release.requested.v1"
	ReleaseValidatedV1Type = "artifact.release.validated.v1"
)

type ArtifactReleaseRequestedCloudEvent struct {
	Envelope
	Type string                   `json:"type" bson:"type" validate:"required,eq=artifact.release.requested.v1"`
	Data ArtifactBuildRequestData `json:"data" bson:"data" validate:"required,dive"`
}

func NewArtifactReleaseRequestedCloudEvent() ArtifactReleaseRequestedCloudEvent {
	return ArtifactReleaseRequestedCloudEvent{
		Envelope: NewEnvelope(),
		Type:     ReleaseRequestedV1Type,
		Data:     NewArtifactReleaseRequestData(),
	}
}

type ArtifactReleaseValidatedCloudEvent struct {
	Envelope
	Type string                        `json:"type" bson:"type" validate:"required,eq=artifact.release.validated.v1"`
	Data ArtifactReleaseArtifactIDData `json:"data" bson:"data" validate:"required,dive"`
}

func NewArtifactReleaseValidatedCloudEvent() ArtifactReleaseValidatedCloudEvent {
	return ArtifactReleaseValidatedCloudEvent{
		Envelope: NewEnvelope(),
		Type:     ReleaseValidatedV1Type,
		Data:     NewArtifactReleaseArtifactIDData(),
	}
}
