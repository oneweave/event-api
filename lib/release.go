package lib

// BuilderSourceRepository matches the repository shape in BuilderSource.
type BuilderSourceRepository struct {
	Kind          string  `json:"kind" validate:"required,oneof=git"`
	Protocol      string  `json:"protocol" validate:"required,oneof=https ssh git"`
	URL           string  `json:"url" validate:"required"`
	CredentialRef *string `json:"credentialRef,omitempty"`
}

func NewBuilderSourceRepository() BuilderSourceRepository {
	return BuilderSourceRepository{
		Kind:     "git",
		Protocol: "https",
	}
}

// BuilderSourceBuildArg matches one build argument item in BuilderSource.
type BuilderSourceBuildArg struct {
	Name  string  `json:"name" validate:"required,alphanum"`
	Value *string `json:"value,omitempty"`
}

func NewBuilderSourceBuildArg() BuilderSourceBuildArg {
	return BuilderSourceBuildArg{}
}

// BuilderSourceImagePushTarget matches the push target section in BuilderSource.
type BuilderSourceImagePushTarget struct {
	Name          string   `json:"name" validate:"required"`
	Kind          *string  `json:"kind,omitempty" validate:"required,oneof=oci-registry"`
	Protocol      *string  `json:"protocol,omitempty" validate:"required,oneof=oci https"`
	BaseURL       string   `json:"baseUrl" validate:"required"`
	Namespace     string   `json:"namespace" validate:"required"`
	Tags          []string `json:"tags,omitempty" validate:"omitempty,dive,required,alphanum"`
	CredentialRef *string  `json:"credentialRef,omitempty"`
}

// BuilderSourceImagePullTarget matches the pull target section in BuilderSource.
type BuilderSourceImagePullTarget struct {
	Name          string   `json:"name" validate:"required"`
	Kind          *string  `json:"kind,omitempty" validate:"required,oneof=oci-registry"`
	Protocol      *string  `json:"protocol,omitempty" validate:"required,oneof=oci https"`
	BaseURL       string   `json:"baseUrl" validate:"required"`
	Namespace     string   `json:"namespace" validate:"required"`
	Tags          []string `json:"tags,omitempty" validate:"omitempty,dive,required,alphanum"`
	CredentialRef *string  `json:"credentialRef,omitempty"`
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
	RequestedArtifactVersion string                       `json:"requestedArtifactVersion" validate:"required"`
	Repository               BuilderSourceRepository      `json:"repository" validate:"required"`
	SourceRevisionRef        string                       `json:"sourceRevisionRef" validate:"required"`
	SourceRevision           string                       `json:"sourceRevision" validate:"required"`
	ManifestFilePath         *string                      `json:"manifestFilePath,omitempty"`
	SubPath                  *string                      `json:"subPath,omitempty" validate:"omitempty"`
	DockerContextPath        *string                      `json:"dockerContextPath,omitempty" validate:"omitempty"`
	DockerfilePath           *string                      `json:"dockerfilePath,omitempty" validate:"omitempty"`
	Platform                 *string                      `json:"platform,omitempty" validate:"omitempty,oneof=linux/amd64"`
	BuildArgs                []BuilderSourceBuildArg      `json:"buildArgs,omitempty" validate:"omitempty,dive"`
	ImagePushTarget          BuilderSourceImagePushTarget `json:"imagePushTarget" validate:"required"`
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
	ReleaseID  string        `json:"releaseId" validate:"required,uuid"`
	ArtifactID string        `json:"artifactId" validate:"required,uuid"`
	BuildInput BuilderSource `json:"buildInput" validate:"required"`
}

func NewArtifactBuildRequestData() ArtifactBuildRequestData {
	return ArtifactBuildRequestData{
		BuildInput: NewBuilderSource(),
	}
}

// ArtifactReleaseArtifactIDData matches the ArtifactReleaseArtifactIdData schema.
type ArtifactReleaseArtifactIDData struct {
	ArtifactID string `json:"artifactId" validate:"required,uuid"`
}

func NewArtifactReleaseArtifactIDData() ArtifactReleaseArtifactIDData {
	return ArtifactReleaseArtifactIDData{}
}

// ArtifactBuildStartedData matches the ArtifactBuildStartedData schema.
type ArtifactBuildStartedData struct {
	BuildID           string `json:"buildId" validate:"required,uuid"`
	ServiceID         string `json:"serviceId" validate:"required,uuid"`
	SourceRevisionRef string `json:"sourceRevisionRef" validate:"required,alphanum"`
	SourceRevision    string `json:"sourceRevision" validate:"required"`
	ServiceVersion    string `json:"serviceVersion" validate:"required"`
}

func NewArtifactBuildStartedData() ArtifactBuildStartedData {
	return ArtifactBuildStartedData{}
}

// ArtifactBuildSucceededData matches the ArtifactBuildSucceededData schema.
type ArtifactBuildSucceededData struct {
	BuildID           string                         `json:"buildId" validate:"required,uuid"`
	ServiceID         string                         `json:"serviceId" validate:"required,uuid"`
	SourceRevisionRef string                         `json:"sourceRevisionRef" validate:"required,alphanum"`
	SourceRevision    string                         `json:"sourceRevision" validate:"required"`
	ServiceVersion    string                         `json:"serviceVersion" validate:"required"`
	ImagePullTargets  []BuilderSourceImagePullTarget `json:"imagePullTargets" validate:"required,min=1,dive"`
}

func NewArtifactBuildSucceededData() ArtifactBuildSucceededData {
	return ArtifactBuildSucceededData{}
}

// ArtifactBuildFailedData matches the ArtifactBuildFailedData schema.
type ArtifactBuildFailedData struct {
	BuildID           string  `json:"buildId" validate:"required,uuid"`
	ServiceID         string  `json:"serviceId" validate:"required,uuid"`
	SourceRevisionRef string  `json:"sourceRevisionRef" validate:"required,alphanum"`
	SourceRevision    string  `json:"sourceRevision" validate:"required"`
	ServiceVersion    string  `json:"serviceVersion" validate:"required"`
	FailureCode       *string `json:"failureCode,omitempty" validate:"omitempty"`
	FailureMessage    *string `json:"failureMessage,omitempty" validate:"omitempty"`
}

func NewArtifactBuildFailedData() ArtifactBuildFailedData {
	return ArtifactBuildFailedData{}
}
