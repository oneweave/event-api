package lib

type PluginManifestMetadata struct {
	Namespace   string  `json:"namespace" bson:"namespace" validate:"required"`
	Name        string  `json:"name" bson:"name" validate:"required"`
	Version     string  `json:"version" bson:"version" validate:"required"`
	Description *string `json:"description,omitempty" bson:"description,omitempty"`
	Owner       *string `json:"owner,omitempty" bson:"owner,omitempty"`
}

func NewPluginManifestMetadata() PluginManifestMetadata { return PluginManifestMetadata{} }

type PluginManifestRestEndpoint struct {
	Path           string   `json:"path" bson:"path" validate:"required,startswith=/"`
	Methods        []string `json:"methods" bson:"methods" validate:"required,dive,required"`
	AllowAnonymous *bool    `json:"allowAnonymous,omitempty" bson:"allow_anonymous,omitempty"`
}

func NewPluginManifestRestEndpoint() PluginManifestRestEndpoint { return PluginManifestRestEndpoint{} }

type PluginManifestEventDescriptor struct {
	Name                 string   `json:"name" bson:"name" validate:"required"`
	RequiresCapabilities []string `json:"requiresCapabilities,omitempty" bson:"requires_capabilities,omitempty" validate:"omitempty,dive,required"`
	Required             *bool    `json:"required,omitempty" bson:"required,omitempty"`
}

func NewPluginManifestEventDescriptor() PluginManifestEventDescriptor {
	return PluginManifestEventDescriptor{
		RequiresCapabilities: []string{},
	}
}

type PluginManifestDependency struct {
	Name     string  `json:"name" bson:"name" validate:"required"`
	Version  *string `json:"version,omitempty" bson:"version,omitempty"`
	Required *bool   `json:"required,omitempty" bson:"required,omitempty"`
}

func NewPluginManifestDependency() PluginManifestDependency { return PluginManifestDependency{} }

type PluginManifestPermissionDataAccess struct {
	Owns      []string `json:"owns" bson:"owns" validate:"required,dive"`
	DependsOn []string `json:"dependsOn" bson:"depends_on" validate:"required,dive"`
}

func NewPluginManifestPermissionDataAccess() PluginManifestPermissionDataAccess {
	return PluginManifestPermissionDataAccess{
		Owns:      []string{},
		DependsOn: []string{},
	}
}

type PluginManifestEnvironmentVariable struct {
	Key         string  `json:"key" bson:"key" validate:"required"`
	Value       *string `json:"value,omitempty" bson:"value,omitempty"`
	Required    *bool   `json:"required,omitempty" bson:"required,omitempty"`
	Description *string `json:"description,omitempty" bson:"description,omitempty"`
}

func NewPluginManifestEnvironmentVariable() PluginManifestEnvironmentVariable {
	return PluginManifestEnvironmentVariable{}
}

type PluginManifestConfigFile struct {
	Path        string  `json:"path" bson:"path" validate:"required,startswith=/"`
	Required    *bool   `json:"required,omitempty" bson:"required,omitempty"`
	Description *string `json:"description,omitempty" bson:"description,omitempty"`
}

func NewPluginManifestConfigFile() PluginManifestConfigFile { return PluginManifestConfigFile{} }

type PluginManifestArtifactRegistry struct {
	Kind          string  `json:"kind" bson:"kind" validate:"required,oneof=oci-registry"`
	Protocol      string  `json:"protocol" bson:"protocol" validate:"required,oneof=oci https"`
	BaseURL       string  `json:"baseUrl" bson:"base_url" validate:"required"`
	Namespace     string  `json:"namespace" bson:"namespace" validate:"required"`
	CredentialRef *string `json:"credentialRef,omitempty" bson:"credential_ref,omitempty"`
}

func NewPluginManifestArtifactRegistry() PluginManifestArtifactRegistry {
	return PluginManifestArtifactRegistry{Kind: "oci-registry", Protocol: "oci"}
}

type PluginManifestResources struct {
	CPU    string `json:"cpu" bson:"cpu" validate:"required"`
	Memory string `json:"memory" bson:"memory" validate:"required"`
}

func NewPluginManifestResources() PluginManifestResources {
	return PluginManifestResources{CPU: "1", Memory: "512Mi"}
}

type PluginManifestHealthProbes struct {
	Liveness  string `json:"liveness" bson:"liveness" validate:"required,startswith=/"`
	Readiness string `json:"readiness" bson:"readiness" validate:"required,startswith=/"`
	Startup   string `json:"startup" bson:"startup" validate:"required,startswith=/"`
}

func NewPluginManifestHealthProbes() PluginManifestHealthProbes {
	return PluginManifestHealthProbes{
		Liveness:  "/live",
		Readiness: "/ready",
		Startup:   "/startup",
	}
}

type PluginManifestObservability struct {
	Logs                string `json:"logs" bson:"logs" validate:"required"`
	Metrics             string `json:"metrics" bson:"metrics" validate:"required"`
	Tracing             string `json:"tracing" bson:"tracing" validate:"required"`
	CorrelationIdHeader string `json:"correlationIdHeader" bson:"correlation_id_header" validate:"required"`
}

func NewPluginManifestObservability() PluginManifestObservability {
	return PluginManifestObservability{}
}

type PluginManifestCompatibility struct {
	PluginAPIVersion string `json:"pluginApiVersion" bson:"plugin_api_version" validate:"required"`
}

func NewPluginManifestCompatibility() PluginManifestCompatibility {
	return PluginManifestCompatibility{}
}

