package lib

// BuilderSourceRepository matches the repository shape in BuilderSource.
type BuilderSourceRepository struct {
	Kind          string  `json:"kind" bson:"kind" validate:"required,oneof=git"`
	Protocol      string  `json:"protocol" bson:"protocol" validate:"required,oneof=https ssh git"`
	URL           string  `json:"url" bson:"url" validate:"required"`
	CredentialRef *string `json:"credentialRef,omitempty" bson:"credential_ref,omitempty"`
}

func NewBuilderSourceRepository() BuilderSourceRepository {
	return BuilderSourceRepository{
		Kind:     "git",
		Protocol: "https",
	}
}

// BuilderSourceBuildArg matches one build argument item in BuilderSource.
type BuilderSourceBuildArg struct {
	Name  string  `json:"name" bson:"name" validate:"required,alphanum"`
	Value *string `json:"value,omitempty" bson:"value,omitempty"`
}

func NewBuilderSourceBuildArg() BuilderSourceBuildArg {
	return BuilderSourceBuildArg{}
}

// BuilderSourceImagePushTarget matches the push target section in BuilderSource.
type BuilderSourceImagePushTarget struct {
	Name          string   `json:"name" bson:"name" validate:"required"`
	Kind          *string  `json:"kind,omitempty" bson:"kind,omitempty" validate:"required,oneof=oci-registry"`
	Protocol      *string  `json:"protocol,omitempty" bson:"protocol,omitempty" validate:"required,oneof=oci https"`
	BaseURL       string   `json:"baseUrl" bson:"base_url" validate:"required"`
	Namespace     string   `json:"namespace" bson:"namespace" validate:"required"`
	Tags          []string `json:"tags,omitempty" bson:"tags,omitempty" validate:"omitempty,dive,required,alphanum"`
	CredentialRef *string  `json:"credentialRef,omitempty" bson:"credential_ref,omitempty"`
}

// BuilderSourceImagePullTarget matches the pull target section in BuilderSource.
type BuilderSourceImagePullTarget struct {
	Name          string   `json:"name" bson:"name" validate:"required"`
	Kind          *string  `json:"kind,omitempty" bson:"kind,omitempty" validate:"required,oneof=oci-registry"`
	Protocol      *string  `json:"protocol,omitempty" bson:"protocol,omitempty" validate:"required,oneof=oci https"`
	BaseURL       string   `json:"baseUrl" bson:"base_url" validate:"required"`
	Namespace     string   `json:"namespace" bson:"namespace" validate:"required"`
	Tags          []string `json:"tags,omitempty" bson:"tags,omitempty" validate:"omitempty,dive,required,alphanum"`
	CredentialRef *string  `json:"credentialRef,omitempty" bson:"credential_ref,omitempty"`
}

func NewBuilderSourceImagePushTarget() BuilderSourceImagePushTarget {
	kind := "oci-registry"
	protocol := "oci"
	return BuilderSourceImagePushTarget{Kind: &kind, Protocol: &protocol}
}

func NewBuilderSourceImagePullTarget() BuilderSourceImagePullTarget {
	kind := "oci-registry"
	protocol := "oci"
	return BuilderSourceImagePullTarget{Kind: &kind, Protocol: &protocol}
}

// BuilderSource matches the buildInput payload shape from artifact.release.requested.v1.
type BuilderSource struct {
	RequestedArtifactVersion string                       `json:"requestedArtifactVersion" bson:"requested_artifact_version" validate:"required"`
	Repository               BuilderSourceRepository      `json:"repository" bson:"repository" validate:"required"`
	SourceRevisionRef        string                       `json:"sourceRevisionRef" bson:"source_revision_ref" validate:"required"`
	SourceRevision           string                       `json:"sourceRevision" bson:"source_revision" validate:"required"`
	ManifestFilePath         *string                      `json:"manifestFilePath,omitempty" bson:"manifest_file_path,omitempty"`
	SubPath                  *string                      `json:"subPath,omitempty" bson:"sub_path,omitempty" validate:"omitempty"`
	DockerContextPath        *string                      `json:"dockerContextPath,omitempty" bson:"docker_context_path,omitempty" validate:"omitempty"`
	DockerfilePath           *string                      `json:"dockerfilePath,omitempty" bson:"dockerfile_path,omitempty" validate:"omitempty"`
	Platform                 *string                      `json:"platform,omitempty" bson:"platform,omitempty" validate:"omitempty,oneof=linux/amd64"`
	BuildArgs                []BuilderSourceBuildArg      `json:"buildArgs,omitempty" bson:"build_args,omitempty" validate:"omitempty,dive"`
	ImagePushTarget          BuilderSourceImagePushTarget `json:"imagePushTarget" bson:"image_push_target" validate:"required"`
}

