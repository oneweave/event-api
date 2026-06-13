package lib

// ReleaseSourceRepository matches the repository shape in ReleaseSource.
type ReleaseSourceRepository struct {
	Kind          string  `json:"kind" bson:"kind" validate:"required,oneof=git"`
	Protocol      string  `json:"protocol" bson:"protocol" validate:"required,oneof=https ssh git"`
	URL           string  `json:"url" bson:"url" validate:"required"`
	CredentialRef *string `json:"credentialRef,omitempty" bson:"credential_ref,omitempty"`
}

func NewReleaseSourceRepository() ReleaseSourceRepository {
	return ReleaseSourceRepository{
		Kind:     "git",
		Protocol: "https",
	}
}

// ReleaseSourceBuildArg matches one build argument item in ReleaseSource.
type ReleaseSourceBuildArg struct {
	Name  string  `json:"name" bson:"name" validate:"required,alphanum"`
	Value *string `json:"value,omitempty" bson:"value,omitempty"`
}

func NewReleaseSourceBuildArg() ReleaseSourceBuildArg {
	return ReleaseSourceBuildArg{}
}

type ImageTarget struct {
	Name          string   `json:"name" bson:"name" validate:"required"`
	Kind          *string  `json:"kind,omitempty" bson:"kind,omitempty" validate:"required,oneof=oci-registry"`
	Protocol      *string  `json:"protocol,omitempty" bson:"protocol,omitempty" validate:"required,oneof=oci https"`
	BaseURL       string   `json:"baseUrl" bson:"base_url" validate:"required"`
	Namespace     string   `json:"namespace" bson:"namespace" validate:"required"`
	Tags          []string `json:"tags,omitempty" bson:"tags,omitempty" validate:"omitempty,dive,required,alphanum"`
	CredentialRef *string  `json:"credentialRef,omitempty" bson:"credential_ref,omitempty"`
}

// ReleaseImagePushTarget matches the push target section in ReleaseSource.
type ReleaseImagePushTarget struct {
	ImageTarget
}

// ReleaseImagePullTarget matches the pull target section in ReleaseSource.
type ReleaseImagePullTarget struct {
	ImageTarget
}

func NewImageTarget() ImageTarget {
	kind := "oci-registry"
	protocol := "oci"
	imageTarget := ImageTarget{
		Kind:     &kind,
		Protocol: &protocol,
	}
	return imageTarget
}

func NewReleaseSourceImagePushTarget() ReleaseImagePushTarget {
	imageTarget := NewImageTarget()
	return ReleaseImagePushTarget{ImageTarget: imageTarget}
}

func NewReleaseSourceImagePullTarget() ReleaseImagePullTarget {
	imageTarget := NewImageTarget()
	return ReleaseImagePullTarget{ImageTarget: imageTarget}
}

func NewReleaseSourceImagePullTargetFromPushTarget(pushTarget ReleaseImagePushTarget) []ReleaseImagePullTarget {
	pullTargets := make([]ReleaseImagePullTarget, len(pushTarget.Tags))
	for i := range pushTarget.Tags {
		pullTargets[i] = ReleaseImagePullTarget{
			ImageTarget: ImageTarget{
				Name:      pushTarget.Name,
				Kind:      pushTarget.Kind,
				Protocol:  pushTarget.Protocol,
				BaseURL:   pushTarget.BaseURL,
				Namespace: pushTarget.Namespace,
				Tags:      []string{pushTarget.Tags[i]},
			},
		}
	}
	return pullTargets
}

// ReleaseSource matches the releaseSource payload shape from artifact.release.requested.v1.
type ReleaseSource struct {
	Repository        ReleaseSourceRepository `json:"repository" bson:"repository" validate:"required"`
	SourceRevisionRef string                  `json:"sourceRevisionRef" bson:"source_revision_ref" validate:"required"`
	SourceRevision    string                  `json:"sourceRevision" bson:"source_revision" validate:"required"`
	ManifestFilePath  *string                 `json:"manifestFilePath,omitempty" bson:"manifest_file_path,omitempty"`
	SubPath           *string                 `json:"subPath,omitempty" bson:"sub_path,omitempty" validate:"omitempty"`
	DockerContextPath *string                 `json:"dockerContextPath,omitempty" bson:"docker_context_path,omitempty" validate:"omitempty"`
	DockerfilePath    *string                 `json:"dockerfilePath,omitempty" bson:"dockerfile_path,omitempty" validate:"omitempty"`
	BuildArgs         []ReleaseSourceBuildArg `json:"buildArgs,omitempty" bson:"build_args,omitempty" validate:"omitempty,dive"`
}

func NewReleaseSource() ReleaseSource {
	subPath := "."
	dockerContextPath := "."
	dockerfilePath := "Dockerfile"
	return ReleaseSource{
		Repository:        NewReleaseSourceRepository(),
		SourceRevisionRef: "",
		SourceRevision:    "",
		SubPath:           &subPath,
		DockerContextPath: &dockerContextPath,
		DockerfilePath:    &dockerfilePath,
	}
}

// ReleaseTarget matches the releaseTarget payload shape from artifact.release.requested.v1.
type ReleaseTarget struct {
	RequestedVersion string                 `json:"requestedVersion" bson:"requested_version" validate:"required"`
	Platform         string                 `json:"platform" bson:"platform" validate:"required,oneof=linux/amd64"`
	ImagePushTarget  ReleaseImagePushTarget `json:"imagePushTarget" bson:"image_push_target" validate:"required"`
}

func NewReleaseTarget() ReleaseTarget {
	platform := "linux/amd64"
	return ReleaseTarget{
		Platform:        platform,
		ImagePushTarget: NewReleaseSourceImagePushTarget(),
	}
}

// ArtifactBuildRequestData matches the ArtifactBuildRequestData schema.
type ArtifactBuildRequestData struct {
	ReleaseID     string        `json:"releaseId" bson:"release_id" validate:"required,uuid"`
	ArtifactID    string        `json:"artifactId" bson:"artifact_id" validate:"required,uuid"`
	ReleaseSource ReleaseSource `json:"releaseSource" bson:"release_source" validate:"required"`
	ReleaseTarget ReleaseTarget `json:"releaseTarget" bson:"release_target" validate:"required"`
}

func NewArtifactReleaseRequestData() ArtifactBuildRequestData {
	return ArtifactBuildRequestData{
		ReleaseSource: NewReleaseSource(),
		ReleaseTarget: NewReleaseTarget(),
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
	BuildID           string                   `json:"buildId" bson:"build_id" validate:"required,uuid"`
	ServiceID         string                   `json:"serviceId" bson:"service_id" validate:"required,uuid"`
	SourceRevisionRef string                   `json:"sourceRevisionRef" bson:"source_revision_ref" validate:"required,alphanum"`
	SourceRevision    string                   `json:"sourceRevision" bson:"source_revision" validate:"required"`
	ServiceVersion    string                   `json:"serviceVersion" bson:"service_version" validate:"required"`
	ImagePullTargets  []ReleaseImagePullTarget `json:"imagePullTargets" bson:"image_pull_targets" validate:"required,min=1,dive"`
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
