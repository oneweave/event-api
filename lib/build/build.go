package build

const (
	ArtifactBuildStartedV1Type   = "artifact.build.started.v1"
	ArtifactBuildSucceededV1Type = "artifact.build.succeeded.v1"
	ArtifactBuildFailedV1Type    = "artifact.build.failed.v1"
)

type ArtifactBuildBaseData struct {
	BuildID   string `json:"buildId" bson:"build_id" validate:"required,uuid"`
	ReleaseID string `json:"releaseId" bson:"release_id" validate:"required,uuid"`
}

func NewArtifactBuildBaseData() ArtifactBuildBaseData {
	return ArtifactBuildBaseData{}
}
