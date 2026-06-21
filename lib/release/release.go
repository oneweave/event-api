package release

const (
	Prefix = "rel"

	ArtifactReleaseValidatedV1Type = "artifact.release.validated.v1"
	ArtifactReleaseRejectedV1Type  = "artifact.release.rejected.v1"

	ArtifactReleaseRequestedV1Type  = "artifact.release.requested.v1"
	ArtifactReleasePublishedV1Type  = "artifact.release.published.v1"
	ArtifactReleaseDeprecatedV1Type = "artifact.release.deprecated.v1"
	ArtifactReleaseRetiredV1Type    = "artifact.release.retired.v1"
)

type ArtifactReleaseBaseData struct {
	ReleaseID  string `json:"releaseId" bson:"release_id" validate:"required,eventid"`
	ArtifactID string `json:"artifactId" bson:"artifact_id" validate:"required,eventid"`
}

func NewArtifactReleaseBaseData() ArtifactReleaseBaseData {
	return ArtifactReleaseBaseData{}
}