type PluginManifestREST struct {
	Public   []PluginManifestRestEndpoint `json:"public,omitempty" bson:"public,omitempty" validate:"omitempty,dive"`
	Internal []PluginManifestRestEndpoint `json:"internal,omitempty" bson:"internal,omitempty" validate:"omitempty,dive"`
}

func NewPluginManifestREST() PluginManifestREST {
	return PluginManifestREST{
		Public:   []PluginManifestRestEndpoint{},
		Internal: []PluginManifestRestEndpoint{},
	}
}

type PluginManifestEvents struct {
	Publishes []PluginManifestEventDescriptor `json:"publishes,omitempty" bson:"publishes,omitempty" validate:"omitempty,dive"`
	Consumes  []PluginManifestEventDescriptor `json:"consumes,omitempty" bson:"consumes,omitempty" validate:"omitempty,dive"`
}

func NewPluginManifestEvents() PluginManifestEvents {
	return PluginManifestEvents{
		Publishes: []PluginManifestEventDescriptor{},
		Consumes:  []PluginManifestEventDescriptor{},
	}
}

type PluginManifestInterfaces struct {
	REST   PluginManifestREST   `json:"rest" bson:"rest" validate:"required"`
	Events PluginManifestEvents `json:"events" bson:"events" validate:"required"`
}

func NewPluginManifestInterfaces() PluginManifestInterfaces {
	return PluginManifestInterfaces{
		REST:   NewPluginManifestREST(),
		Events: NewPluginManifestEvents(),
	}
}

type PluginManifestDependencies struct {
	Services []PluginManifestDependency `json:"services,omitempty" bson:"services,omitempty" validate:"omitempty,dive"`
}

func NewPluginManifestDependencies() PluginManifestDependencies {
	return PluginManifestDependencies{
		Services: []PluginManifestDependency{},
	}
}

type PluginManifestPermissions struct {
	DataAccess PluginManifestPermissionDataAccess `json:"dataAccess" bson:"data_access" validate:"required"`
}

func NewPluginManifestPermissions() PluginManifestPermissions {
	return PluginManifestPermissions{DataAccess: NewPluginManifestPermissionDataAccess()}
}

type ReplicaConfiguration struct {
	MinReplicas int `json:"minReplicas" bson:"min_replicas" validate:"required,gt=0"`
	MaxReplicas int `json:"maxReplicas" bson:"max_replicas" validate:"required,gt=0,gtefield=MinReplicas"`
}

func NewReplicaConfiguration() ReplicaConfiguration {
	return ReplicaConfiguration{
		MinReplicas: 0,
		MaxReplicas: 5,
	}
}

type PluginManifestConfiguration struct {
	Health               PluginManifestHealthProbes          `json:"health" bson:"health" validate:"required"`
	Replicas             ReplicaConfiguration                `json:"replicas" bson:"replicas" validate:"required,dive"`
	Resources            PluginManifestResources             `json:"resources" bson:"resources" validate:"required"`
	EnvironmentVariables []PluginManifestEnvironmentVariable `json:"environmentVariables,omitempty" bson:"environment_variables,omitempty" validate:"omitempty,dive"`
	Files                []PluginManifestConfigFile          `json:"files,omitempty" bson:"files,omitempty" validate:"omitempty,dive"`
}

func NewPluginManifestConfiguration() PluginManifestConfiguration {
	return PluginManifestConfiguration{
		Health:               NewPluginManifestHealthProbes(),
		Replicas:             NewReplicaConfiguration(),
		Resources:            NewPluginManifestResources(),
		EnvironmentVariables: []PluginManifestEnvironmentVariable{},
		Files:                []PluginManifestConfigFile{},
	}
}

type PluginManifestSpec struct {
	Compatibility   PluginManifestCompatibility `json:"compatibility" bson:"compatibility" validate:"required"`
	Interfaces      PluginManifestInterfaces    `json:"interfaces" bson:"interfaces" validate:"required"`
	Dependencies    PluginManifestDependencies  `json:"dependencies" bson:"dependencies" validate:"required"`
	Permissions     PluginManifestPermissions   `json:"permissions" bson:"permissions" validate:"required"`
	Configuration   PluginManifestConfiguration `json:"configuration" bson:"configuration" validate:"required"`
	ImagePullTarget ReleaseImagePullTarget      `json:"imagePullTarget" bson:"image_pull_target" validate:"required,dive"`
	Observability   PluginManifestObservability `json:"observability" bson:"observability" validate:"required"`
}

func NewPluginManifestSpec() PluginManifestSpec {
	return PluginManifestSpec{
		Compatibility:   NewPluginManifestCompatibility(),
		Interfaces:      NewPluginManifestInterfaces(),
		Dependencies:    NewPluginManifestDependencies(),
		Permissions:     NewPluginManifestPermissions(),
		Configuration:   NewPluginManifestConfiguration(),
		ImagePullTarget: NewReleaseSourceImagePullTarget(),
		Observability:   NewPluginManifestObservability(),
	}
}

type PluginManifest struct {
	APIVersion string                 `json:"apiVersion" bson:"api_version" validate:"required,eq=oneweave/v1alpha"`
	Metadata   PluginManifestMetadata `json:"metadata" bson:"metadata" validate:"required"`
	Spec       PluginManifestSpec     `json:"spec" bson:"spec" validate:"required"`
}

func NewV1AlphaPluginManifest() PluginManifest {
	return PluginManifest{
		APIVersion: "oneweave/v1alpha",
		Metadata:   NewPluginManifestMetadata(),
		Spec:       NewPluginManifestSpec(),
	}
}

func NewPluginManifest() PluginManifest {
	return NewV1AlphaPluginManifest()
}
