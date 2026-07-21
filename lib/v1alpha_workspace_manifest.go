package lib

const (
	workspaceManifestKind = "Workspace"
)

type WorkspaceService struct {
	Name             string  `json:"name" bson:"name" validate:"required"`
	Path             string  `json:"path" bson:"path" validate:"required"`
	ManifestFilePath *string `json:"manifestFilePath,omitempty" bson:"manifest_file_path,omitempty" validate:"omitempty"`
}

func NewWorkspaceService() WorkspaceService {
	manifestFilePath := "weave.yaml"
	return WorkspaceService{
		Path:             ".",
		ManifestFilePath: &manifestFilePath,
	}
}

type WorkspaceManifestSpec struct {
	Services []WorkspaceService `json:"services" bson:"services" validate:"required"`
}

type WorkspaceManifest struct {
	APIVersion string                 `json:"apiVersion" bson:"api_version" validate:"required,eq=oneweave/v1alpha"`
	Kind       string                 `json:"kind" bson:"kind" validate:"required,eq=Workspace"`
	Metadata   PluginManifestMetadata `json:"metadata" bson:"metadata" validate:"required"`
	Spec       WorkspaceManifestSpec  `json:"spec" bson:"spec" validate:"required"`
}

func NewWorkspaceManifest() WorkspaceManifest {
	return WorkspaceManifest{
		APIVersion: v1alphaPluginManifestAPIVersion,
		Kind:       workspaceManifestKind,
		Metadata:   NewPluginManifestMetadata(),
		Spec:       NewWorkspaceManifestSpec(),
	}
}

func NewWorkspaceManifestSpec() WorkspaceManifestSpec {
	return WorkspaceManifestSpec{
		Services: []WorkspaceService{},
	}
}