func NewBuilderSource() BuilderSource {
	subPath := "."
	dockerContextPath := "."
	dockerfilePath := "Dockerfile"
	platform := "linux/amd64"
	return BuilderSource{
		Repository:               NewBuilderSourceRepository(),
		RequestedArtifactVersion: "",
		SourceRevisionRef:        "",
		SourceRevision:           "",
		SubPath:                  &subPath,
		DockerContextPath:        &dockerContextPath,
		DockerfilePath:           &dockerfilePath,
		Platform:                 &platform,
		ImagePushTarget:          NewBuilderSourceImagePushTarget(),
	}
}

// ArtifactBuildRequestData matches the ArtifactBuildRequestData schema.
type ArtifactBuildRequestData struct {
	ReleaseID  string        `json:"releaseId" bson:"release_id" validate:"required,uuid"`
	ArtifactID string        `json:"artifactId" bson:"artifact_id" validate:"required,uuid"`
	BuildInput BuilderSource `json:"buildInput" bson:"build_input" validate:"required"`
}

func NewArtifactBuildRequestData() ArtifactBuildRequestData {
	return ArtifactBuildRequestData{
		BuildInput: NewBuilderSource(),
	}
}

// ArtifactReleaseArtifactIDData matches the ArtifactReleaseArtifactIdData schema.
type ArtifactReleaseArtifactIDData struct {
	ArtifactID string `json:"artifactId" bson:"artifact_id" validate:"required,uuid"`
}

func NewArtifactReleaseArtifactIDData() ArtifactReleaseArtifactIDData {
	return ArtifactReleaseArtifactIDData{}
}

// ArtifactBuildStartedData matches the ArtifactBuildStartedData schema.
type ArtifactBuildStartedData struct {
	BuildID           string `json:"buildId" bson:"build_id" validate:"required,uuid"`
	ServiceID         string `json:"serviceId" bson:"service_id" validate:"required,uuid"`
	SourceRevisionRef string `json:"sourceRevisionRef" bson:"source_revision_ref" validate:"required,alphanum"`
	SourceRevision    string `json:"sourceRevision" bson:"source_revision" validate:"required"`
	ServiceVersion    string `json:"serviceVersion" bson:"service_version" validate:"required"`
}

func NewArtifactBuildStartedData() ArtifactBuildStartedData {
	return ArtifactBuildStartedData{}
}

// ArtifactBuildSucceededData matches the ArtifactBuildSucceededData schema.
type ArtifactBuildSucceededData struct {
	BuildID           string                         `json:"buildId" bson:"build_id" validate:"required,uuid"`
	ServiceID         string                         `json:"serviceId" bson:"service_id" validate:"required,uuid"`
	SourceRevisionRef string                         `json:"sourceRevisionRef" bson:"source_revision_ref" validate:"required,alphanum"`
	SourceRevision    string                         `json:"sourceRevision" bson:"source_revision" validate:"required"`
	ServiceVersion    string                         `json:"serviceVersion" bson:"service_version" validate:"required"`
	ImagePullTargets  []BuilderSourceImagePullTarget `json:"imagePullTargets" bson:"image_pull_targets" validate:"required,min=1,dive"`
}

func NewArtifactBuildSucceededData() ArtifactBuildSucceededData {
	return ArtifactBuildSucceededData{}
}

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
