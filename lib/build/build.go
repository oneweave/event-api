package build

const (
	Prefix = "bld"

	ArtifactBuildRequestedV1Type = "artifact.build.requested.v1"
	ArtifactBuildStartedV1Type   = "artifact.build.started.v1"
	ArtifactBuildSucceededV1Type = "artifact.build.succeeded.v1"
	ArtifactBuildFailedV1Type    = "artifact.build.failed.v1"
)

type BuildSourceInfo struct {
	Revision  string `json:"revision" bson:"revision" validate:"required"`
	Commitish string `json:"commitish" bson:"commitish" validate:"required,alphanum,len=40"`
}

func NewBuildSourceInfo() BuildSourceInfo {
	return BuildSourceInfo{}
}

type ArtifactBuildBaseData struct {
	BuildID    string          `json:"buildId" bson:"build_id" validate:"required,eventid"`
	ReleaseID  string          `json:"releaseId" bson:"release_id" validate:"required,eventid"`
	SourceInfo BuildSourceInfo `json:"sourceInfo" bson:"source_info" validate:"omitempty"`
}

func NewArtifactBuildBaseData() ArtifactBuildBaseData {
	return ArtifactBuildBaseData{}
}
